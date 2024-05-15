package web_experimentation

import (
	"net/http"
	"strconv"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
)

type VariationWERequester struct {
	*common.ResourceRequest
}

func (v *VariationWERequester) HTTPGetVariation(testID, id int) (models.VariationWE, error) {
	return common.HTTPGetItem[models.VariationWE](utils.GetWebExperimentationHost() + "/v1/accounts/" + v.AccountID + "/tests/" + strconv.Itoa(testID) + "/variations/" + strconv.Itoa(id))
}

func (v *VariationWERequester) HTTPDeleteVariation(testID, id int) error {
	_, err := common.HTTPRequest[models.VariationWE](http.MethodDelete, utils.GetWebExperimentationHost()+"/v1/accounts/"+v.AccountID+"/tests/"+strconv.Itoa(testID)+"/variations/"+strconv.Itoa(id), nil)
	return err
}
