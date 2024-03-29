package httprequest

import (
	"encoding/json"
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPUpdatePanic(panicStatus string) ([]byte, error) {
	panicRequestJSON, err := json.Marshal(models.PanicRequest{Panic: panicStatus})
	if err != nil {
		return nil, err
	}
	resp, err := HTTPRequest(http.MethodPatch, utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/panic", panicRequestJSON)
	return resp, err
}
