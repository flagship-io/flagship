package flag

import (
	"encoding/json"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.APIFlag()
	m.Run()
}

var testFlag models.Flag
var testFlagList []models.Flag

func TestFlagCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(FlagCmd)
	assert.Contains(t, output, "Manage your flags in your account")
}

func TestFlagHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(FlagCmd, "--help")
	assert.Contains(t, output, "Manage your flags in your account")
}

func TestFlagGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(FlagCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(FlagCmd, "get", "--id=testFlagID")

	err := json.Unmarshal([]byte(successOutput), &testFlag)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestFlag, testFlag)
}

func TestFlagListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(FlagCmd, "list")

	err := json.Unmarshal([]byte(output), &testFlagList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestFlagList, testFlagList)
}

func TestFlagCreateCommand(t *testing.T) {

	successOutput, _ := utils.ExecuteCommand(FlagCmd, "create", "--data-raw='{\"name\":\"testFlagName\",\"type\":\"string\",\"description\":\"testFlagDescription\",\"source\":\"cli\"}'")

	err := json.Unmarshal([]byte(successOutput), &testFlag)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestFlag, testFlag)
}

func TestFlagEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(FlagCmd, "edit")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\", \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(FlagCmd, "edit", "--id=testFlagID", "--data-raw='{\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"cli\"}'")

	err := json.Unmarshal([]byte(successOutput), &testFlag)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestFlagEdit, testFlag)
}

func TestFlagDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(FlagCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(FlagCmd, "delete", "--id=testFlagID")
	assert.Equal(t, "Flag deleted\n", successOutput)
}
