package objects

type V1AuthenticationObjectRequest struct {
	Code string `json:"code"`
}

type V1AuthenticationObjectResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresAt    int    `json:"expires_at"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
