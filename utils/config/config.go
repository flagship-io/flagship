package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

var v = viper.New()

func Unset(key string) error {
	configMap := viper.AllSettings()
	delete(configMap, key)
	encodedConfig, _ := json.MarshalIndent(configMap, "", " ")
	err := viper.ReadConfig(bytes.NewReader(encodedConfig))
	if err != nil {
		return err
	}
	viper.WriteConfig()
	return nil
}

func CheckABTastyHomeDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()

	if _, err := os.Stat(homeDir + "/.flagship/credentials/" + utils.FEATURE_EXPERIMENTATION); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(homeDir+"/.flagship/credentials/"+utils.FEATURE_EXPERIMENTATION, os.ModePerm)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}
	}

	if _, err := os.Stat(homeDir + "/.flagship/credentials/" + utils.WEB_EXPERIMENTATION); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(homeDir+"/.flagship/credentials/"+utils.WEB_EXPERIMENTATION, os.ModePerm)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}
	}

	return homeDir, err
}

func CredentialPath(product, username string) string {
	homeDir, _ := CheckABTastyHomeDirectory()
	filepath, err := filepath.Abs(homeDir + "/.flagship/credentials/" + product + "/" + username + ".yaml")
	if err != nil {
		log.Fatalf("error occured: %s", err)
	}
	return filepath
}

func GetUsernames(product string) ([]string, error) {
	homeDir, err := CheckABTastyHomeDirectory()
	r := regexp.MustCompile(`(?P<Username>[^/]+)\.yaml`)
	var fileNames []string

	f, err := os.Open(homeDir + "/.flagship/credentials/" + product)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred: %s", err)
		return nil, err
	}

	files, err := f.Readdir(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occurred: %s", err)
		return nil, err
	}

	for _, v := range files {
		match := r.FindStringSubmatch(v.Name())
		userName := r.SubexpIndex("Username")
		if len(match) == 0 {
			log.Fatalln("Error: File not found")
		}

		fileNames = append(fileNames, match[userName])
	}
	return fileNames, nil
}

func CreateAuthFile(product, username, clientId, clientSecret string, authenticationResponse models.TokenResponse) {
	v := viper.New()
	filepath := CredentialPath(product, username)

	v.Set("username", username)
	v.Set("client_id", clientId)
	v.Set("client_secret", clientSecret)
	v.Set("token", authenticationResponse.AccessToken)
	v.Set("refresh_token", authenticationResponse.RefreshToken)

	err := v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}

func ReadAuth(product, AuthName string) *viper.Viper {
	v := viper.New()
	configFilepath := CredentialPath(product, AuthName)
	if _, err := os.Stat(configFilepath); errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "error occurred: %v \n", err)
	}
	v.SetConfigFile(configFilepath)
	v.MergeInConfig()
	return v
}

func SelectAuth(product, AuthName string) {
	var v = viper.New()

	filepath := CredentialPath(product, utils.HOME_CLI)
	v.Set("current_used_credential", AuthName)

	err := v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

}

func SetAccountID(product, accountID string) {
	var v = viper.New()
	configFilepath := CredentialPath(product, utils.HOME_CLI)
	v.SetConfigFile(configFilepath)
	v.MergeInConfig()

	v.Set("account_id", accountID)

	err := v.WriteConfigAs(configFilepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}

func SetAccountEnvID(product, accountEnvID string) {
	var v = viper.New()
	configFilepath := CredentialPath(product, utils.HOME_CLI)
	v.SetConfigFile(configFilepath)
	v.MergeInConfig()

	v.Set("account_environment_id", accountEnvID)

	err := v.WriteConfigAs(configFilepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}

func ReadCredentialsFromFile(AuthFile string) *viper.Viper {
	var v = viper.New()
	v.SetConfigFile(AuthFile)
	err := v.MergeInConfig()
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	return v
}

func WriteToken(product, AuthName string, authenticationResponse models.TokenResponse) {
	v := viper.New()
	configFilepath := CredentialPath(product, AuthName)

	v.SetConfigFile(configFilepath)

	v.MergeInConfig()
	v.Set("token", authenticationResponse.AccessToken)
	v.Set("refresh_token", authenticationResponse.RefreshToken)

	err := v.WriteConfigAs(configFilepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}
