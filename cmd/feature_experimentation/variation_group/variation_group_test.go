package variation_group

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
	mockfunction_fe.APIVariationGroup()

	m.Run()
}

var testVariationGroup models.VariationGroup
var testVariationGroupList []models.VariationGroup

func TestVariationGroupCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(VariationGroupCmd)
	assert.Contains(t, output, "Error: required flag(s) \"campaign-id\" not set")

	output1, _ := utils.ExecuteCommand(VariationGroupCmd, "--campaign-id=campaignID")
	assert.Contains(t, output1, "Manage your variation groups")
}

func TestVariationGroupHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(VariationGroupCmd, "--help")
	assert.Contains(t, output, "Manage your variation groups")
}

func TestVariationGroupGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(VariationGroupCmd, "get", "--campaign-id=campaignID")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(VariationGroupCmd, "get", "--id=testVariationGroupID", "--campaign-id=campaignID")

	err := json.Unmarshal([]byte(successOutput), &testVariationGroup)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestVariationGroup, testVariationGroup)
}

func TestVariationGroupListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(VariationGroupCmd, "list", "--campaign-id=campaignID")

	err := json.Unmarshal([]byte(output), &testVariationGroupList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestVariationGroupList, testVariationGroupList)
}

func TestVariationGroupCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(VariationGroupCmd, "create", "--campaign-id=campaignID")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\" not set")

	successOutput, _ := utils.ExecuteCommand(VariationGroupCmd, "create", "--campaign-id=campaignID", "--data-raw='{\"name\":\"testVariationGroupName\",\"variations\":[{\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}'")

	err := json.Unmarshal([]byte(successOutput), &testVariationGroup)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestVariationGroup, testVariationGroup)
}

func TestVariationGroupEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(VariationGroupCmd, "edit", "--campaign-id=campaignID")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\", \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(VariationGroupCmd, "edit", "--id=testVariationGroupID", "--campaign-id=campaignID", "--data-raw='{\"name\":\"testVariationGroupName1\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}'")

	err := json.Unmarshal([]byte(successOutput), &testVariationGroup)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestVariationGroupEdit, testVariationGroup)
}

func TestVariationGroupDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(VariationGroupCmd, "delete", "--campaign-id=campaignID")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(VariationGroupCmd, "delete", "--id=testVariationGroupID", "--campaign-id=campaignID")
	assert.Equal(t, "Variation group deleted\n", successOutput)
}
