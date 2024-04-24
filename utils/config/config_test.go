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
)
var authResponse = models.TokenResponse{
	AccessToken:  accessToken,
	RefreshToken: refreshToken,
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

	// Save initial state
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	filepath, err := CredentialPath(product, username)
	if err != nil {
		t.Errorf("CredentialPath() error = %v", err)
	}

	// Assertions
	expectedPath := homeDir + "/.flagship/credentials/" + product + "/" + username + ".yaml"
	assert.Equal(t, expectedPath, filepath)

}

func TestGetUsernames(t *testing.T) {

	// Create authentication response

	// Execute the function
	CreateAuthFile(product, username, clientID, clientSecret, authResponse)

	// Execute the function
	fileNames, err := GetUsernames(product)
	if err != nil {
		t.Errorf("GetUsernames() error = %v", err)
	}

	// Assertions
	if len(fileNames) != 1 || fileNames[0] != "test_user" {
		t.Errorf("GetUsernames() returned unexpected file names: %v", fileNames)
	}
}

func TestCreateAuthFile(t *testing.T) {

	// Save initial state
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get user home directory: %v", err)
	}

	// Execute the function
	CreateAuthFile(product, username, clientID, clientSecret, authResponse)

	// Read the content of the created file
	fileContent, err := os.ReadFile(homeDir + "/.flagship/credentials/" + product + "/" + username + ".yaml")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Assert the content of the file
	expectedContent := fmt.Sprintf(`client_id: %s
client_secret: %s
refresh_token: %s
token: %s
username: %s
`, clientID, clientSecret, refreshToken, accessToken, username)

	assert.Equal(t, expectedContent, string(fileContent))

}
