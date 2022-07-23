package httprequest

import (
	"net/http"

	"github.com/Chadiii/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPListUsers() ([]byte, error) {
	respBody, err := HTTPRequest(http.MethodGet, utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/users", nil)
	return respBody, err
}

func HTTPBatchUpdateUsers(data string) ([]byte, error) {
	respBody, err := HTTPRequest(http.MethodPut, utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/users", []byte(data))
	return respBody, err
}

func HTTPDeleteUsers(email string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/users?emails[]="+email, nil)
	return err
}
