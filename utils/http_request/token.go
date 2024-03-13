package http_request

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPCreateToken(client_id, client_secret, grant_type, scope string, expiration int) (models.TokenResponse, error) {
	var authenticationResponse models.TokenResponse
	authRequest := models.ClientCredentialsRequest{
		ClientID:     client_id,
		ClientSecret: client_secret,
		Scope:        scope,
		GrantType:    "client_credentials",
	}
	authRequestJSON, err := json.Marshal(authRequest)
	if err != nil {
		return models.TokenResponse{}, err
	}

	respBody, err := HTTPRequest(http.MethodPost, utils.GetHostFeatureExperimentationAuth()+"/"+viper.GetString("account_id")+"/token?expires_in="+strconv.Itoa(expiration), authRequestJSON)
	if err != nil {
		return models.TokenResponse{}, err
	}

	err = json.Unmarshal(respBody, &authenticationResponse)
	if err != nil {
		return models.TokenResponse{}, err
	}

	return authenticationResponse, err
}

func HTTPRefreshToken(client_id, refresh_token string) (models.TokenResponse, error) {
	var authenticationResponse models.TokenResponse
	authRequest := models.RefreshTokenRequest{
		ClientID:     client_id,
		GrantType:    "refresh_token",
		RefreshToken: refresh_token,
	}
	authRequestJSON, err := json.Marshal(authRequest)
	if err != nil {
		return models.TokenResponse{}, err
	}

	respBody, err := HTTPRequest(http.MethodPost, utils.GetHostFeatureExperimentationAuth()+"/"+viper.GetString("account_id")+"/token", authRequestJSON)
	if err != nil {
		return models.TokenResponse{}, err
	}

	err = json.Unmarshal(respBody, &authenticationResponse)
	if err != nil {
		return models.TokenResponse{}, err
	}

	return authenticationResponse, err
}

func HTTPCheckToken(token string) (models.Token, error) {
	return HTTPGetItem[models.Token](utils.GetHostFeatureExperimentationAuth() + "/token?access_token=" + token)
}
