package feature_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/viper"
)

func HTTPListFlag() ([]models.Flag, error) {
	return httprequest.HTTPGetAllPages[models.Flag](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/flags")
}

func HTTPGetFlag(id string) (models.Flag, error) {
	return httprequest.HTTPGetItem[models.Flag](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/flags/" + id)
}

func HTTPCreateFlag(data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/flags", []byte(data))
}

func HTTPEditFlag(id, data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+id, []byte(data))
}

func HTTPDeleteFlag(id string) error {
	_, err := httprequest.HTTPRequest(http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+id, nil)
	return err
}
