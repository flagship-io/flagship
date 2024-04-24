package feature_experimentation

import (
	"log"
	"os"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
)

var TestAuth = models.Auth{
	Username:     "test_auth",
	ClientID:     "testAuthClientID",
	ClientSecret: "testAuthClientSecret",
	Token:        "testAccessToken",
	RefreshToken: "testRefreshToken",
}

func InitMockAuth() {
	credPath, err := config.CredentialPath(utils.FEATURE_EXPERIMENTATION, "test_auth")
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}

	os.Remove(credPath)
}
