package feature_experimentation

import (
	"os"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
)

func InitMockAuth() {
	os.Remove(config.CredentialPath(utils.FEATURE_EXPERIMENTATION, "test_configuration"))
}
