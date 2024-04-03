package web_experimentation

import (
	"strconv"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/spf13/viper"
)

type VariationWERequester struct {
	*common.ResourceRequest
}

func (v *VariationWERequester) HTTPGetVariation(testID, id int) (models.VariationWE, error) {
	return common.HTTPGetItem[models.VariationWE](utils.GetWebExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/tests/" + strconv.Itoa(testID) + "/variations/" + strconv.Itoa(id))
}
