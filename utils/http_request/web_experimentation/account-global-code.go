package web_experimentation

import (
	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type AccountGlobalCodeRequester struct {
	*common.ResourceRequest
}

func (a *AccountGlobalCodeRequester) HTTPGetAccountGlobalCode(id string) (string, error) {
	resp, err := common.HTTPGetItem[models.AccountWE](utils.GetWebExperimentationHost() + "/v1/accounts/" + id)
	return resp.GlobalCode.Value, err
}
