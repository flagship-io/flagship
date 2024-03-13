package feature_experimentation

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/viper"
)

func HTTPListTargetingKey() ([]models.TargetingKey, error) {
	return httprequest.HTTPGetAllPages[models.TargetingKey](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/targeting_keys")
}

func HTTPGetTargetingKey(id string) (models.TargetingKey, error) {
	return httprequest.HTTPGetItem[models.TargetingKey](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/targeting_keys/" + id)
}

func HTTPCreateTargetingKey(data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys", []byte(data))
}

func HTTPEditTargetingKey(id, data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys/"+id, []byte(data))
}

func HTTPDeleteTargetingKey(id string) error {
	_, err := httprequest.HTTPRequest(http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys/"+id, nil)
	return err
}
