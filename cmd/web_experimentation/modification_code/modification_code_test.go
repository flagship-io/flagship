/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/

package modification_code

import (
	"testing"

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
	assert.Contains(t, failOutput, "Error: required flag(s) \"campaign-id\", \"id\" not set\nUsage")

	successOutput, _ := utils.ExecuteCommand(ModificationCodeCmd, "get", "-i=120003", "--campaign-id=100000")
	assert.Equal(t, "console.log(\"test modification\")\n", successOutput)
}
