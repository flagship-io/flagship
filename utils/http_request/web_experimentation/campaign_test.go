package web_experimentation

import (
	"testing"

	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/stretchr/testify/assert"
)

var campaignRequester = CampaignWERequester{&common.ResourceRequest{AccountID: "account_id"}}

func TestHTTPGetCampaign(t *testing.T) {

	respBody, err := campaignRequester.HTTPGetCampaign("100000")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, 100000, respBody.Id)
	assert.Equal(t, "testCampaignName", respBody.Name)
	assert.Equal(t, "console.log(\"Hello World!\")", respBody.GlobalCodeCampaign)
	assert.Equal(t, "testCampaignDescription", respBody.Description)
	assert.Equal(t, "ab", respBody.Type)

}

func TestHTTPListCampaign(t *testing.T) {

	respBody, err := campaignRequester.HTTPListCampaign()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, 100000, respBody[0].Id)
	assert.Equal(t, "testCampaignName", respBody[0].Name)
	assert.Equal(t, "console.log(\"Hello World!\")", respBody[0].GlobalCodeCampaign)
	assert.Equal(t, "testCampaignDescription", respBody[0].Description)
	assert.Equal(t, "ab", respBody[0].Type)

	assert.Equal(t, 100001, respBody[1].Id)
	assert.Equal(t, "testCampaignName1", respBody[1].Name)
	assert.Equal(t, "console.log(\"Hello Earth!\")", respBody[1].GlobalCodeCampaign)
	assert.Equal(t, "testCampaignDescription1", respBody[1].Description)
	assert.Equal(t, "ab", respBody[1].Type)

}

func TestHTTPDeleteCampaign(t *testing.T) {

	err := campaignRequester.HTTPDeleteCampaign("100000")

	assert.Nil(t, err)
}
