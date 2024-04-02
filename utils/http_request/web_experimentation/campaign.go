package web_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type TestRequester struct {
	*common.ResourceRequest
}

func (t *TestRequester) HTTPListTest() ([]models.Test, error) {
	return common.HTTPGetAllPagesWE[models.Test](utils.GetWebExperimentationHost() + "/v1/accounts/" + t.AccountID + "/tests")
}

func (t *TestRequester) HTTPGetTest(id string) (models.Test, error) {
	return common.HTTPGetItem[models.Test](utils.GetWebExperimentationHost() + "/v1/accounts/" + t.AccountID + "/tests/" + id)
}

func (t *TestRequester) HTTPCreateTest(data string) ([]byte, error) {
	return common.HTTPRequest[models.Test](http.MethodPost, utils.GetWebExperimentationHost()+"/v1/accounts/"+t.AccountID+"/tests", []byte(data))
}

func (t *TestRequester) HTTPEditTest(id, data string) ([]byte, error) {
	return common.HTTPRequest[models.Test](http.MethodPatch, utils.GetWebExperimentationHost()+"/v1/accounts/"+t.AccountID+"/tests/"+id, []byte(data))
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

func (t *TestRequester) HTTPDeleteTest(id string) error {
	_, err := common.HTTPRequest[models.Test](http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+t.AccountID+"/tests/"+id, nil)
	return err
}
