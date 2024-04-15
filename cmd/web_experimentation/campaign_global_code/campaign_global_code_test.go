package campaign_global_code

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

	mockfunction_we.APICampaign()

	m.Run()
}

func TestCampaignGlobalCodeCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(CampaignGlobalCodeCmd)
	assert.Contains(t, output, "Get campaign global code")
}

func TestCampaignGlobalCodeHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(CampaignGlobalCodeCmd, "--help")
	assert.Contains(t, output, "Get campaign global code")
}

func TestCampaignGlobalCodeGetCommand(t *testing.T) {
	failOutput, _ := utils.ExecuteCommand(CampaignGlobalCodeCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(CampaignGlobalCodeCmd, "get", "-i=100000")
	assert.Equal(t, "console.log(\"Hello World!\")\n", successOutput)
}
