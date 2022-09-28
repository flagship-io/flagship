package test

import (
	"encoding/json"
	"testing"

	"github.com/flagship-io/flagship/cmd/flag"
	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/stretchr/testify/assert"
)

var testFlag models.Flag
var testFlagList []models.Flag
var testFlagUsageList []models.FlagUsage

func TestFlagCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(flag.FlagCmd)
	assert.Contains(t, output, "Manage your flags in your account")
}

func TestFlagHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(flag.FlagCmd, "--help")
	assert.Contains(t, output, "Manage your flags in your account")
}

func TestFlagGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "get", "--id=testFlagID")

	err := json.Unmarshal([]byte(successOutput), &testFlag)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestFlag, testFlag)
}

func TestFlagListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(flag.FlagCmd, "list")

	err := json.Unmarshal([]byte(output), &testFlagList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestFlagList, testFlagList)
}

func TestFlagCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "create")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\" not set")

	successOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "create", "--data-raw='{\"name\":\"testFlagName\",\"type\":\"string\",\"description\":\"testFlagDescription\",\"source\":\"manual\"}'")

	err := json.Unmarshal([]byte(successOutput), &testFlag)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestFlag, testFlag)
}

func TestFlagEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "edit")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\", \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "edit", "--id=testFlagID", "--data-raw='{\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"manual\"}'")

	err := json.Unmarshal([]byte(successOutput), &testFlag)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestFlagEdit, testFlag)
}

func TestFlagDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "delete", "--id=testFlagID")
	assert.Equal(t, "Flag deleted\n", successOutput)
}

func TestFlagUsageListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(flag.FlagCmd, "usage", "get")

	err := json.Unmarshal([]byte(output), &testFlagUsageList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestFlagUsageList, testFlagUsageList)
}
