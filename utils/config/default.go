package config

import "github.com/flagship-io/flagship/cmd/version"

const (
	OutputFormat         = "table"
	GrantType            = "client_credentials"
	Expiration           = 86400
	Scope                = "*"
	ClientID             = "clientID"
	ClientSecret         = "clientSecret"
	Token                = "token"
	AccountID            = "accountID"
	AccountEnvironmentID = "accountEnvironmentID"
)

var DefaultUserAgent = "flagship-cli/" + version.Version
