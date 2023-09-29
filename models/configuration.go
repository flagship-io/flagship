package models

type Configuration struct {
	Name                 string `json:"name"`
	ClientID             string `json:"client_id"`
	ClientSecret         string `json:"client_secret"`
	AccountID            string `json:"account_id"`
	AccountEnvironmentID string `json:"account_environment_id"`
}

type ConfigurationYaml struct {
	Name                 string `yaml:"name"`
	ClientID             string `yaml:"client_id"`
	ClientSecret         string `yaml:"client_secret"`
	AccountID            string `yaml:"account_id"`
	AccountEnvironmentID string `yaml:"account_environment_id"`
}
