package feature_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type VariationRequester struct {
	*common.ResourceRequest
}

func (v *VariationRequester) HTTPListVariation(campaignID, variationGroupID string) ([]models.Variation, error) {
	return common.HTTPGetAllPages[models.Variation](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + v.AccountID + "/account_environments/" + v.AccountEnvID + "/campaigns/" + campaignID + "/variation_groups/" + variationGroupID + "/variations")
}

func (v *VariationRequester) HTTPGetVariation(campaignID, variationGroupID, id string) (models.Variation, error) {
	return common.HTTPGetItem[models.Variation](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + v.AccountID + "/account_environments/" + v.AccountEnvID + "/campaigns/" + campaignID + "/variation_groups/" + variationGroupID + "/variations/" + id)
}

func (v *VariationRequester) HTTPCreateVariation(campaignID, variationGroupID, data string) ([]byte, error) {
	return common.HTTPRequest(http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+v.AccountID+"/account_environments/"+v.AccountEnvID+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations", []byte(data))
}

func (v *VariationRequester) HTTPEditVariation(campaignID, variationGroupID, id, data string) ([]byte, error) {
	return common.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+v.AccountID+"/account_environments/"+v.AccountEnvID+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+id, []byte(data))
}

func (v *VariationRequester) HTTPDeleteVariation(campaignID, variationGroupID, id string) error {
	_, err := common.HTTPRequest(http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+v.AccountID+"/account_environments/"+v.AccountEnvID+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+id, nil)
	return err
}
