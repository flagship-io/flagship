package common

import (
	"encoding/json"
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
)

func HTTPRefreshTokenFE(client_id, refresh_token string) (models.TokenResponse, error) {
	var authenticationResponse models.TokenResponse
	authRequest := models.RefreshTokenRequestFE{
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

func HTTPRefreshTokenWE(client_id, client_secret, refresh_token string) (models.TokenResponse, error) {
	var authenticationResponse models.TokenResponse
	authRequest := models.RefreshTokenRequestWE{
		ClientID:     client_id,
		GrantType:    "refresh_token",
		RefreshToken: refresh_token,
		ClientSecret: client_secret,
	}
	authRequestJSON, err := json.Marshal(authRequest)
	if err != nil {
		return models.TokenResponse{}, err
	}

	respBody, err := HTTPRequest[models.TokenWE](http.MethodPost, utils.GetHostWebExperimentationAuth()+"/v1"+"/token", authRequestJSON)
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

func HTTPCreateTokenWEAuthorizationCode(client_id, client_secret, code string) (models.TokenResponse, error) {
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

func HTTPCreateTokenWEPassword(client_id, client_secret, username, password, mfaCode string) (models.TokenResponse, error) {
	var authenticationResponse models.TokenResponse
	var mfaResponse models.MfaRequestWE
	var mfmResponse models.MfaRequestWE

	authRequest := models.PasswordRequest{
		ClientID:     client_id,
		ClientSecret: client_secret,
		GrantType:    "password",
		Username:     username,
		Password:     password,
	}
	authRequestJSON, err := json.Marshal(authRequest)
	if err != nil {
		return models.TokenResponse{}, err
	}

	mfaRespBody, err := HTTPRequest[models.MfaRequestWE](http.MethodPost, utils.GetHostWebExperimentationAuth()+"/v1/token", authRequestJSON)
	if err != nil {
		return models.TokenResponse{}, err
	}

	err = json.Unmarshal(mfaRespBody, &mfaResponse)
	if err != nil {
		return models.TokenResponse{}, err
	}

	mfmRequest := models.MultiFactorMethodRequestWE{
		GrantType: "multi_factor_methods",
		MfaToken:  mfaResponse.MfaToken,
		MfaMethod: "totp",
	}

	mfmRequestJSON, err := json.Marshal(mfmRequest)
	if err != nil {
		return models.TokenResponse{}, err
	}

	mfmRespBody, err := HTTPRequest[models.MfaRequestWE](http.MethodPost, utils.GetHostWebExperimentationAuth()+"/v1/token", mfmRequestJSON)
	if err != nil {
		return models.TokenResponse{}, err
	}

	err = json.Unmarshal(mfmRespBody, &mfmResponse)
	if err != nil {
		return models.TokenResponse{}, err
	}

	mfRequest := models.MultiFactorRequestWE{
		GrantType: "multi_factor",
		MfaToken:  mfmResponse.MfaToken,
		MfaMethod: "totp",
		Code:      mfaCode,
	}

	mfRequestJSON, err := json.Marshal(mfRequest)
	if err != nil {
		return models.TokenResponse{}, err
	}

	respBody, err := HTTPRequest[models.MfaRequestWE](http.MethodPost, utils.GetHostWebExperimentationAuth()+"/v1/token", mfRequestJSON)
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
	authRequest := models.RefreshTokenRequestFE{
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

func HTTPCheckToken() (models.TokenFE, error) {
	return HTTPGetItem[models.TokenFE](utils.GetHostFeatureExperimentationAuth() + "/token?access_token=" + cred.Token)
}
