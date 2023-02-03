package httprequest

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPListVariation(campaignID, variationGroupID string) ([]models.Variation, error) {
	return HTTPGetAllPages[models.Variation](utils.GetHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns/" + campaignID + "/variation_groups/" + variationGroupID + "/variations")
}

func HTTPGetVariation(campaignID, variationGroupID, id string) (models.Variation, error) {
	return HTTPGetItem[models.Variation](utils.GetHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns/" + campaignID + "/variation_groups/" + variationGroupID + "/variations/" + id)
}

func HTTPCreateVariation(campaignID, variationGroupID, data string) ([]byte, error) {
	return HTTPRequest(http.MethodPost, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations", []byte(data))
}

func HTTPEditVariation(campaignID, variationGroupID, id, data string) ([]byte, error) {
	return HTTPRequest(http.MethodPatch, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+id, []byte(data))
}

func HTTPDeleteVariation(campaignID, variationGroupID, id string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+id, nil)
	return err
}
