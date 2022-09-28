package test

import (
	"testing"

	"github.com/flagship-io/flagship/cmd/campaign"
	"github.com/flagship-io/flagship/utils"
	"github.com/stretchr/testify/assert"
)

func TestCampaignCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(campaign.CampaignCmd)
	assert.Equal(t, "Manage your campaigns\n\nUsage:\n  campaign [create|edit|get|list|delete|toggle] [flags]\n  campaign [command]\n\nAvailable Commands:\n  completion  Generate the autocompletion script for the specified shell\n  create      Create a campaign\n  delete      Delete a campaign\n  edit        Edit a campaign\n  get         Get a campaign\n  help        Help about any command\n  list        List all campaigns\n  toggle      Toggle a campaign\n\nFlags:\n  -h, --help   help for campaign\n\nUse \"campaign [command] --help\" for more information about a command.\n", output)
}

func TestCampaignHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(campaign.CampaignCmd, "--help")
	assert.Equal(t, "Manage your campaigns\n\nUsage:\n  campaign [create|edit|get|list|delete|toggle] [flags]\n  campaign [command]\n\nAvailable Commands:\n  completion  Generate the autocompletion script for the specified shell\n  create      Create a campaign\n  delete      Delete a campaign\n  edit        Edit a campaign\n  get         Get a campaign\n  help        Help about any command\n  list        List all campaigns\n  toggle      Toggle a campaign\n\nFlags:\n  -h, --help   help for campaign\n\nUse \"campaign [command] --help\" for more information about a command.\n", output)
}

func TestCampaignGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(campaign.CampaignCmd, "get")
	assert.Equal(t, "Error: required flag(s) \"id\" not set\nUsage:\n  campaign get [-i <campaign-id> | --id=<campaign-id>] [flags]\n\nFlags:\n  -h, --help        help for get\n  -i, --id string   id of the campaign you want to display\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(campaign.CampaignCmd, "get", "--id=testCampaignID")
	assert.Equal(t, "{\"id\":\"testCampaignID\",\"project_id\":\"testProjectID\",\"name\":\"testCampaignName\",\"description\":\"testCampaignDescription\",\"type\":\"toggle\",\"status\":\"\",\"variation_groups\":[{\"id\":\"\",\"name\":\"variationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"blue\"}}},{\"id\":\"\",\"name\":\"My variation 2\",\"reference\":false,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}\n", successOutput)
}

func TestCampaignListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(campaign.CampaignCmd, "list")
	assert.Equal(t, "[{\"id\":\"testCampaignID\",\"project_id\":\"testProjectID\",\"name\":\"testCampaignName\",\"description\":\"testCampaignDescription\",\"type\":\"toggle\",\"status\":\"\",\"variation_groups\":[{\"id\":\"\",\"name\":\"variationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"blue\"}}},{\"id\":\"\",\"name\":\"My variation 2\",\"reference\":false,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}},{\"id\":\"testCampaignID1\",\"project_id\":\"testProjectID1\",\"name\":\"testCampaignName1\",\"description\":\"testCampaignDescription1\",\"type\":\"toggle\",\"status\":\"\",\"variation_groups\":[{\"id\":\"\",\"name\":\"variationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"blue\"}}},{\"id\":\"\",\"name\":\"My variation 2\",\"reference\":false,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}]\n", output)
}

func TestCampaignCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(campaign.CampaignCmd, "create")
	assert.Equal(t, "Error: required flag(s) \"data-raw\" not set\nUsage:\n  campaign create [-d <data-raw> | --data-raw=<data-raw>] [flags]\n\nFlags:\n  -d, --data-raw string   raw data contains all the info to create your campaign, check the doc for details\n  -h, --help              help for create\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(campaign.CampaignCmd, "create", "--data-raw='{\"project_id\":\"testProjectID\",\"name\":\"testCampaignName\",\"description\":\"testCampaignDescription\",\"type\":\"toggle\",\"status\":\"\",\"variation_groups\":[{\"id\":\"\",\"name\":\"variationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"blue\"}}},{\"id\":\"\",\"name\":\"My variation 2\",\"reference\":false,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}'")
	assert.Equal(t, "campaign created: {\"id\":\"testCampaignID\",\"project_id\":\"testProjectID\",\"name\":\"testCampaignName\",\"description\":\"testCampaignDescription\",\"type\":\"toggle\",\"status\":\"\",\"variation_groups\":[{\"id\":\"\",\"name\":\"variationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"blue\"}}},{\"id\":\"\",\"name\":\"My variation 2\",\"reference\":false,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}\n", successOutput)
}

func TestCampaignEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(campaign.CampaignCmd, "edit")
	assert.Equal(t, "Error: required flag(s) \"data-raw\", \"id\" not set\nUsage:\n  campaign edit [-i <campaign-id> | --id=<campaign-id>] [ -d <data-raw> | --data-raw=<data-raw>] [flags]\n\nFlags:\n  -d, --data-raw string   raw data contains all the info to edit your campaign, check the doc for details\n  -h, --help              help for edit\n  -i, --id string         id of the campaign you want to edit\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(campaign.CampaignCmd, "edit", "--id=testCampaignID", "--data-raw={\"project_id\":\"testProjectID1\",\"name\":\"testCampaignName1\",\"description\":\"testCampaignDescription1\",\"type\":\"toggle\",\"status\":\"\",\"variation_groups\":[{\"id\":\"\",\"name\":\"variationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"blue\"}}},{\"id\":\"\",\"name\":\"My variation 2\",\"reference\":false,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}")
	assert.Equal(t, "campaign updated: {\"id\":\"testCampaignID\",\"project_id\":\"testProjectID1\",\"name\":\"testCampaignName1\",\"description\":\"testCampaignDescription1\",\"type\":\"toggle\",\"status\":\"\",\"variation_groups\":[{\"id\":\"\",\"name\":\"variationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"blue\"}}},{\"id\":\"\",\"name\":\"My variation 2\",\"reference\":false,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}\n", successOutput)
}

func TestCampaignDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(campaign.CampaignCmd, "delete")
	assert.Equal(t, "Error: required flag(s) \"id\" not set\nUsage:\n  campaign delete [-i <campaign-id> | --id=<campaign-id>] [flags]\n\nFlags:\n  -h, --help        help for delete\n  -i, --id string   id of the campaign you want to delete\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(campaign.CampaignCmd, "delete", "--id=testCampaignID")
	assert.Equal(t, "Campaign deleted\n", successOutput)
}
