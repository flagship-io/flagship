package httprequest

import (
	"net/http"

	"github.com/flagship-io/flagship-cli/models"
	"github.com/flagship-io/flagship-cli/utils"
	"github.com/spf13/viper"
)

func HTTPListTargetingKey() ([]models.TargetingKey, error) {
	return HTTPGetAllPages[models.TargetingKey](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/targeting_keys")
}

func HTTPGetTargetingKey(id string) (models.TargetingKey, error) {
	return HTTPGetItem[models.TargetingKey](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/targeting_keys/" + id)
}

func HTTPCreateTargetingKey(data string) ([]byte, error) {
	return HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys", []byte(data))
}

func HTTPEditTargetingKey(id, data string) ([]byte, error) {
	return HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys/"+id, []byte(data))
}

func HTTPDeleteTargetingKey(id string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys/"+id, nil)
	return err
}
