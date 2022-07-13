package models

type AuthenticationResponse struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type AuthenticationRequest struct {
	Grant_type    string `json:"grant_type"`
	Scope         string `json:"scope"`
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
}
