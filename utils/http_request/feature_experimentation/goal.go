package feature_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/viper"
)

func HTTPListGoal() ([]models.Goal, error) {
	return httprequest.HTTPGetAllPages[models.Goal](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/goals")
}

func HTTPGetGoal(id string) (models.Goal, error) {
	return httprequest.HTTPGetItem[models.Goal](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/goals/" + id)
}

func HTTPCreateGoal(data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/goals", []byte(data))
}

func HTTPEditGoal(id, data string) ([]byte, error) {
	return httprequest.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/goals/"+id, []byte(data))
}

func HTTPDeleteGoal(id string) error {
	_, err := httprequest.HTTPRequest(http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/goals/"+id, nil)
	return err
}
