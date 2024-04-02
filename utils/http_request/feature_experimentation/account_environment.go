package feature_experimentation

import (
	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type AccountEnvironmentRequester struct {
	*common.ResourceRequest
}

func (a *AccountEnvironmentRequester) HTTPListAccountEnvironment() ([]models.AccountEnvironment, error) {
	return common.HTTPGetAllPages[models.AccountEnvironment](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + a.AccountID + "/account_environments")
}

func (a *AccountEnvironmentRequester) HTTPGetAccountEnvironment(id string) (models.AccountEnvironment, error) {
	return common.HTTPGetItem[models.AccountEnvironment](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + a.AccountID + "/account_environments/" + id)
}
