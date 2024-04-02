package auth

import (
	"encoding/json"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
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

var testConfiguration models.Auth
var testConfigurationList []models.Auth

func TestAuthCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(AuthCmd)
	assert.Contains(t, output, "Manage your CLI authentication for feature experimentation\n\nUsage:\n  authentication [login|get|list]")
}

func TestAuthHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(AuthCmd, "--help")
	assert.Contains(t, output, "Manage your CLI authentication for feature experimentation\n\nUsage:\n  authentication [login|get|list]")
}

func TestAuthLoginCommand(t *testing.T) {
	successOutput, _ := utils.ExecuteCommand(AuthCmd, "login", "-u=test_configuration", "-i=testConfigurationClientID", "-s=testConfigurationClientSecret", "-a=account_id")
	assert.Equal(t, "Credential created successfully\n", successOutput)
}

func TestAuthListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(AuthCmd, "list")

	err := json.Unmarshal([]byte(output), &testConfigurationList)

	assert.Nil(t, err)

	//byt, err := json.Marshal(mockfunction.TestConfiguration)

	assert.Nil(t, err)

	//assert.Contains(t, output, string(byt))
}

func TestAuthGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(AuthCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"username\" not set")

	successOutput, _ := utils.ExecuteCommand(AuthCmd, "get", "--username=test_configuration")
	err := json.Unmarshal([]byte(successOutput), &testConfiguration)

	assert.Nil(t, err)

	//assert.Equal(t, mockfunction.TestConfiguration, testConfiguration)
}
