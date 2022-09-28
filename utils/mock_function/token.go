package mockfunction

import (
	"net/http"
	"strconv"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
)

func APIToken() {
	config.SetViper()

	token := "token"
	tokenExpiration := 0

	testToken := models.Token{
		ClientID:  "client_id",
		AccountID: "account_id",
		ExpiresIn: 0,
		Scope:     "*",
	}

	testAuthenticationResponse := models.AuthenticationResponse{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}

	httpmock.RegisterResponder("GET", utils.HostAuth+"/token?access_token="+token,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testToken)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.HostAuth+"/"+viper.GetString("account_id")+"/token?expires_in="+strconv.Itoa(tokenExpiration),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testAuthenticationResponse)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
}
