package web_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type CampaignWERequester struct {
	*common.ResourceRequest
}

func (t *CampaignWERequester) HTTPListCampaign() ([]models.CampaignWE, error) {
	return common.HTTPGetAllPagesWE[models.CampaignWE](utils.GetWebExperimentationHost() + "/v1/accounts/" + t.AccountID + "/tests")
}

func (t *CampaignWERequester) HTTPGetCampaign(id string) (models.CampaignWE, error) {
	return common.HTTPGetItem[models.CampaignWE](utils.GetWebExperimentationHost() + "/v1/accounts/" + t.AccountID + "/tests/" + id)
}

func (t *CampaignWERequester) HTTPDeleteCampaign(id string) error {
	_, err := common.HTTPRequest[models.CampaignWE](http.MethodDelete, utils.GetWebExperimentationHost()+"/v1/accounts/"+t.AccountID+"/tests/"+id, nil)

	return err
}
