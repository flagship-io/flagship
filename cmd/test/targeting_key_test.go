package test

import (
	"encoding/json"
	"testing"

	"github.com/flagship-io/flagship/cmd/targeting_key"
	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/stretchr/testify/assert"
)

var testTargetingKey models.TargetingKey
var testTargetingKeyList []models.TargetingKey

func TestTargetingKeyCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(targeting_key.TargetingKeyCmd)
	assert.Contains(t, output, "Manage your targeting keys in your account")
}

func TestTargetingKeyHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(targeting_key.TargetingKeyCmd, "--help")
	assert.Contains(t, output, "Manage your targeting keys in your account")
}

func TestTargetingKeyGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(targeting_key.TargetingKeyCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(targeting_key.TargetingKeyCmd, "get", "--id=testTargetingKeyID")

	err := json.Unmarshal([]byte(successOutput), &testTargetingKey)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestTargetingKey, testTargetingKey)
}

func TestTargetingKeyListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(targeting_key.TargetingKeyCmd, "list")

	err := json.Unmarshal([]byte(output), &testTargetingKeyList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestTargetingKeyList, testTargetingKeyList)
}

func TestTargetingKeyCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(targeting_key.TargetingKeyCmd, "create")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\" not set")

	successOutput, _ := utils.ExecuteCommand(targeting_key.TargetingKeyCmd, "create", "--data-raw='{\"name\":\"testTargetingKeyName\",\"type\":\"string\",\"description\":\"testTargetingKeyDescription\"}'")

	err := json.Unmarshal([]byte(successOutput), &testTargetingKey)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestTargetingKey, testTargetingKey)
}

func TestTargetingKeyEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(targeting_key.TargetingKeyCmd, "edit")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\", \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(targeting_key.TargetingKeyCmd, "edit", "--id=testTargetingKeyID", "--data-raw='{\"name\":\"testTargetingKeyName1\",\"type\":\"string\",\"description\":\"testTargetingKeyDescription1\"}'")

	err := json.Unmarshal([]byte(successOutput), &testTargetingKey)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestTargetingKeyEdit, testTargetingKey)
}

func TestTargetingKeyDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(targeting_key.TargetingKeyCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(targeting_key.TargetingKeyCmd, "delete", "--id=testTargetingKeyID")
	assert.Equal(t, "Targeting key deleted\n", successOutput)
}
