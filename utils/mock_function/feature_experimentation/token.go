package feature_experimentation

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
)

var TestToken = models.Token{
	ClientID:  "client_id",
	AccountID: "account_id",
	ExpiresIn: 0,
	Scope:     "scope",
}

func APIToken() {

	token := "access_token"

	testAuthenticationResponse := models.TokenResponse{
		AccessToken:  "testAccessToken",
		RefreshToken: "testRefreshToken",
	}

	httpmock.RegisterResponder("GET", utils.GetHostFeatureExperimentationAuth()+"/token?access_token="+token,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestToken)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.GetHostFeatureExperimentationAuth()+"/"+mockfunction.Auth.AccountID+"/token?expires_in=43200",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, testAuthenticationResponse)
			return resp, nil
		},
	)
}
