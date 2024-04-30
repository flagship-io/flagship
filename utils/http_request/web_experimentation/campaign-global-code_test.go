package web_experimentation

import (
	"testing"

	"github.com/flagship-io/flagship/utils/http_request/common"
	mockfunction_we "github.com/flagship-io/flagship/utils/mock_function/web_experimentation"
	"github.com/stretchr/testify/assert"
)

var campaignGlobalCodeRequester = CampaignGlobalCodeRequester{&common.ResourceRequest{AccountID: "account_id"}}

func TestHTTPCampaignGlobalCode(t *testing.T) {

	respBody, err := campaignGlobalCodeRequester.HTTPGetCampaignGlobalCode("100000")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "console.log(\"Hello World!\")", respBody)

}

func TestHTTPCampaignGlobalCodeInfo(t *testing.T) {

	respBody, err := campaignGlobalCodeRequester.HTTPGetCampaignGlobalCodeInfo("100000")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, mockfunction_we.TestCampaign, respBody)
}
