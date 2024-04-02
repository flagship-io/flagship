package models

type TokenWE struct {
	ClientID  string `json:"client_id"`
	AccountID string `json:"account"`
	ExpiresIn int    `json:"expires_in"`
	Scope     string `json:"scope"`
}

type TokenFE struct {
	ClientID  string `json:"client_id"`
	AccountID string `json:"account"`
	ExpiresIn int    `json:"expires_in"`
	Scope     string `json:"scope"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ClientCredentialsRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthorizationCodeRequest struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	RefreshToken string `json:"refresh_token"`
}
