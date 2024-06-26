package flag

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
	mockfunction_fe.APIFlag()
	m.Run()
}

var testFlag models.Flag
var testFlagList []models.Flag

func TestFlagCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(FlagCmd)
	assert.Contains(t, output, "Manage your flags")
}

func TestFlagHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(FlagCmd, "--help")
	assert.Contains(t, output, "Manage your flags")
}

func TestFlagGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(FlagCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(FlagCmd, "get", "--id=testFlagID")

	err := json.Unmarshal([]byte(successOutput), &testFlag)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestFlag, testFlag)
}

func TestFlagListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(FlagCmd, "list")

	err := json.Unmarshal([]byte(output), &testFlagList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestFlagList, testFlagList)
}

func TestFlagCreateCommand(t *testing.T) {

	successOutput, _ := utils.ExecuteCommand(FlagCmd, "create", "--data-raw='{\"name\":\"testFlagName\",\"type\":\"string\",\"description\":\"testFlagDescription\",\"source\":\"cli\"}'")

	err := json.Unmarshal([]byte(successOutput), &testFlag)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestFlag, testFlag)
}

func TestFlagEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(FlagCmd, "edit")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\", \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(FlagCmd, "edit", "--id=testFlagID", "--data-raw='{\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"cli\"}'")

	err := json.Unmarshal([]byte(successOutput), &testFlag)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestFlagEdit, testFlag)
}

func TestFlagDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(FlagCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(FlagCmd, "delete", "--id=testFlagID")
	assert.Equal(t, "Flag deleted\n", successOutput)
}
