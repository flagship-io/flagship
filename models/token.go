package models

type Token struct {
	ClientID  string `json:"client_id"`
	AccountID string `json:"account"`
	ExpiresIn int    `json:"expires_in"`
	Scope     string `json:"scope"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
