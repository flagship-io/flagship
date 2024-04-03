package web_experimentation

import (
	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type GlobalCodeRequester struct {
	*common.ResourceRequest
}

func (g *GlobalCodeRequester) HTTPListGlobalCode() ([]models.GlobalCode, error) {
	return common.HTTPGetAllPagesWE[models.GlobalCode](utils.GetWebExperimentationHost() + "/v1/accounts/" + g.AccountID + "/global-codes")
}

func (g *GlobalCodeRequester) HTTPGetGlobalCode(id string) (models.GlobalCode, error) {
	return common.HTTPGetItem[models.GlobalCode](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + g.AccountID + "/global-codes/" + id)
}
