package account_global_code

import (
	"encoding/json"
	"testing"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	mockfunction_we "github.com/flagship-io/flagship/utils/mock_function/web_experimentation"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	defer mockfunction_we.InitMockAuth()

	mockfunction.SetMock(&http_request.ResourceRequester)

	mockfunction_we.APIAccount()

	m.Run()
}

func TestAccountGlobalCodeCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(AccountGlobalCodeCmd)
	assert.Contains(t, output, "Get account global code")
}

func TestAccountGlobalCodeHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(AccountGlobalCodeCmd, "--help")
	assert.Contains(t, output, "Get account global code")
}

func TestAccountGlobalCodeGetCommand(t *testing.T) {
	failOutput, _ := utils.ExecuteCommand(AccountGlobalCodeCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(AccountGlobalCodeCmd, "get", "-i=account_id")
	assert.Equal(t, "console.log(\"test\")\n", successOutput)
}

func TestAccountGlobalCodePushCommand(t *testing.T) {
	var testAccount models.AccountWE

	failOutput, _ := utils.ExecuteCommand(AccountGlobalCodeCmd, "push")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(AccountGlobalCodeCmd, "push", "-i=account_id", "--code=console.log(\"test\")")
	err := json.Unmarshal([]byte(successOutput), &testAccount)

	assert.Nil(t, err)
	assert.Equal(t, mockfunction_we.TestAccountGlobalCode, testAccount)
}
