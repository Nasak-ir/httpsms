package authlocal

import (
	"time"

	"github.com/NdoleStudio/httpsms/pkg/entities"
	"github.com/golang-jwt/jwt/v5"
	"github.com/palantir/stacktrace"
)

const issuer = "nasak-sms-local-auth"

// Sign creates a local API bearer token for environments where Firebase auth is unreachable.
func Sign(secret string, authUser entities.AuthContext, ttl time.Duration) (string, error) {
	if secret == "" {
		return "", stacktrace.NewError("LOCAL_AUTH_SECRET is not configured")
	}

	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"iss":   issuer,
		"sub":   authUser.ID.String(),
		"email": authUser.Email,
		"iat":   now.Unix(),
		"exp":   now.Add(ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Verify validates a local bearer token and returns its auth context.
func Verify(secret string, rawToken string) (entities.AuthContext, error) {
	if secret == "" {
		return entities.AuthContext{}, stacktrace.NewError("LOCAL_AUTH_SECRET is not configured")
	}

	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, stacktrace.NewError("unexpected local auth signing method")
		}
		return []byte(secret), nil
	}, jwt.WithIssuer(issuer))
	if err != nil || !token.Valid {
		return entities.AuthContext{}, stacktrace.Propagate(err, "invalid local auth token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return entities.AuthContext{}, stacktrace.NewError("invalid local auth claims")
	}

	userID, _ := claims.GetSubject()
	email, _ := claims["email"].(string)
	authUser := entities.AuthContext{
		ID:    entities.UserID(userID),
		Email: email,
	}
	if authUser.IsNoop() {
		return entities.AuthContext{}, stacktrace.NewError("invalid local auth context")
	}
	return authUser, nil
}
