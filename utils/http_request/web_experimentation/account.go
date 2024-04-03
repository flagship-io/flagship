package web_experimentation

import (
	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type AccountWERequester struct {
	*common.ResourceRequest
}

func (a *AccountWERequester) HTTPListAccount() ([]models.AccountWE, error) {
	return common.HTTPGetAllPagesWE[models.AccountWE](utils.GetWebExperimentationHost() + "/v1/accounts/")
}
