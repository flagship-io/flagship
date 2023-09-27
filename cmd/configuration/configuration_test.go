package configuration

import (
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

	successOutput, _ := utils.ExecuteCommand(ConfigurationCmd, "create", "--name=testConfigurationName1", "-i=testConfigurationClientID", "-s=testConfigurationClientSecret", "-a=testConfigurationAccountID", "-e=testConfigurationAccountEnvID")
	assert.Equal(t, "Configuration created successfully\n", successOutput)

}

func TestConfigurationDeleteCommand(t *testing.T) {
	config.CreateConfigurationFile(mockfunction.TestConfiguration.Name, mockfunction.TestConfiguration.ClientID, mockfunction.TestConfiguration.ClientSecret, mockfunction.TestConfiguration.AccountID, mockfunction.TestConfiguration.AccountEnvironmentID)

	failOutput, _ := utils.ExecuteCommand(ConfigurationCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"name\" not set")

	successOutput, _ := utils.ExecuteCommand(ConfigurationCmd, "delete", "--name=testConfigurationName")
	assert.Equal(t, "Configuration deleted successfully\n", successOutput)
}
