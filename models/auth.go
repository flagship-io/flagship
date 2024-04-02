package models

type Auth struct {
	Username     string `json:"username"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthYaml struct {
	Username     string `yaml:"username"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

type AccountYaml struct {
	CurrentUsedCredential string `yaml:"current_used_credential"`
	AccountID             string `yaml:"account_id"`
	AccountEnvironmentID  string `yaml:"account_environment_id"`
}

type AccountJSON struct {
	CurrentUsedCredential string `json:"current_used_credential"`
	AccountID             string `json:"account_id"`
	AccountEnvironmentID  string `json:"account_environment_id"`
}
