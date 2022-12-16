package httprequest

import (
	"encoding/json"
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPListCampaign() ([]models.Campaign, error) {
	return HTTPGetAllPages[models.Campaign](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns")
}

func HTTPGetCampaign(id string) (models.Campaign, error) {
	return HTTPGetItem[models.Campaign](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns/" + id)
}

func HTTPCreateCampaign(data string) ([]byte, error) {
	return HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns", []byte(data))
}

func HTTPEditCampaign(id, data string) ([]byte, error) {
	return HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+id, []byte(data))
}

func HTTPToggleCampaign(id, state string) error {
	campaignToggleRequest := models.CampaignToggleRequest{
		State: state,
	}

	campaignToggleRequestJSON, err := json.Marshal(campaignToggleRequest)
	if err != nil {
		return err
	}

	_, err = HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+id+"/toggle", campaignToggleRequestJSON)
	return err
}

func HTTPDeleteCampaign(id string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+id, nil)
	return err
}
