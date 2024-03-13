package feature_experimentation

import (
	"encoding/json"
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/viper"
)

func HTTPListCampaign() ([]models.Campaign, error) {
	return httprequest.HTTPGetAllPages[models.Campaign](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns")
}

func HTTPGetCampaign(id string) (models.Campaign, error) {
	return httprequest.HTTPGetItem[models.Campaign](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns/" + id)
}

func HTTPCreateCampaign(data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns", []byte(data))
}

func HTTPEditCampaign(id, data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+id, []byte(data))
}

func HTTPSwitchCampaign(id, state string) error {
	campaignSwitchRequest := models.CampaignSwitchRequest{
		State: state,
	}

	campaignSwitchRequestJSON, err := json.Marshal(campaignSwitchRequest)
	if err != nil {
		return err
	}

	_, err = httprequest.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+id+"/toggle", campaignSwitchRequestJSON)
	return err
}

func HTTPDeleteCampaign(id string) error {
	_, err := httprequest.HTTPRequest(http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+id, nil)
	return err
}
