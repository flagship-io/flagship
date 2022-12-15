package mockfunction

import (
	"net/http"

	"github.com/flagship-io/flagship-cli/utils"
	"github.com/flagship-io/flagship-cli/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
)

func APIPanic() {
	config.SetViper()

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/panic",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, "")
			return resp, nil
		},
	)
}
