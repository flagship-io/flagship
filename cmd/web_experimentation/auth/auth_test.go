package auth

import (
	"encoding/json"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/flagship-io/flagship/utils/mock_function/web_experimentation"
	mockfunction_we "github.com/flagship-io/flagship/utils/mock_function/web_experimentation"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	//defer mockfunction_we.InitMockAuth()

	mockfunction.SetMock(&http_request.ResourceRequester)

	mockfunction_we.APIToken()

	m.Run()
}

var testAuth models.Auth
var testAuthList []models.Auth

func TestAuthCommand(t *testing.T) {
	output, err := utils.ExecuteCommand(AuthCmd)

	assert.Nil(t, err)
	assert.Contains(t, output, "Manage your CLI authentication for web experimentation\n\nUsage:\n  authentication [login|get|list|delete]")
}

func TestAuthHelpCommand(t *testing.T) {
	output, err := utils.ExecuteCommand(AuthCmd, "--help")

	assert.Nil(t, err)
	assert.Contains(t, output, "Manage your CLI authentication for web experimentation\n\nUsage:\n  authentication [login|get|list|delete]")
}

func TestAuthLoginCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(AuthCmd, "login", "-u=test_auth", "--client-id=CI", "--client-secret=CS")

	assert.Equal(t, "Error while login, required fields (username, client ID, client secret, account id)\n", failOutput)

	successOutput, err := utils.ExecuteCommand(AuthCmd, "login", "-u=test_auth", "--client-id=CI", "--client-secret=CS", "--account-id=AI")

	assert.Nil(t, err)
	assert.Equal(t, "Credential created successfully\n", successOutput)
}

func TestAuthListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(AuthCmd, "list")

	err := json.Unmarshal([]byte(output), &testAuthList)

	assert.Nil(t, err)

	byt, err := json.Marshal(web_experimentation.TestAuth)

	assert.Nil(t, err)

	assert.Contains(t, output, string(byt))
}

func TestAuthGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(AuthCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"username\" not set")

	successOutput, _ := utils.ExecuteCommand(AuthCmd, "get", "--username=test_auth")
	err := json.Unmarshal([]byte(successOutput), &testAuth)

	assert.Nil(t, err)

	assert.Equal(t, web_experimentation.TestAuth, testAuth)
}

func TestAuthDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(AuthCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"username\" not set")

	output, _ := utils.ExecuteCommand(AuthCmd, "delete", "--username=test_auth")

	assert.Contains(t, output, "Credential deleted successfully")
}
