package variation

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

	mockfunction.SetMock(&http_request.ResourceRequester)
	mockfunction_we.APIVariation()

	m.Run()
}

var testVariation models.VariationWE

func TestVariationCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(VariationCmd)
	assert.Contains(t, output, "Error: required flag(s) \"campaign-id\" not set")

	output1, _ := utils.ExecuteCommand(VariationCmd, "--campaign-id=100000")
	assert.Contains(t, output1, "Manage your variation")

}

func TestVariationHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(VariationCmd, "--help")
	assert.Contains(t, output, "Manage your variation")
}

func TestVariationGetCommand(t *testing.T) {

	successOutput, _ := utils.ExecuteCommand(VariationCmd, "get", "-i=110000", "--campaign-id=100000")

	err := json.Unmarshal([]byte(successOutput), &testVariation)

	assert.Nil(t, err)
	assert.Equal(t, mockfunction_we.TestVariation, testVariation)
}

func TestVariationDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(VariationCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(VariationCmd, "delete", "--campaign-id=100000", "--id=110000")
	assert.Equal(t, "Variation deleted\n", successOutput)
}
