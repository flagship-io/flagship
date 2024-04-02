package web_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/spf13/viper"
)

func HTTPListTest() ([]models.Test, error) {
	return common.HTTPGetAllPagesWe[models.Test](utils.GetWebExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/tests")
}

func HTTPGetTest(id string) (models.Test, error) {
	return common.HTTPGetItem[models.Test](utils.GetWebExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/tests/" + id)
}

func HTTPCreateTest(data string) ([]byte, error) {
	return common.HTTPRequest[models.Test](http.MethodPost, utils.GetWebExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/tests", []byte(data))
}

func HTTPEditTest(id, data string) ([]byte, error) {
	return common.HTTPRequest[models.Test](http.MethodPatch, utils.GetWebExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/tests/"+id, []byte(data))
}

/* func HTTPSwitchCampaign(id, state string) error {
	campaignSwitchRequest := models.CampaignSwitchRequest{
		State: state,
	}

	campaignSwitchRequestJSON, err := json.Marshal(campaignSwitchRequest)
	if err != nil {
		return err
	}

	_, err = common.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+id+"/toggle", campaignSwitchRequestJSON)
	return err
} */

func HTTPDeleteTest(id string) error {
	_, err := common.HTTPRequest[models.Test](http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+viper.GetString("account_id")+"/tests/"+id, nil)
	return err
}
