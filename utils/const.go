package utils

import "os"

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
		return "https://api-auth.abtasty.com"
	}

	return "https://api-auth.abtasty.com"
}

const FEATURE_EXPERIMENTATION = "fe"
const WEB_EXPERIMENTATION = "we"
const HOME_CLI = ".cli"
const CLIENT_ID = "590_397jchm4asqo04g0ogcggcoo88wo44sg0c0owk0448gkwck8s8"
const CLIENT_SECRET = "5xc41rk3etwc404o8scs0s8c0sogcgc48g4sk0ggc44so4w00c"
