package utils

import (
	"log"
	"os"
)

func GetFeatureExperimentationHost() string {
	if os.Getenv("FS_STAGING") == "true" {
		return "https://staging-api.flagship.io"
	}

	return "https://api.flagship.io"
}

func GetWebExperimentationHost() string {
	if os.Getenv("FS_STAGING") == "true" {
		return "https://staging-api.abtasty.com/api"
	}

	return "https://api.abtasty.com/api"
}

func GetHostFeatureExperimentationAuth() string {
	if os.Getenv("FS_STAGING") == "true" {
		return "https://staging-auth.flagship.io"
	}

	return "https://auth.flagship.io"
}

func GetHostWebExperimentationAuth() string {
	if os.Getenv("FS_STAGING") == "true" {
		return "https://staging-api-auth.abtasty.com"
	}

	return "https://api-auth.abtasty.com"
}

func DefaultGlobalCodeWorkingDir() string {
	wdDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}

	return wdDir
}

const FEATURE_EXPERIMENTATION = "fe"
const WEB_EXPERIMENTATION = "we"
const HOME_CLI = ".cli"
