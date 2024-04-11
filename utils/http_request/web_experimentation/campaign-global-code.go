package web_experimentation

import (
	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type CampaignGlobalCodeRequester struct {
	*common.ResourceRequest
}

func (c *CampaignGlobalCodeRequester) HTTPGetCampaignGlobalCode(id string) (string, error) {
	resp, err := common.HTTPGetItem[models.CampaignWE](utils.GetWebExperimentationHost() + "/v1/accounts/" + c.AccountID + "/tests/" + id)
	return resp.GlobalCodeCampaign, err
}
