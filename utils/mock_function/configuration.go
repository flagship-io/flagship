package mockfunction

import (
	"os"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils/config"
)

var TestConfiguration = models.Configuration{
	Name:                 "test_configuration",
	ClientID:             "testConfigurationClientID",
	ClientSecret:         "testConfigurationClientSecret",
	AccountID:            "testConfigurationAccountID",
	AccountEnvironmentID: "testConfigurationAccountEnvID",
}

func InitMockConfiguration() {
	config.SetViperMock()
	os.Remove(config.SetPathForConfigName("test_configuration"))
}
