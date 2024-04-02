package feature_experimentation

import (
	"encoding/json"
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type PanicRequester struct {
	*common.ResourceRequest
}

func (p *PanicRequester) HTTPUpdatePanic(panicStatus string) ([]byte, error) {
	panicRequestJSON, err := json.Marshal(models.PanicRequest{Panic: panicStatus})
	if err != nil {
		return nil, err
	}
	resp, err := common.HTTPRequest[models.PanicRequest](http.MethodPatch, utils.GetFeatureExperimentationHost()+"/v1/accounts/"+p.AccountID+"/account_environments/"+p.AccountEnvironmentID+"/panic", panicRequestJSON)
	return resp, err
}
