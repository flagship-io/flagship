package httprequest

import (
	"encoding/json"
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPUpdatePanic(panic string) error {
	panicRequestJSON, err := json.Marshal(models.PanicRequest{Panic: panic})
	if err != nil {
		return err
	}
	_, err = HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/panic", panicRequestJSON)
	return err
}
