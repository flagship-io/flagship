package httprequest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var CampaignID = "campaignID"

func TestHTTPGetVariationGroup(t *testing.T) {

	respBody, err := HTTPGetVariationGroup(CampaignID, "testVariationGroupID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testVariationGroupID", respBody.Id)
	assert.Equal(t, "testVariationGroupName", respBody.Name)
}

func TestHTTPListVariationGroup(t *testing.T) {

	respBody, err := HTTPListVariationGroup(CampaignID)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testVariationGroupID", respBody[0].Id)
	assert.Equal(t, "testVariationGroupName", respBody[0].Name)

	assert.Equal(t, "testVariationGroupID1", respBody[1].Id)
	assert.Equal(t, "testVariationGroupName1", respBody[1].Name)
}

func TestHTTPCreateVariationGroup(t *testing.T) {

	data := "{\"name\":\"testVariationGroupName\",\"variations\":[{\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"operator\":\"CONTAINS\",\"key\":\"isVIP\",\"value\":true}]}]}}"

	respBody, err := HTTPCreateVariationGroup(CampaignID, data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testVariationGroupID\",\"name\":\"testVariationGroupName\",\"variations\":[{\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}"), respBody)
}

func TestHTTPEditVariationGroup(t *testing.T) {

	data := "{\"name\":\"testVariationGroupName1\",\"variations\":[{\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"operator\":\"CONTAINS\",\"key\":\"isVIP\",\"value\":true}]}]}}"

	respBody, err := HTTPEditVariationGroup(CampaignID, "testVariationGroupID", data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testVariationGroupID\",\"name\":\"testVariationGroupName1\",\"variations\":[{\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}"), respBody)
}

func TestHTTPDeleteVariationGroup(t *testing.T) {

	err := HTTPDeleteVariationGroup(CampaignID, "testVariationGroupID")

	assert.Nil(t, err)
}
