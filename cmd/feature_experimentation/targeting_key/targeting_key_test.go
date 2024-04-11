package targetingkey

import (
	"encoding/json"
	"testing"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
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

	mockfunction.SetMock(&http_request.ResourceRequester)
	mockfunction_fe.APITargetingKey()

	m.Run()
}

var testTargetingKey models.TargetingKey
var testTargetingKeyList []models.TargetingKey

func TestTargetingKeyCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(TargetingKeyCmd)
	assert.Contains(t, output, "Manage your targeting keys")
}

func TestTargetingKeyHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(TargetingKeyCmd, "--help")
	assert.Contains(t, output, "Manage your targeting keys")
}

func TestTargetingKeyGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(TargetingKeyCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(TargetingKeyCmd, "get", "--id=testTargetingKeyID")

	err := json.Unmarshal([]byte(successOutput), &testTargetingKey)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestTargetingKey, testTargetingKey)
}

func TestTargetingKeyListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(TargetingKeyCmd, "list")

	err := json.Unmarshal([]byte(output), &testTargetingKeyList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestTargetingKeyList, testTargetingKeyList)
}

func TestTargetingKeyCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(TargetingKeyCmd, "create")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\" not set")

	successOutput, _ := utils.ExecuteCommand(TargetingKeyCmd, "create", "--data-raw='{\"name\":\"testTargetingKeyName\",\"type\":\"string\",\"description\":\"testTargetingKeyDescription\"}'")

	err := json.Unmarshal([]byte(successOutput), &testTargetingKey)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestTargetingKey, testTargetingKey)
}

func TestTargetingKeyEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(TargetingKeyCmd, "edit")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\", \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(TargetingKeyCmd, "edit", "--id=testTargetingKeyID", "--data-raw='{\"name\":\"testTargetingKeyName1\",\"type\":\"string\",\"description\":\"testTargetingKeyDescription1\"}'")

	err := json.Unmarshal([]byte(successOutput), &testTargetingKey)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestTargetingKeyEdit, testTargetingKey)
}

func TestTargetingKeyDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(TargetingKeyCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(TargetingKeyCmd, "delete", "--id=testTargetingKeyID")
	assert.Equal(t, "Targeting key deleted\n", successOutput)
}
