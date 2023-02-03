package utils

import "os"

func GetHost() string {
	if os.Getenv("STAGING") == "true" {
		return "https://staging-api.flagship.io"
	}

	return "https://api.flagship.io"
}

func GetHostAuth() string {
	if os.Getenv("STAGING") == "true" {
		return "https://staging-auth.flagship.io"
	}

	return "https://auth.flagship.io"
}
