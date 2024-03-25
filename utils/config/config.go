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

func CheckFlagshipHomeDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()

	if _, err := os.Stat(homeDir + "/.flagship/credentials/fe"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(homeDir+"/.flagship/credentials/fe", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(homeDir + "/.flagship/credentials/we"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(homeDir+"/.flagship/credentials/we", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	return homeDir, err
}

func SetPathForConfigName(fileName string) string {
	homeDir, _ := CheckFlagshipHomeDirectory()

	filepath, err := filepath.Abs(homeDir + "/.flagship/credentials/fe/" + fileName + ".yaml")
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

	return filepath
}

func SetPathForCredentials(product, fileName string) string {
	homeDir, _ := CheckFlagshipHomeDirectory()
	filepath, err := filepath.Abs(homeDir + "/.flagship/credentials/" + product + "/" + fileName + ".yaml")
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

	return filepath
}

func GetConfigurationsName() ([]string, error) {
	homeDir, _ := CheckFlagshipHomeDirectory()
	r := regexp.MustCompile(`(?P<ConfigurationName>[^/]+)\.yaml`)
	var fileNames []string

	f, err := os.Open(homeDir + "/.flagship/configurations")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, v := range files {
		match := r.FindStringSubmatch(v.Name())
		configurationName := r.SubexpIndex("ConfigurationName")
		if len(match) == 0 {
			log.Fatalln("Error: File not found")
		}

		if match[configurationName] == ".cli" {
			continue
		}

		fileNames = append(fileNames, match[configurationName])
	}
	return fileNames, nil
}

func GetUsernames(product string) ([]string, error) {
	homeDir, err := CheckFlagshipHomeDirectory()
	r := regexp.MustCompile(`(?P<UserName>[^/]+)\.yaml`)
	var fileNames []string

	f, err := os.Open(homeDir + "/.flagship/credentials/" + product)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, v := range files {
		match := r.FindStringSubmatch(v.Name())
		userName := r.SubexpIndex("UserName")
		if len(match) == 0 {
			log.Fatalln("Error: File not found")
		}

		fileNames = append(fileNames, match[userName])
	}
	return fileNames, nil
}

func SetOptionalsDefault(grantType, scope string, expiration int) {
	viper.Set("grant_type", grantType)
	viper.Set("scope", scope)
	viper.Set("expiration", expiration)
}

func CreateConfigurationFile(configurationName, clientId, clientSecret string) {
	filepath := SetPathForConfigName(configurationName)

	v.Set("name", configurationName)
	v.Set("client_id", clientId)
	v.Set("client_secret", clientSecret)

	err := v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

}

func CreateCredentialsFile(product, username, clientId, clientSecret string, authenticationResponse models.TokenResponse) {
	filepath := SetPathForCredentials(product, username)

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

func SelectConfiguration(configurationName string) {
	filepath := SetPathForConfigName(".cli")
	v.Set("current_used_credential", configurationName)

	err := v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	ReadConfiguration(configurationName)
}

func ReadConfiguration(configurationName string) *viper.Viper {
	v := viper.New()
	configFilepath := SetPathForConfigName(configurationName)
	v.SetConfigFile(configFilepath)
	v.MergeInConfig()
	return v
}

func SelectCredentials(product, configurationName string) {
	var v = viper.New()

	filepath := SetPathForCredentials(product, utils.HOME_CLI)
	v.Set("current_used_credential", configurationName)

	err := v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	ReadCredentials(product, configurationName)
}

func ReadCredentials(product, configurationName string) {
	configFilepath := SetPathForCredentials(product, configurationName)
	viper.SetConfigFile(configFilepath)
	viper.MergeInConfig()
}

func EditConfigurationFile(configurationName, newName, clientId, clientSecret, accountId, accountEnvId string) {
	filepath := SetPathForConfigName(configurationName)

	v.Set("name", newName)
	v.Set("client_id", clientId)
	v.Set("client_secret", clientSecret)
	v.Set("account_id", accountId)
	v.Set("account_environment_id", accountEnvId)

	err := v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	e := os.Rename(filepath, SetPathForConfigName(newName))
	if e != nil {
		log.Fatal(e)
	}
}

func SetAccountID(product, accountID string) {
	var v = viper.New()
	configFilepath := SetPathForCredentials(product, utils.HOME_CLI)
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
	configFilepath := SetPathForCredentials(product, utils.HOME_CLI)
	v.SetConfigFile(configFilepath)
	v.MergeInConfig()

	v.Set("account_environment_id", accountEnvID)

	err := viper.WriteConfigAs(configFilepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}

func ReadCredentialsFromFile(configurationFile string) {
	viper.SetConfigFile(configurationFile)
	err := viper.MergeInConfig()

	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}

func WriteToken(configurationName string, authenticationResponse models.TokenResponse) {
	configFilepath := SetPathForConfigName(configurationName)

	viper.SetConfigFile(configFilepath)
	err := viper.MergeInConfig()
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	viper.Set("token", authenticationResponse.AccessToken)
	viper.Set("refresh_token", authenticationResponse.RefreshToken)
	viper.Set("current_used_credential", nil)
	Unset("current_used_credential")

	err = viper.WriteConfigAs(configFilepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}

func WriteToken_new(product, username string) {
	configFilepath := SetPathForCredentials(product, username)

	viper.SetConfigFile(configFilepath)
	err := viper.MergeInConfig()
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	viper.Set("token", "token")
	viper.Set("refresh_token", "refresh_token")
	viper.Set("current_used_credential", nil)
	Unset("current_used_credential")

	err = viper.WriteConfigAs(configFilepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}

func SetViperMock() {
	viper.GetViper().Set("account_id", "account_id")
	viper.GetViper().Set("account_environment_id", "account_environment_id")
	viper.GetViper().Set("client_id", "client_id")
	viper.GetViper().Set("client_secret", "client_secret")
	viper.GetViper().Set("token", "access_token")
	viper.GetViper().Set("refresh_token", "refresh_token")
	viper.GetViper().Set("output_format", "json")
}
