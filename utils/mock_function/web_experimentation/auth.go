package web_experimentation

import (
	"log"
	"os"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
)

var TestAuth = models.Auth{
	Username:     "test_auth",
	ClientID:     "CI",
	ClientSecret: "CS",
	Token:        "testAccessToken",
}

func InitMockAuth() {
	credPath, err := config.CredentialPath(utils.WEB_EXPERIMENTATION, "test_auth")
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}

	os.Remove(credPath)
}
