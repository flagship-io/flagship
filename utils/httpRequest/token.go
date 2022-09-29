package httprequest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPCreateToken(client_id, client_secret, grant_type, scope string, expiration int) (string, error) {
	var authenticationResponse models.AuthenticationResponse
	authRequest := models.AuthenticationRequest{
		ClientID:     client_id,
		ClientSecret: client_secret,
		Scope:        scope,
		GrantType:    grant_type,
	}
	authRequestJSON, err := json.Marshal(authRequest)
	if err != nil {
		return "", err
	}

	respBody, err := HTTPRequest(http.MethodPost, utils.HostAuth+"/"+viper.GetString("account_id")+"/token?expires_in="+strconv.Itoa(expiration), authRequestJSON)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(respBody, &authenticationResponse)
	if err != nil {
		return "", err
	}

	return authenticationResponse.AccessToken, err
}

func HTTPCheckToken(token string) (models.Token, error) {
	return HTTPGetItem[models.Token](utils.HostAuth + "/token?access_token=" + token)
}
