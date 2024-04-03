package feature_experimentation

import (
	"testing"

	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/stretchr/testify/assert"
)

var variationGroupID = "variationGroupID"
var variationRequester = VariationFERequester{&common.ResourceRequest{AccountID: "account_id", AccountEnvironmentID: "account_environment_id"}}

func TestHTTPGetVariation(t *testing.T) {

	respBody, err := variationRequester.HTTPGetVariation(CampaignID, variationGroupID, "testVariationID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testVariationID", respBody.Id)
	assert.Equal(t, "testVariationName", respBody.Name)
}

func TestHTTPListVariation(t *testing.T) {

	respBody, err := variationRequester.HTTPListVariation(CampaignID, variationGroupID)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testVariationID", respBody[0].Id)
	assert.Equal(t, "testVariationName", respBody[0].Name)
	assert.Equal(t, true, respBody[0].Reference)
	assert.Equal(t, 50, respBody[0].Allocation)

	assert.Equal(t, "testVariationID1", respBody[1].Id)
	assert.Equal(t, "testVariationName1", respBody[1].Name)
	assert.Equal(t, false, respBody[1].Reference)
	assert.Equal(t, 80, respBody[1].Allocation)
}

func TestHTTPCreateVariation(t *testing.T) {

	data := "{\"name\":\"testVariationName\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}"

	respBody, err := variationRequester.HTTPCreateVariation(CampaignID, variationGroupID, data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testVariationID\",\"name\":\"testVariationName\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}"), respBody)
}

func TestHTTPEditVariation(t *testing.T) {

	data := "{\"name\":\"testVariationName1\",\"reference\":false,\"allocation\":80,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP1\"}}"

	respBody, err := variationRequester.HTTPEditVariation(CampaignID, variationGroupID, "testVariationID", data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testVariationID\",\"name\":\"testVariationName1\",\"reference\":false,\"allocation\":80,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP1\"}}"), respBody)
}

func TestHTTPDeleteVariation(t *testing.T) {

	err := variationRequester.HTTPDeleteVariation(CampaignID, variationGroupID, "testVariationID")

	assert.Nil(t, err)
}
