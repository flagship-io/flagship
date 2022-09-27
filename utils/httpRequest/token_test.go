package httprequest

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestHTTPCheckToken(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var token string = "token"

	testToken := models.Token{
		ClientID:  "client_id",
		AccountID: "account_id",
		ExpiresIn: 0,
		Scope:     "*",
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

	respBody, err := HTTPCheckToken(token)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "client_id", respBody.ClientID)
	assert.Equal(t, "account_id", respBody.AccountID)
	assert.Equal(t, 0, respBody.ExpiresIn)
	assert.Equal(t, "*", respBody.Scope)
}

func TestHTTPCreateToken(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var tokenExpiration int = 0

	testAuthenticationRequest := models.AuthenticationRequest{
		ClientID:     "client_id",
		ClientSecret: "client_secret",
		GrantType:    "client_credentials",
		Scope:        "*",
	}

	testAuthenticationResponse := models.AuthenticationResponse{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}

	httpmock.RegisterResponder("POST", utils.HostAuth+"/"+viper.GetString("account_id")+"/token?expires_in="+strconv.Itoa(tokenExpiration),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testAuthenticationResponse)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPCreateToken(testAuthenticationRequest.ClientID, testAuthenticationRequest.ClientSecret, testAuthenticationRequest.GrantType, testAuthenticationRequest.Scope, tokenExpiration)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "access_token", respBody)
}
