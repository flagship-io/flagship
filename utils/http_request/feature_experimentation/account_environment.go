package feature_experimentation

import (
	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type AccountEnvironmentFERequester struct {
	*common.ResourceRequest
}

func (a *AccountEnvironmentFERequester) HTTPListAccountEnvironment(accountID string) ([]models.AccountEnvironmentFE, error) {
	if accountID == "" {
		accountID = a.AccountID
	}

	return common.HTTPGetAllPagesFE[models.AccountEnvironmentFE](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + accountID + "/account_environments")
}

func (a *AccountEnvironmentFERequester) HTTPGetAccountEnvironment(id string) (models.AccountEnvironmentFE, error) {
	return common.HTTPGetItem[models.AccountEnvironmentFE](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + a.AccountID + "/account_environments/" + id)
}
