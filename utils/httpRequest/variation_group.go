package httprequest

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPListVariationGroup(campaignID string) ([]models.VariationGroup, error) {
	return HTTPGetAllPages[models.VariationGroup](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns/" + campaignID + "/variation_groups")
}

func HTTPGetVariationGroup(campaignID, id string) (models.VariationGroup, error) {
	return HTTPGetItem[models.VariationGroup](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns/" + campaignID + "/variation_groups/" + id)
}

func HTTPCreateVariationGroup(campaignID, data string) ([]byte, error) {
	return HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups", []byte(data))
}

func HTTPEditVariationGroup(campaignID, id, data string) ([]byte, error) {
	return HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+id, []byte(data))
}

func HTTPDeleteVariationGroup(campaignID, id string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+id, nil)
	return err
}
