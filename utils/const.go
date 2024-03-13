package utils

import "os"

func GetFeatureExperimentationHost() string {
	if os.Getenv("STAGING") == "true" {
		return "https://staging-api.flagship.io"
	}

	return "https://api.flagship.io"
}

func GetWebExperimentationHost() string {
	if os.Getenv("STAGING") == "true" {
		return "https://staging-api.abtasty.com/api"
	}

	return "https://api.abtasty.com/api"
}

func GetHostFeatureExperimentationAuth() string {
	if os.Getenv("STAGING") == "true" {
		return "https://staging-auth.flagship.io"
	}

	return "https://auth.flagship.io"
}
