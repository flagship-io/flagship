package httprequest

import (
	"encoding/json"
	"net/http"

	"github.com/Chadiii/flagship/models"
	"github.com/Chadiii/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPToggleCampaign(id, state string) error {
	campaignToggleRequest := models.CampaignToggleRequest{
		State: state,
	}

	campaignToggleRequestJSON, err := json.Marshal(campaignToggleRequest)
	if err != nil {
		return err
	}

	_, err = HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns/"+id+"/toggle", campaignToggleRequestJSON)
	return err
}

func HTTPListCampaign() ([]byte, error) {
	respBody, err := HTTPRequest(http.MethodGet, utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns", nil)
	return respBody, err
}

func HTTPCreateCampaign(data string) ([]byte, error) {
	respBody, err := HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns", []byte(data))
	return respBody, err
}

func HTTPGetCampaign(id string) ([]byte, error) {
	respBody, err := HTTPRequest(http.MethodGet, utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns/"+id, nil)
	return respBody, err
}

func HTTPDeleteCampaign(id string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns/"+id, nil)
	return err
}

func HTTPEditCampaign(id, data string) ([]byte, error) {
	respBody, err := HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns/"+id, []byte(data))
	return respBody, err
}
