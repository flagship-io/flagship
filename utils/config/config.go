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
	"github.com/spf13/cobra"
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

func CheckFlagshipHomeDirectory() string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if _, err := os.Stat(homeDir + "/.flagship/configurations"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(homeDir+"/.flagship/configurations", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	return homeDir
}

func SetPathForConfigName(fileName string) string {
	homeDir := CheckFlagshipHomeDirectory()

	filepath, err := filepath.Abs(homeDir + "/.flagship/configurations/" + fileName + ".yaml")
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

	return filepath
}

func GetConfigurationsName() ([]string, error) {
	homeDir := CheckFlagshipHomeDirectory()
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

func SetOptionalsDefault(grantType, scope string, expiration int) {
	viper.Set("grant_type", grantType)
	viper.Set("scope", scope)
	viper.Set("expiration", expiration)
}

func CreateConfigurationFile(configurationName, clientId, clientSecret, accountId, accountEnvId string) {
	filepath := SetPathForConfigName(configurationName)

	v.Set("name", configurationName)
	v.Set("client_id", clientId)
	v.Set("client_secret", clientSecret)
	v.Set("account_id", accountId)
	v.Set("account_environment_id", accountEnvId)

	err := v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

}

func SelectConfiguration(configurationName string) {
	filepath := SetPathForConfigName(".cli")
	v.Set("current_used_configuration", configurationName)

	err := v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	ReadConfiguration(configurationName)
}

func ReadConfiguration(configurationName string) {
	configFilepath := SetPathForConfigName(configurationName)
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
	viper.Set("current_used_configuration", nil)
	Unset("current_used_configuration")

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
