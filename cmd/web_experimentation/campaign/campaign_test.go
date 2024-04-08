package campaign

import (
	"encoding/json"
	"strconv"
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
	mockfunction_we.APICampaign()

	m.Run()
}

var testCampaign models.CampaignWE
var testCampaignList []models.CampaignWE

func TestCampaignCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(CampaignCmd)
	assert.Contains(t, output, "Manage your campaigns")
}

func TestCampaignHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(CampaignCmd, "--help")
	assert.Contains(t, output, "Manage your campaigns")
}

func TestCampaignGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(CampaignCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(CampaignCmd, "get", "--id="+strconv.Itoa(100000))

	err := json.Unmarshal([]byte(successOutput), &testCampaign)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_we.TestCampaign, testCampaign)
}

func TestCampaignListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(CampaignCmd, "list")

	err := json.Unmarshal([]byte(output), &testCampaignList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_we.TestCampaignlist, testCampaignList)
}
