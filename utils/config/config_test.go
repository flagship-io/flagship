package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/stretchr/testify/assert"
)

var (
	product      = "test_product"
	username     = "test_user"
	clientID     = "client_id"
	clientSecret = "client_secret"
	accessToken  = "access_token"
	refreshToken = "refresh_token"
	scope        = "scope"
	accountID    = "account_id"
	accountEnvID = "account_environment_id"
)
var authResponse = models.TokenResponse{
	AccessToken:  accessToken,
	RefreshToken: refreshToken,
	Scope:        scope,
}

func TestMain(m *testing.M) {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v", err)
	}

	if _, err := os.Stat(homeDir + "/.flagship/credentials/" + product); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(homeDir+"/.flagship/credentials/"+product, os.ModePerm)
	}

	defer os.RemoveAll(currentDir + "/abtasty")
	defer os.RemoveAll(homeDir + "/.flagship/credentials/" + product)

	m.Run()
}

func TestCheckABTastyHomeDirectory(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	abtastyHome, err := CheckABTastyHomeDirectory()
	if err != nil {
		t.Errorf("CheckABTastyHomeDirectory() error = %v", err)
	}

	assert.Equal(t, homeDir, abtastyHome)
	assert.Equal(t, homeDir+"/.flagship/credentials/"+utils.FEATURE_EXPERIMENTATION, abtastyHome+"/.flagship/credentials/"+utils.FEATURE_EXPERIMENTATION)
	assert.Equal(t, homeDir+"/.flagship/credentials/"+utils.WEB_EXPERIMENTATION, abtastyHome+"/.flagship/credentials/"+utils.WEB_EXPERIMENTATION)
	assert.Equal(t, homeDir+"/.flagship/credentials/"+product, abtastyHome+"/.flagship/credentials/"+product)

}

func TestCredentialPath(t *testing.T) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	filepath, err := CredentialPath(product, username)
	if err != nil {
		t.Errorf("CredentialPath() error = %v", err)
	}

	expectedPath := homeDir + "/.flagship/credentials/" + product + "/" + username + ".yaml"
	assert.Equal(t, expectedPath, filepath)

}

func TestGetUsernames(t *testing.T) {

	err := CreateAuthFile(product, username, clientID, clientSecret, authResponse)
	if err != nil {
		t.Errorf("GetUsernames() error = %v", err)
	}

	fileNames, err := GetUsernames(product)
	if err != nil {
		t.Errorf("GetUsernames() error = %v", err)
	}

	if len(fileNames) != 1 || fileNames[0] != "test_user" {
		t.Errorf("GetUsernames() returned unexpected file names: %v", fileNames)
	}
}

func TestCreateAuthFile(t *testing.T) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	err = CreateAuthFile(product, username, clientID, clientSecret, authResponse)
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	fileContent, err := os.ReadFile(homeDir + "/.flagship/credentials/" + product + "/" + username + ".yaml")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expectedContent := fmt.Sprintf(`client_id: %s
client_secret: %s
refresh_token: %s
scope: %s
token: %s
username: %s
`, clientID, clientSecret, refreshToken, scope, accessToken, username)

	assert.Equal(t, expectedContent, string(fileContent))

}

func TestReadAuth(t *testing.T) {
	err := CreateAuthFile(product, username, clientID, clientSecret, authResponse)
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	v, err := ReadAuth(product, username)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	assert.Equal(t, v.GetString("client_id"), clientID)
	assert.Equal(t, v.GetString("client_secret"), clientSecret)
	assert.Equal(t, v.GetString("username"), username)
	assert.Equal(t, v.GetString("token"), authResponse.AccessToken)
	assert.Equal(t, v.GetString("refresh_token"), authResponse.RefreshToken)
}

func TestSelectAuth(t *testing.T) {
	err := CreateAuthFile(product, username, clientID, clientSecret, authResponse)
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	err = SelectAuth(product, username)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	filepath, err := CredentialPath(product, utils.HOME_CLI)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	assert.Equal(t, string(yamlFile), "current_used_credential: test_user\n")
}

func TestSetAccountID(t *testing.T) {
	err := CreateAuthFile(product, username, clientID, clientSecret, authResponse)
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	err = SelectAuth(product, username)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	err = SetAccountID(product, accountID)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	filepath, err := CredentialPath(product, utils.HOME_CLI)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	assert.Equal(t, string(yamlFile), "account_id: account_id\ncurrent_used_credential: test_user\n")
}

func TestSetAccountEnvironmentID(t *testing.T) {
	err := CreateAuthFile(product, username, clientID, clientSecret, authResponse)
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	err = SelectAuth(product, username)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	err = SetAccountEnvID(product, accountEnvID)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	filepath, err := CredentialPath(product, utils.HOME_CLI)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	assert.Equal(t, string(yamlFile), "account_environment_id: account_environment_id\ncurrent_used_credential: test_user\n")
}

func TestReadCredentialsFromFile(t *testing.T) {
	err := CreateAuthFile(product, username, clientID, clientSecret, authResponse)
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	filepath, err := CredentialPath(product, username)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	v, err := ReadCredentialsFromFile(filepath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	assert.Equal(t, v.GetString("client_id"), clientID)
	assert.Equal(t, v.GetString("client_secret"), clientSecret)
	assert.Equal(t, v.GetString("username"), username)
	assert.Equal(t, v.GetString("token"), authResponse.AccessToken)
	assert.Equal(t, v.GetString("refresh_token"), authResponse.RefreshToken)
}

func TestRewriteToken(t *testing.T) {
	err := CreateAuthFile(product, username, clientID, clientSecret, models.TokenResponse{})
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	err = RewriteToken(product, username, authResponse)
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	v, err := ReadAuth(product, username)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	assert.Equal(t, v.GetString("client_id"), clientID)
	assert.Equal(t, v.GetString("client_secret"), clientSecret)
	assert.Equal(t, v.GetString("username"), username)
	assert.Equal(t, v.GetString("token"), authResponse.AccessToken)
	assert.Equal(t, v.GetString("refresh_token"), authResponse.RefreshToken)
}
