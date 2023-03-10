package utils

import (
	"os"

	"github.com/spf13/viper"
)

func GetHost() string {
	if os.Getenv("STAGING") == "true" {
		return "https://staging-api.flagship.io"
	}

	return "https://api.flagship.io"
}

func GetDecisionAPIHost() string {
	if os.Getenv("STAGING") == "true" {
		return "https://staging-decision.flagship.io"
	}

	if viper.GetString("self-hosted") != "" {
		return viper.GetString("self-hosted")
	}

	return "https://decision.flagship.io"
}

func GetHostAuth() string {
	if os.Getenv("STAGING") == "true" {
		return "https://staging-auth.flagship.io"
	}

	return "https://auth.flagship.io"
}
