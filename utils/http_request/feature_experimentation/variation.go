package feature_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/viper"
)

func HTTPListVariation(campaignID, variationGroupID string) ([]models.Variation, error) {
	return httprequest.HTTPGetAllPages[models.Variation](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns/" + campaignID + "/variation_groups/" + variationGroupID + "/variations")
}

func HTTPGetVariation(campaignID, variationGroupID, id string) (models.Variation, error) {
	return httprequest.HTTPGetItem[models.Variation](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns/" + campaignID + "/variation_groups/" + variationGroupID + "/variations/" + id)
}

func HTTPCreateVariation(campaignID, variationGroupID, data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations", []byte(data))
}

func HTTPEditVariation(campaignID, variationGroupID, id, data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+id, []byte(data))
}

func HTTPDeleteVariation(campaignID, variationGroupID, id string) error {
	_, err := httprequest.HTTPRequest(http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+id, nil)
	return err
}
