package web_experimentation

import (
	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/spf13/viper"
)

func HTTPListGlobalCode() ([]models.GlobalCode, error) {
	return common.HTTPGetAllPagesWe[models.GlobalCode](utils.GetWebExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/global-codes")
}

func HTTPGetGlobalCode(id string) (models.GlobalCode, error) {
	return common.HTTPGetItem[models.GlobalCode](utils.GetFeatureExperimentationHost() + "/v1/accounts/" + viper.GetString("account_id") + "/global-codes/" + id)
}
