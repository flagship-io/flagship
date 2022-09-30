package httprequest

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPListFlag() ([]models.Flag, error) {
	return HTTPGetAllPages[models.Flag](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/flags")
}

func HTTPFlagUsage() ([]models.FlagUsage, error) {
	return HTTPGetAllPages[models.FlagUsage](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/flags_usage")
}

func HTTPGetFlag(id string) (models.Flag, error) {
	return HTTPGetItem[models.Flag](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/flags/" + id)
}

func HTTPCreateFlag(data string) ([]byte, error) {
	return HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags", []byte(data))
}

func HTTPEditFlag(id, data string) ([]byte, error) {
	return HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+id, []byte(data))
}

func HTTPDeleteFlag(id string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+id, nil)
	return err
}
