package web_experimentation

import (
	"encoding/json"
	"net/http"

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

func (c *CampaignGlobalCodeRequester) HTTPPushCampaignGlobalCode(id string, code []byte) ([]byte, error) {
	var gc = models.GlobalCodeStr{
		GlobalCode: string(code),
	}

	data, err := json.Marshal(gc)
	if err != nil {
		return nil, err
	}

	return common.HTTPRequest[models.CampaignWE](http.MethodPatch, utils.GetWebExperimentationHost()+"/v1/accounts/"+c.AccountID+"/tests/"+id, []byte(data))
}

func (c *CampaignGlobalCodeRequester) HTTPGetCampaignGlobalCodeInfo(id string) (models.CampaignWE, error) {
	resp, err := common.HTTPGetItem[models.CampaignWE](utils.GetWebExperimentationHost() + "/v1/accounts/" + c.AccountID + "/tests/" + id)
	return resp, err
}
