package web_experimentation

import (
	"encoding/json"
	"net/http"

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

func (a *AccountGlobalCodeRequester) HTTPPushAccountGlobalCode(id string, code []byte) ([]byte, error) {
	var gc = models.GlobalCodeStr{
		GlobalCode: string(code),
	}

	data, err := json.Marshal(gc)
	if err != nil {
		return nil, err
	}

	return common.HTTPRequest[models.AccountWE](http.MethodPatch, utils.GetWebExperimentationHost()+"/v1/accounts/"+a.AccountID, []byte(data))
}
