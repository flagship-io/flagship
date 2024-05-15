/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/

package modification_code

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

	mockfunction_we.APIModification()

	m.Run()
}

func TestModificationCodeCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(ModificationCodeCmd)
	assert.Contains(t, output, "Get modification code")
}

func TestModificationCodeHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(ModificationCodeCmd, "--help")
	assert.Contains(t, output, "Get modification code")
}

func TestModificationCodeGetCommand(t *testing.T) {
	failOutput, _ := utils.ExecuteCommand(ModificationCodeCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"campaign-id\", \"id\" not set\n")

	successOutput, _ := utils.ExecuteCommand(ModificationCodeCmd, "get", "-i=120003", "--campaign-id=100000")
	assert.Equal(t, "console.log(\"test modification\")\n", successOutput)
}

func TestModificationCodePushCommand(t *testing.T) {
	var testModification models.Modification

	failOutput, _ := utils.ExecuteCommand(ModificationCodeCmd, "push")
	assert.Contains(t, failOutput, "Error: required flag(s) \"campaign-id\" not set")

	successOutput, _ := utils.ExecuteCommand(ModificationCodeCmd, "push", "-i=120003", "--campaign-id=100000", "--code=console.log(\"test modification\")")
	err := json.Unmarshal([]byte(successOutput), &testModification)

	assert.Nil(t, err)
	assert.Equal(t, mockfunction_we.TestElementModification, testModification)
}
