package configuration

import (
	"encoding/json"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.InitMockConfiguration()

	m.Run()
}

var testConfiguration models.Configuration
var testConfigurationList []models.Configuration

func TestConfigurationCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(ConfigurationCmd)
	assert.Contains(t, output, "Manage your CLI configurations in your account")
}

func TestConfigurationHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(ConfigurationCmd, "--help")
	assert.Contains(t, output, "Manage your CLI configurations in your account")
}

func TestConfigurationCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(ConfigurationCmd, "create")

	assert.Contains(t, failOutput, "Configuration not created")

	successOutput, _ := utils.ExecuteCommand(ConfigurationCmd, "create", "--name=test_configuration", "-i=testConfigurationClientID", "-s=testConfigurationClientSecret", "-a=testConfigurationAccountID", "-e=testConfigurationAccountEnvID")
	assert.Equal(t, "Configuration created successfully\n", successOutput)

}

func TestConfigurationDeleteCommand(t *testing.T) {
	config.CreateConfigurationFile(mockfunction.TestConfiguration.Name, mockfunction.TestConfiguration.ClientID, mockfunction.TestConfiguration.ClientSecret, mockfunction.TestConfiguration.AccountID, mockfunction.TestConfiguration.AccountEnvironmentID)

	failOutput, _ := utils.ExecuteCommand(ConfigurationCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"name\" not set")

	successOutput, _ := utils.ExecuteCommand(ConfigurationCmd, "delete", "--name=test_configuration")
	assert.Equal(t, "Configuration deleted successfully\n", successOutput)
}

func TestConfigurationListCommand(t *testing.T) {
	config.CreateConfigurationFile(mockfunction.TestConfiguration.Name, mockfunction.TestConfiguration.ClientID, mockfunction.TestConfiguration.ClientSecret, mockfunction.TestConfiguration.AccountID, mockfunction.TestConfiguration.AccountEnvironmentID)

	output, _ := utils.ExecuteCommand(ConfigurationCmd, "list")

	err := json.Unmarshal([]byte(output), &testConfigurationList)

	byt, err := json.Marshal(mockfunction.TestConfiguration)

	assert.Nil(t, err)

	assert.Contains(t, output, string(byt))
}

func TestConfigurationGetCommand(t *testing.T) {
	config.CreateConfigurationFile(mockfunction.TestConfiguration.Name, mockfunction.TestConfiguration.ClientID, mockfunction.TestConfiguration.ClientSecret, mockfunction.TestConfiguration.AccountID, mockfunction.TestConfiguration.AccountEnvironmentID)

	failOutput, _ := utils.ExecuteCommand(ConfigurationCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"name\" not set")

	successOutput, _ := utils.ExecuteCommand(ConfigurationCmd, "get", "--name=test_configuration")
	err := json.Unmarshal([]byte(successOutput), &testConfiguration)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestConfiguration, testConfiguration)
}
