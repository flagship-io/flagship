package variation_global_code

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

func TestVariationGlobalCodeCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(VariationGlobalCodeCmd)
	assert.Contains(t, output, "Get variation global code")
}

func TestVariationGlobalCodeHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(VariationGlobalCodeCmd, "--help")
	assert.Contains(t, output, "Get variation global code")
}

func TestVariationGlobalCodeGetJSCommand(t *testing.T) {
	failOutput, _ := utils.ExecuteCommand(VariationGlobalCodeCmd, "get-js")
	assert.Contains(t, failOutput, "Error: required flag(s) \"campaign-id\", \"id\" not set\nUsage")

	successOutput, _ := utils.ExecuteCommand(VariationGlobalCodeCmd, "get-js", "-i=110000", "--campaign-id=100000")
	assert.Equal(t, "console.log(\"test modification\")\n", successOutput)
}

func TestVariationGlobalCodeGetCSSCommand(t *testing.T) {
	failOutput, _ := utils.ExecuteCommand(VariationGlobalCodeCmd, "get-css")
	assert.Contains(t, failOutput, "Error: required flag(s) \"campaign-id\", \"id\" not set\nUsage")

	successOutput, _ := utils.ExecuteCommand(VariationGlobalCodeCmd, "get-css", "-i=110000", "--campaign-id=100000")
	assert.Equal(t, ".id{\"color\": \"black\"}\n", successOutput)
}
