package responses

// FirebaseEmailAuthResponse is returned by the same-origin email auth proxy.
type FirebaseEmailAuthResponse struct {
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	LocalID      string `json:"local_id"`
	Email        string `json:"email"`
}
