package httprequest

import (
	"encoding/json"
	"net/http"

	"github.com/Chadiii/flagship/models"
	"github.com/Chadiii/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPCreateToken(client_id, client_secret, grant_type, scope, expiration string) (string, error) {
	var authenticationResponse models.AuthenticationResponse
	authRequest := models.AuthenticationRequest{
		Client_id:     client_id,
		Client_secret: client_secret,
		Scope:         scope,
		Grant_type:    grant_type,
	}
	authRequestJSON, err := json.Marshal(authRequest)
	if err != nil {
		return "", err
	}

	respBody, err := HTTPRequest(http.MethodPost, utils.HostAuth+"/"+viper.GetString("account_id")+"/token?expires_in="+expiration, authRequestJSON)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(respBody, &authenticationResponse)
	if err != nil {
		return "", err
	}

	return authenticationResponse.Access_token, err
}

func HTTPCheckToken(token string) error {
	_, err := HTTPRequest(http.MethodGet, utils.HostAuth+"/token?access_token="+token, nil)
	return err
}
