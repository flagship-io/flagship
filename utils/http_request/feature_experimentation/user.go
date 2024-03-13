package feature_experimentation

import (
	"net/http"
	"net/url"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/viper"
)

func HTTPListUsers() ([]models.User, error) {
	return httprequest.HTTPGetAllPages[models.User](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/users")
}

func HTTPBatchUpdateUsers(data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPut, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users", []byte(data))
}

func HTTPDeleteUsers(email string) error {
	_, err := httprequest.HTTPRequest(http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users?emails[]="+url.QueryEscape(email), nil)
	return err
}