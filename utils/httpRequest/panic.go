package httprequest

import (
	"encoding/json"
	"net/http"

	"github.com/Chadiii/flagship/models"
	"github.com/Chadiii/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPUpdatePanic(panic string) error {
	panicRequestJSON, err := json.Marshal(models.PanicRequest{Panic: panic})
	if err != nil {
		return err
	}
	_, err = HTTPRequest(http.MethodPatch, utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/panic", panicRequestJSON)
	return err
}
