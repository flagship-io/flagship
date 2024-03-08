package feature_experimentation

import (
	"net/http"
	"net/url"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPListUsers() ([]models.User, error) {
	return HTTPGetAllPages[models.User](utils.GetHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/users")
}

func HTTPBatchUpdateUsers(data string) ([]byte, error) {
	return HTTPRequest(http.MethodPut, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users", []byte(data))
}

func HTTPDeleteUsers(email string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users?emails[]="+url.QueryEscape(email), nil)
	return err
}
