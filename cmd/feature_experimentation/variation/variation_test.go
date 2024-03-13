package variation

import (
	"encoding/json"
	"testing"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function/feature_experimentation"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.APIVariation()

	m.Run()
}

var testVariation models.Variation
var testVariationList []models.Variation

func TestVariationCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(VariationCmd)
	assert.Contains(t, output, "Error: required flag(s) \"campaign-id\", \"variation-group-id\" not set")

	output1, _ := utils.ExecuteCommand(VariationCmd, "--campaign-id=campaignID", "--variation-group-id=variationGroupID")
	assert.Contains(t, output1, "Manage your variations in your variation group")
}

func TestVariationHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(VariationCmd, "--help")
	assert.Contains(t, output, "Manage your variations in your variation group")
}

func TestVariationGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(VariationCmd, "get", "--campaign-id=campaignID", "--variation-group-id=variationGroupID")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(VariationCmd, "get", "--id=testVariationID", "--campaign-id=campaignID", "--variation-group-id=variationGroupID")

	err := json.Unmarshal([]byte(successOutput), &testVariation)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestVariation, testVariation)
}

func TestVariationListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(VariationCmd, "list", "--campaign-id=campaignID", "--variation-group-id=variationGroupID")

	err := json.Unmarshal([]byte(output), &testVariationList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestVariationList, testVariationList)
}

func TestVariationCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(VariationCmd, "create", "--campaign-id=campaignID", "--variation-group-id=variationGroupID")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\" not set")

	successOutput, _ := utils.ExecuteCommand(VariationCmd, "create", "--campaign-id=campaignID", "--variation-group-id=variationGroupID", "--data-raw='{\"name\":\"testVariationName\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}'")

	err := json.Unmarshal([]byte(successOutput), &testVariation)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestVariation, testVariation)
}

func TestVariationEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(VariationCmd, "edit", "--campaign-id=campaignID", "--variation-group-id=variationGroupID")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\", \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(VariationCmd, "edit", "--id=testVariationID", "--campaign-id=campaignID", "--variation-group-id=variationGroupID", "--data-raw='{\"name\":\"testVariationName1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP1\"}}'")

	err := json.Unmarshal([]byte(successOutput), &testVariation)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestVariationEdit, testVariation)
}

func TestVariationDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(VariationCmd, "delete", "--campaign-id=campaignID", "--variation-group-id=variationGroupID")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(VariationCmd, "delete", "--id=testVariationID", "--campaign-id=campaignID", "--variation-group-id=variationGroupID")
	assert.Equal(t, "Variation deleted\n", successOutput)
}
