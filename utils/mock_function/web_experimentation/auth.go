package web_experimentation

import (
	"os"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
)

var TestAuth = models.Auth{
	Username:     "test_auth",
	ClientID:     "",
	ClientSecret: "",
	Token:        "testAccessToken",
	RefreshToken: "testRefreshToken",
}

func InitMockAuth() {
	os.Remove(config.CredentialPath(utils.WEB_EXPERIMENTATION, "test_auth"))
}
