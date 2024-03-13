package web_experimentation

import (
	"strconv"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/viper"
)

func HTTPGetVariation(testID, id int) (models.TestVariation, error) {
	return httprequest.HTTPGetItem[models.TestVariation](utils.GetWebExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/tests/" + strconv.Itoa(testID) + "/variations/" + strconv.Itoa(id))
}
