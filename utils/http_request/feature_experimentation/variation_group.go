package feature_experimentation

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/viper"
)

func HTTPListVariationGroup(campaignID string) ([]models.VariationGroup, error) {
	return httprequest.HTTPGetAllPages[models.VariationGroup](utils.GetHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns/" + campaignID + "/variation_groups")
}

func HTTPGetVariationGroup(campaignID, id string) (models.VariationGroup, error) {
	return httprequest.HTTPGetItem[models.VariationGroup](utils.GetHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns/" + campaignID + "/variation_groups/" + id)
}

func HTTPCreateVariationGroup(campaignID, data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPost, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups", []byte(data))
}

func HTTPEditVariationGroup(campaignID, id, data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPatch, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+id, []byte(data))
}

func HTTPDeleteVariationGroup(campaignID, id string) error {
	_, err := httprequest.HTTPRequest(http.MethodDelete, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+id, nil)
	return err
}
