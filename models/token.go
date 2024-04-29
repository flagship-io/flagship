package models

type MfaRequestWE struct {
	MfaToken   string   `json:"mfa_token"`
	MfaMethods []string `json:"mfa_methods"`
}

type MultiFactorMethodRequestWE struct {
	MfaToken  string `json:"token"`
	MfaMethod string `json:"mfa_method"`
	GrantType string `json:"grant_type"`
}

type MultiFactorRequestWE struct {
	MfaToken  string `json:"token"`
	MfaMethod string `json:"mfa_method"`
	GrantType string `json:"grant_type"`
	Code      string `json:"code"`
}

type MultiFactorMethodResponseWE struct {
	MfaToken  string `json:"token"`
	MfaMethod string `json:"mfa_method"`
	GrantType string `json:"grant_type"`
}

type Token struct {
	ClientID  string `json:"client_id"`
	AccountID string `json:"account"`
	ExpiresIn int    `json:"expires_in"`
	Scope     string `json:"scope"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type ClientCredentialsRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope,omitempty"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthorizationCodeRequest struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type PasswordRequest struct {
	GrantType    string `json:"grant_type"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type RefreshTokenRequestFE struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	RefreshToken string `json:"refresh_token"`
}
type RefreshTokenRequestWE struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	RefreshToken string `json:"refresh_token"`
	ClientSecret string `json:"client_secret"`
}
