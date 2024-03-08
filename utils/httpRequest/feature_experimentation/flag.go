package feature_experimentation

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPListFlag() ([]models.Flag, error) {
	return HTTPGetAllPages[models.Flag](utils.GetHost() + "/v1/accounts/" + viper.GetString("account_id") + "/flags")
}

func HTTPGetFlag(id string) (models.Flag, error) {
	return HTTPGetItem[models.Flag](utils.GetHost() + "/v1/accounts/" + viper.GetString("account_id") + "/flags/" + id)
}

func HTTPCreateFlag(data string) ([]byte, error) {
	return HTTPRequest(http.MethodPost, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/flags", []byte(data))
}

func HTTPEditFlag(id, data string) ([]byte, error) {
	return HTTPRequest(http.MethodPatch, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+id, []byte(data))
}

func HTTPDeleteFlag(id string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+id, nil)
	return err
}
