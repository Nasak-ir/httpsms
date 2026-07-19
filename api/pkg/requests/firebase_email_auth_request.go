package requests

import "strings"

// FirebaseEmailAuthRequest contains email/password auth credentials.
type FirebaseEmailAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Mode     string `json:"mode"`
}

// Sanitize trims user-controlled auth fields without logging secrets.
func (request FirebaseEmailAuthRequest) Sanitize() FirebaseEmailAuthRequest {
	request.Email = strings.ToLower(strings.TrimSpace(request.Email))
	request.Mode = strings.ToLower(strings.TrimSpace(request.Mode))
	return request
}

// FirebaseEmailRefreshRequest contains a Firebase refresh token.
type FirebaseEmailRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
