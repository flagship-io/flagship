package models

type Token struct {
	ClientID  string `json:"client_id"`
	AccountID string `json:"account"`
	ExpiresIn int    `json:"expires_in"`
	Scope     string `json:"scope"`
}
