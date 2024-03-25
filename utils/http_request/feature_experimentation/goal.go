package feature_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type GoalRequester struct {
	*common.ResourceRequest
}

func (g *GoalRequester) HTTPListGoal() ([]models.Goal, error) {
	return common.HTTPGetAllPages[models.Goal](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + g.AccountID + "/account_environments/" + g.AccountEnvID + "/goals")
}

func (g *GoalRequester) HTTPGetGoal(id string) (models.Goal, error) {
	return common.HTTPGetItem[models.Goal](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + g.AccountID + "/account_environments/" + g.AccountEnvID + "/goals/" + id)
}

func (g *GoalRequester) HTTPCreateGoal(data string) ([]byte, error) {
	return common.HTTPRequest(http.MethodPost, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+g.AccountID+"/account_environments/"+g.AccountEnvID+"/goals", []byte(data))
}

func (g *GoalRequester) HTTPEditGoal(id, data string) ([]byte, error) {
	return common.HTTPRequest(http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+g.AccountID+"/account_environments/"+g.AccountEnvID+"/goals/"+id, []byte(data))
}

func (g *GoalRequester) HTTPDeleteGoal(id string) error {
	_, err := common.HTTPRequest(http.MethodDelete, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+g.AccountID+"/account_environments/"+g.AccountEnvID+"/goals/"+id, nil)
	return err
}
