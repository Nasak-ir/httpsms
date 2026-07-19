package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/mail"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/NdoleStudio/httpsms/pkg/authlocal"
	"github.com/NdoleStudio/httpsms/pkg/entities"
	"github.com/NdoleStudio/httpsms/pkg/repositories"
	"github.com/NdoleStudio/httpsms/pkg/requests"
	"github.com/NdoleStudio/httpsms/pkg/responses"
	"github.com/NdoleStudio/httpsms/pkg/telemetry"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/palantir/stacktrace"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/proxy"
	"gorm.io/gorm"
)

// FirebaseEmailAuthHandler proxies email auth through the API server.
type FirebaseEmailAuthHandler struct {
	handler
	logger         telemetry.Logger
	tracer         telemetry.Tracer
	db             *gorm.DB
	userRepository repositories.UserRepository
	client         *http.Client
}

// NewFirebaseEmailAuthHandler creates a FirebaseEmailAuthHandler.
func NewFirebaseEmailAuthHandler(logger telemetry.Logger, tracer telemetry.Tracer, db *gorm.DB, userRepository repositories.UserRepository) *FirebaseEmailAuthHandler {
	logger = logger.WithService("handlers.FirebaseEmailAuthHandler")
	return &FirebaseEmailAuthHandler{
		logger:         logger,
		tracer:         tracer,
		db:             db,
		userRepository: userRepository,
		client: &http.Client{
			Timeout:   20 * time.Second,
			Transport: firebaseAuthTransport(logger),
		},
	}
}

// RegisterRoutes registers public Firebase email auth proxy routes.
func (h *FirebaseEmailAuthHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/v1/auth/email", h.Authenticate)
	router.Post("/v1/auth/email/refresh", h.Refresh)
}

// Authenticate signs a user in or creates an account through Firebase Identity Toolkit.
func (h *FirebaseEmailAuthHandler) Authenticate(c fiber.Ctx) error {
	_, span, ctxLogger := h.tracer.StartFromFiberCtxWithLogger(c, h.logger)
	defer span.End()

	var request requests.FirebaseEmailAuthRequest
	if err := c.Bind().Body(&request); err != nil {
		ctxLogger.Warn(stacktrace.Propagate(err, "cannot parse firebase email auth request"))
		return h.responseBadRequest(c, err)
	}
	request = request.Sanitize()

	if errors := validateFirebaseEmailAuth(request); len(errors) != 0 {
		return h.responseUnprocessableEntity(c, errors, "validation errors while authenticating")
	}

	endpoint := "accounts:signInWithPassword"
	if request.Mode == "sign_up" {
		endpoint = "accounts:signUp"
	}

	authResponse, status, err := h.callIdentityToolkit(endpoint, map[string]any{
		"email":             request.Email,
		"password":          request.Password,
		"returnSecureToken": true,
	})
	if err != nil {
		ctxLogger.Warn(stacktrace.Propagate(err, "firebase email auth failed with status [%d], trying local auth", status))
		localResponse, localErr := h.authenticateLocally(c.Context(), request)
		if localErr != nil {
			ctxLogger.Warn(stacktrace.Propagate(localErr, "local email auth failed after firebase status [%d]", status))
			return h.firebaseError(c, status, err)
		}
		return h.responseOK(c, "local email authentication successful", localResponse)
	}

	return h.responseOK(c, "firebase email authentication successful", authResponse)
}

// Refresh exchanges a Firebase refresh token for a new ID token.
func (h *FirebaseEmailAuthHandler) Refresh(c fiber.Ctx) error {
	_, span, ctxLogger := h.tracer.StartFromFiberCtxWithLogger(c, h.logger)
	defer span.End()

	var request requests.FirebaseEmailRefreshRequest
	if err := c.Bind().Body(&request); err != nil {
		ctxLogger.Warn(stacktrace.Propagate(err, "cannot parse firebase email refresh request"))
		return h.responseBadRequest(c, err)
	}
	if strings.TrimSpace(request.RefreshToken) == "" {
		errors := url.Values{}
		errors.Add("refresh_token", "refresh token is required")
		return h.responseUnprocessableEntity(c, errors, "validation errors while refreshing authentication")
	}

	authResponse, status, err := h.callSecureToken(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": request.RefreshToken,
	})
	if err != nil {
		ctxLogger.Warn(stacktrace.Propagate(err, "firebase email token refresh failed with status [%d], trying local auth", status))
		localResponse, localErr := h.refreshLocalToken(request.RefreshToken)
		if localErr != nil {
			ctxLogger.Warn(stacktrace.Propagate(localErr, "local token refresh failed after firebase status [%d]", status))
			return h.firebaseError(c, status, err)
		}
		return h.responseOK(c, "local email token refreshed successfully", localResponse)
	}

	return h.responseOK(c, "firebase email token refreshed successfully", authResponse)
}

func (h *FirebaseEmailAuthHandler) authenticateLocally(ctx context.Context, request requests.FirebaseEmailAuthRequest) (*responses.FirebaseEmailAuthResponse, error) {
	if strings.TrimSpace(os.Getenv("LOCAL_AUTH_SECRET")) == "" {
		return nil, stacktrace.NewError("LOCAL_AUTH_SECRET is not configured")
	}

	credential := new(entities.LocalAuthCredential)
	err := h.db.WithContext(ctx).Where("email = ?", request.Email).First(credential).Error
	if err == nil {
		if bcrypt.CompareHashAndPassword([]byte(credential.PasswordHash), []byte(request.Password)) != nil {
			return nil, stacktrace.NewError("invalid local email or password")
		}
		return h.localAuthResponse(entities.AuthContext{ID: credential.UserID, Email: request.Email})
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, stacktrace.Propagate(err, "cannot load local auth credential")
	}
	if request.Mode != "sign_up" {
		return nil, stacktrace.NewError("local auth credential does not exist")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot hash local auth password")
	}
	authUser := entities.AuthContext{
		ID:    entities.UserID("local:" + uuid.NewString()),
		Email: request.Email,
	}
	if existingUser, loadErr := h.userRepository.LoadByEmail(ctx, request.Email); loadErr == nil {
		authUser.ID = existingUser.ID
	}
	if _, _, err = h.userRepository.LoadOrStore(ctx, authUser); err != nil {
		return nil, stacktrace.Propagate(err, "cannot create local auth user")
	}

	credential = &entities.LocalAuthCredential{
		Email:        request.Email,
		UserID:       authUser.ID,
		PasswordHash: string(passwordHash),
	}
	if err = h.db.WithContext(ctx).Create(credential).Error; err != nil {
		return nil, stacktrace.Propagate(err, "cannot save local auth credential")
	}

	return h.localAuthResponse(authUser)
}

func (h *FirebaseEmailAuthHandler) refreshLocalToken(rawToken string) (*responses.FirebaseEmailAuthResponse, error) {
	authUser, err := authlocal.Verify(os.Getenv("LOCAL_AUTH_SECRET"), rawToken)
	if err != nil {
		return nil, err
	}
	return h.localAuthResponse(authUser)
}

func (h *FirebaseEmailAuthHandler) localAuthResponse(authUser entities.AuthContext) (*responses.FirebaseEmailAuthResponse, error) {
	const ttl = 7 * 24 * time.Hour
	token, err := authlocal.Sign(os.Getenv("LOCAL_AUTH_SECRET"), authUser, ttl)
	if err != nil {
		return nil, err
	}
	return &responses.FirebaseEmailAuthResponse{
		IDToken:      token,
		RefreshToken: token,
		ExpiresIn:    fmt.Sprintf("%.0f", ttl.Seconds()),
		LocalID:      authUser.ID.String(),
		Email:        authUser.Email,
	}, nil
}

func validateFirebaseEmailAuth(request requests.FirebaseEmailAuthRequest) url.Values {
	errors := url.Values{}
	if _, err := mail.ParseAddress(request.Email); err != nil {
		errors.Add("email", "please enter a valid email address")
	}
	if len(request.Password) < 6 {
		errors.Add("password", "password should be at least 6 characters")
	}
	if request.Mode != "sign_in" && request.Mode != "sign_up" {
		errors.Add("mode", "mode should be sign_in or sign_up")
	}
	return errors
}

func (h *FirebaseEmailAuthHandler) callIdentityToolkit(endpoint string, payload map[string]any) (*responses.FirebaseEmailAuthResponse, int, error) {
	apiKey := strings.TrimSpace(os.Getenv("FIREBASE_API_KEY"))
	if apiKey == "" {
		return nil, fiber.StatusServiceUnavailable, stacktrace.NewError("FIREBASE_API_KEY is not configured")
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fiber.StatusInternalServerError, stacktrace.Propagate(err, "cannot marshal firebase auth payload")
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/%s?key=%s", endpoint, url.QueryEscape(apiKey)), bytes.NewReader(body))
	if err != nil {
		return nil, fiber.StatusInternalServerError, stacktrace.Propagate(err, "cannot create firebase auth request")
	}
	request.Header.Set("Content-Type", "application/json")
	setFirebaseBrowserHeaders(request)

	return h.doFirebaseRequest(request)
}

func (h *FirebaseEmailAuthHandler) callSecureToken(form map[string]string) (*responses.FirebaseEmailAuthResponse, int, error) {
	apiKey := strings.TrimSpace(os.Getenv("FIREBASE_API_KEY"))
	if apiKey == "" {
		return nil, fiber.StatusServiceUnavailable, stacktrace.NewError("FIREBASE_API_KEY is not configured")
	}

	values := url.Values{}
	for key, value := range form {
		values.Set(key, value)
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://securetoken.googleapis.com/v1/token?key=%s", url.QueryEscape(apiKey)), strings.NewReader(values.Encode()))
	if err != nil {
		return nil, fiber.StatusInternalServerError, stacktrace.Propagate(err, "cannot create firebase token refresh request")
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	setFirebaseBrowserHeaders(request)

	return h.doFirebaseRequest(request)
}

func (h *FirebaseEmailAuthHandler) doFirebaseRequest(request *http.Request) (*responses.FirebaseEmailAuthResponse, int, error) {
	response, err := h.client.Do(request)
	if err != nil {
		if strings.TrimSpace(os.Getenv("FIREBASE_AUTH_PROXY_URL")) == "" {
			return nil, fiber.StatusBadGateway, stacktrace.Propagate(err, "cannot reach firebase authentication service")
		}

		retryRequest, retryErr := cloneFirebaseRequest(request)
		if retryErr != nil {
			return nil, fiber.StatusBadGateway, stacktrace.Propagate(err, "cannot reach firebase authentication service")
		}
		response, err = (&http.Client{Timeout: 20 * time.Second}).Do(retryRequest)
		if err != nil {
			return nil, fiber.StatusBadGateway, stacktrace.Propagate(err, "cannot reach firebase authentication service")
		}
	}
	defer response.Body.Close()

	data, err := io.ReadAll(io.LimitReader(response.Body, 1<<20))
	if err != nil {
		return nil, fiber.StatusBadGateway, stacktrace.Propagate(err, "cannot read firebase authentication response")
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, response.StatusCode, stacktrace.NewError("firebase authentication rejected the request: %s", firebaseSafeError(data))
	}

	var firebaseResponse struct {
		IDToken      string `json:"idToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
		LocalID      string `json:"localId"`
		Email        string `json:"email"`
	}
	if err := json.Unmarshal(data, &firebaseResponse); err != nil {
		return nil, fiber.StatusBadGateway, stacktrace.Propagate(err, "cannot decode firebase authentication response")
	}

	return &responses.FirebaseEmailAuthResponse{
		IDToken:      firebaseResponse.IDToken,
		RefreshToken: firebaseResponse.RefreshToken,
		ExpiresIn:    firebaseResponse.ExpiresIn,
		LocalID:      firebaseResponse.LocalID,
		Email:        firebaseResponse.Email,
	}, response.StatusCode, nil
}

func cloneFirebaseRequest(request *http.Request) (*http.Request, error) {
	if request.GetBody == nil {
		return nil, stacktrace.NewError("firebase auth request body cannot be replayed")
	}
	body, err := request.GetBody()
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot replay firebase auth request body")
	}
	clone := request.Clone(request.Context())
	clone.Body = body
	return clone, nil
}

func setFirebaseBrowserHeaders(request *http.Request) {
	referer := strings.TrimSpace(os.Getenv("FIREBASE_AUTH_REFERER_URL"))
	if referer == "" {
		referer = strings.TrimSpace(os.Getenv("APP_URL"))
	}
	if referer == "" {
		referer = "https://sms.nasak.ir"
	}
	if !strings.HasSuffix(referer, "/") {
		referer += "/"
	}
	request.Header.Set("Origin", strings.TrimRight(referer, "/"))
	request.Header.Set("Referer", referer)
	request.Header.Set("User-Agent", "NasakSMS/1.0")
}

func firebaseSafeError(data []byte) string {
	var payload struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(data, &payload); err != nil || payload.Error.Message == "" {
		return "unknown firebase error"
	}
	return payload.Error.Message
}

func (h *FirebaseEmailAuthHandler) firebaseError(c fiber.Ctx, status int, err error) error {
	if status == fiber.StatusServiceUnavailable || status == fiber.StatusBadGateway {
		return c.Status(status).JSON(fiber.Map{
			"status":  "error",
			"message": "Authentication service is temporarily unavailable.",
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  "error",
		"message": firebaseClientMessage(err.Error()),
	})
}

func firebaseClientMessage(message string) string {
	switch {
	case strings.Contains(message, "EMAIL_NOT_FOUND"):
		return "No account found with this email address."
	case strings.Contains(message, "INVALID_PASSWORD"), strings.Contains(message, "INVALID_LOGIN_CREDENTIALS"):
		return "Invalid email or password."
	case strings.Contains(message, "EMAIL_EXISTS"):
		return "An account already exists with this email."
	case strings.Contains(message, "WEAK_PASSWORD"):
		return "Password should be at least 6 characters."
	case strings.Contains(message, "USER_DISABLED"):
		return "This account has been disabled."
	case strings.Contains(message, "TOO_MANY_ATTEMPTS_TRY_LATER"):
		return "Too many failed attempts. Please try again later."
	default:
		return "Authentication failed. Please check your email and password."
	}
}

func firebaseAuthTransport(logger telemetry.Logger) http.RoundTripper {
	proxyURL := strings.TrimSpace(os.Getenv("FIREBASE_AUTH_PROXY_URL"))
	if proxyURL == "" {
		return http.DefaultTransport
	}

	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		logger.Warn(stacktrace.Propagate(err, "invalid FIREBASE_AUTH_PROXY_URL"))
		return http.DefaultTransport
	}

	switch parsedURL.Scheme {
	case "http", "https":
		return &http.Transport{Proxy: http.ProxyURL(parsedURL)}
	case "socks5", "socks5h":
		dialer, err := proxy.FromURL(parsedURL, proxy.Direct)
		if err != nil {
			logger.Warn(stacktrace.Propagate(err, "cannot configure firebase auth SOCKS proxy"))
			return http.DefaultTransport
		}
		return &http.Transport{
			DialContext: func(ctx context.Context, network string, address string) (net.Conn, error) {
				type result struct {
					conn net.Conn
					err  error
				}
				done := make(chan result, 1)
				go func() {
					conn, err := dialer.Dial(network, address)
					done <- result{conn: conn, err: err}
				}()

				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				case result := <-done:
					return result.conn, result.err
				}
			},
		}
	default:
		logger.Warn(stacktrace.NewError("unsupported FIREBASE_AUTH_PROXY_URL scheme [%s]", parsedURL.Scheme))
		return http.DefaultTransport
	}
}
