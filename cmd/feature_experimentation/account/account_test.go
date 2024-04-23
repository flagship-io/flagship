package account

import (
	"encoding/json"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/flagship-io/flagship/utils/mock_function/feature_experimentation"
	mockfunction_fe "github.com/flagship-io/flagship/utils/mock_function/feature_experimentation"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	defer mockfunction_fe.InitMockAuth()

	mockfunction.SetMock(&http_request.ResourceRequester)

	mockfunction_fe.APIToken()

	m.Run()
}

var testAccount models.AccountJSON

func TestAccountCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(AccountCmd)
	assert.Contains(t, output, "Manage your CLI authentication\n\nUsage:\n  account [use|current]")
}

func TestAccountHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(AccountCmd, "--help")
	assert.Contains(t, output, "Manage your CLI authentication\n\nUsage:\n  account [use|current]")
}

func TestAccountUseCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(AccountCmd, "use")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(AccountCmd, "use", "-i=account_id")
	assert.Equal(t, "Account ID set to : account_id\n", successOutput)
}

func TestAccountCurrentCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(AccountCmd, "current")

	err := json.Unmarshal([]byte(output), &testAccount)

	assert.Nil(t, err)

	assert.Equal(t, feature_experimentation.TestAccount.AccountID, testAccount.AccountID)
}
