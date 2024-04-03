package campaign

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
	mockfunction_fe.APICampaign()

	m.Run()
}

var testCampaign models.CampaignFE
var testCampaignList []models.CampaignFE

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

	successOutput, _ := utils.ExecuteCommand(CampaignCmd, "get", "--id=testCampaignID")

	err := json.Unmarshal([]byte(successOutput), &testCampaign)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestCampaign, testCampaign)
}

func TestCampaignListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(CampaignCmd, "list")

	err := json.Unmarshal([]byte(output), &testCampaignList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestCampaignlist, testCampaignList)
}

func TestCampaignCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(CampaignCmd, "create")
	assert.Contains(t, "Error: required flag(s) \"data-raw\" not set\nUsage:\n  campaign create [-d <data-raw> | --data-raw=<data-raw>] [flags]\n\nFlags:\n  -d, --data-raw string   raw data contains all the info to create your campaign, check the doc for details\n  -h, --help              help for create\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(CampaignCmd, "create", "--data-raw='{\"project_id\":\"testProjectID\",\"name\":\"testCampaignName\",\"description\":\"testCampaignDescription\",\"type\":\"toggle\",\"status\":\"\",\"variation_groups\":[{\"id\":\"\",\"name\":\"variationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"blue\"}}},{\"id\":\"\",\"name\":\"My variation 2\",\"reference\":false,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}'")

	err := json.Unmarshal([]byte(successOutput), &testCampaign)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestCampaign, testCampaign)
}

func TestCampaignEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(CampaignCmd, "edit")
	assert.Contains(t, "Error: required flag(s) \"data-raw\", \"id\" not set\nUsage:\n  campaign edit [-i <campaign-id> | --id=<campaign-id>] [ -d <data-raw> | --data-raw=<data-raw>] [flags]\n\nFlags:\n  -d, --data-raw string   raw data contains all the info to edit your campaign, check the doc for details\n  -h, --help              help for edit\n  -i, --id string         id of the campaign you want to edit\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(CampaignCmd, "edit", "--id=testCampaignID", "--data-raw={\"project_id\":\"testProjectID1\",\"name\":\"testCampaignName1\",\"description\":\"testCampaignDescription1\",\"type\":\"toggle\",\"status\":\"\",\"variation_groups\":[{\"id\":\"\",\"name\":\"variationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"blue\"}}},{\"id\":\"\",\"name\":\"My variation 2\",\"reference\":false,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}")

	err := json.Unmarshal([]byte(successOutput), &testCampaign)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestCampaignEdit, testCampaign)
}

func TestCampaignDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(CampaignCmd, "delete")
	assert.Contains(t, "Error: required flag(s) \"id\" not set\nUsage:\n  campaign delete [-i <campaign-id> | --id=<campaign-id>] [flags]\n\nFlags:\n  -h, --help        help for delete\n  -i, --id string   id of the campaign you want to delete\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(CampaignCmd, "delete", "--id=testCampaignID")
	assert.Equal(t, "Campaign deleted\n", successOutput)
}

func TestCampaignSwitchCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(CampaignCmd, "switch")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\", \"status\" not set")

	failOutput1, _ := utils.ExecuteCommand(CampaignCmd, "switch", "--id=testCampaignID", "--status=notKnown")
	assert.Equal(t, "Status can only have 3 values: active or paused or interrupted\n", failOutput1)

	successOutput, _ := utils.ExecuteCommand(CampaignCmd, "switch", "--id=testCampaignID", "--status=active")
	assert.Equal(t, "campaign status set to active\n", successOutput)
}
