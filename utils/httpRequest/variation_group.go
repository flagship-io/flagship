package httprequest

import (
	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPListVariationGroup(campaingID string) ([]models.VariationGroup, error) {
	return HTTPGetAllPages[models.VariationGroup](utils.Host + "/v1/accounts/" + viper.GetString("account_id") + "/account_environments/" + viper.GetString("account_environment_id") + "/campaigns/" + campaingID + "/variation_groups")
}
