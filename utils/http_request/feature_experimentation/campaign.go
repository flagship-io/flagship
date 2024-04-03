package feature_experimentation

import (
	"encoding/json"
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type CampaignFERequester struct {
	*common.ResourceRequest
}

func (c *CampaignFERequester) HTTPListCampaign() ([]models.CampaignFE, error) {
	return common.HTTPGetAllPagesFE[models.CampaignFE](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + c.AccountID + "/account_environments/" + c.AccountEnvironmentID + "/campaigns")
}

func (c *CampaignFERequester) HTTPGetCampaign(id string) (models.CampaignFE, error) {
	return common.HTTPGetItem[models.CampaignFE](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + c.AccountID + "/account_environments/" + c.AccountEnvironmentID + "/campaigns/" + id)
}

func (c *CampaignFERequester) HTTPCreateCampaign(data string) ([]byte, error) {
	return common.HTTPRequest[models.CampaignFE](http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+c.AccountID+"/account_environments/"+c.AccountEnvironmentID+"/campaigns", []byte(data))
}

func (c *CampaignFERequester) HTTPEditCampaign(id, data string) ([]byte, error) {
	return common.HTTPRequest[models.CampaignFE](http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+c.AccountID+"/account_environments/"+c.AccountEnvironmentID+"/campaigns/"+id, []byte(data))
}

func (c *CampaignFERequester) HTTPSwitchCampaign(id, state string) error {
	campaignSwitchRequest := models.CampaignFESwitchRequest{
		State: state,
	}

	campaignSwitchRequestJSON, err := json.Marshal(campaignSwitchRequest)
	if err != nil {
		return err
	}

	_, err = common.HTTPRequest[models.CampaignFE](http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+c.AccountID+"/account_environments/"+c.AccountEnvironmentID+"/campaigns/"+id+"/toggle", campaignSwitchRequestJSON)
	return err
}

func (c *CampaignFERequester) HTTPDeleteCampaign(id string) error {
	_, err := common.HTTPRequest[models.CampaignFE](http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+c.AccountID+"/account_environments/"+c.AccountEnvironmentID+"/campaigns/"+id, nil)
	return err
}
