package common

import (
	"encoding/json"
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
)

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

	respBody, err := HTTPRequest[models.TokenWE](http.MethodPost, utils.GetHostFeatureExperimentationAuth()+"/"+cred.AccountID+"/token", authRequestJSON)
	if err != nil {
		return models.TokenResponse{}, err
	}

	err = json.Unmarshal(respBody, &authenticationResponse)
	if err != nil {
		return models.TokenResponse{}, err
	}

	return authenticationResponse, err
}

func HTTPCreateTokenFE(clientId, clientSecret, accountId string) (models.TokenResponse, error) {
	var authenticationResponse models.TokenResponse
	authRequest := models.ClientCredentialsRequest{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scope:        "*",
		GrantType:    "client_credentials",
	}
	authRequestJSON, err := json.Marshal(authRequest)
	if err != nil {
		return models.TokenResponse{}, err
	}

	respBody, err := HTTPRequest[models.TokenFE](http.MethodPost, utils.GetHostFeatureExperimentationAuth()+"/"+accountId+"/token?expires_in=86400", authRequestJSON)
	if err != nil {
		return models.TokenResponse{}, err
	}

	err = json.Unmarshal(respBody, &authenticationResponse)
	if err != nil {
		return models.TokenResponse{}, err
	}

	return authenticationResponse, err
}

func HTTPCreateTokenWE(client_id, client_secret, code string) (models.TokenResponse, error) {
	var authenticationResponse models.TokenResponse
	authRequest := models.AuthorizationCodeRequest{
		ClientID:     client_id,
		ClientSecret: client_secret,
		GrantType:    "authorization_code",
		Code:         code,
	}
	authRequestJSON, err := json.Marshal(authRequest)
	if err != nil {
		return models.TokenResponse{}, err
	}

	respBody, err := HTTPRequest[models.TokenWE](http.MethodPost, utils.GetHostWebExperimentationAuth()+"/v1/token", authRequestJSON)
	if err != nil {
		return models.TokenResponse{}, err
	}

	err = json.Unmarshal(respBody, &authenticationResponse)
	if err != nil {
		return models.TokenResponse{}, err
	}

	return authenticationResponse, err
}

func HTTPRefreshToken_(product, client_id, refresh_token string) (models.TokenResponse, error) {
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

	respBody, err := HTTPRequest[models.TokenWE](http.MethodPost, utils.GetHostFeatureExperimentationAuth()+"/"+cred.AccountID+"/token", authRequestJSON)
	if err != nil {
		return models.TokenResponse{}, err
	}

	err = json.Unmarshal(respBody, &authenticationResponse)
	if err != nil {
		return models.TokenResponse{}, err
	}

	return authenticationResponse, err
}

func HTTPCheckToken(token string) (models.TokenFE, error) {
	return HTTPGetItem[models.TokenFE](utils.GetHostFeatureExperimentationAuth() + "/token?access_token=" + token)
}
