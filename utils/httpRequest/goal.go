package httprequest

import (
	"net/http"

	"github.com/flagship-io/flagship-cli/models"
	"github.com/flagship-io/flagship-cli/utils"
	"github.com/spf13/viper"
)

func HTTPListGoal() ([]models.Goal, error) {
	return HTTPGetAllPages[models.Goal](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/goals")
}

func HTTPGetGoal(id string) (models.Goal, error) {
	return HTTPGetItem[models.Goal](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/goals/" + id)
}

func HTTPCreateGoal(data string) ([]byte, error) {
	return HTTPRequest(http.MethodPost, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/goals", []byte(data))
}

func HTTPEditGoal(id, data string) ([]byte, error) {
	return HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/goals/"+id, []byte(data))
}

func HTTPDeleteGoal(id string) error {
	_, err := HTTPRequest(http.MethodDelete, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/goals/"+id, nil)
	return err
}
