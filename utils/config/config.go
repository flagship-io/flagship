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

func SetPathForConfigName(fileName string) (filePath string) {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if _, err := os.Stat(homeDir + "/.flagship/configurations"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(homeDir+"/.flagship/configurations", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	filepath, _ := filepath.Abs(homeDir + "/.flagship/configurations/" + fileName + ".yaml")
	v.SetConfigFile(filepath)

	return filepath
}

func GetConfigurationsName() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	r := regexp.MustCompile(`(?P<ConfigurationName>[^/]+)\.yaml`)
	var fileNames []string

	if _, err := os.Stat(homeDir + "/.flagship/configurations"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(homeDir+"/.flagship/configurations", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

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
		fileNames = append(fileNames, match[configurationName])
	}

	return fileNames, nil
}

func SetOptionalsDefault(grantType, scope string, expiration int) (*Config, error) {
	viper.Set("grant_type", grantType)
	viper.Set("scope", scope)
	viper.Set("expiration", expiration)

	return &Config{viper.GetViper()}, nil
}

func CreateConfigurationFile(configurationName, clientId, clientSecret, accountId, accountEnvId string) (*Config, error) {

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

	return &Config{v}, nil

}

func SelectConfiguration(configurationName string) (*Config, error) {

	filepath := SetPathForConfigName(".cli")
	configFilepath := SetPathForConfigName(configurationName)
	v.SetConfigFile(configFilepath)
	v.MergeInConfig()

	err := v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	viper.SetConfigFile(configFilepath)
	viper.MergeInConfig()

	return &Config{v}, nil

}

func EditConfigurationFile(configurationName, newName, clientId, clientSecret, accountId, accountEnvId string) (*Config, error) {

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

	return &Config{v}, nil

}

func ReadCredentialsFromFile(configurationFile string) *Config {
	viper.SetConfigFile(configurationFile)
	viper.MergeInConfig()

	return &Config{viper.GetViper()}
}

func WriteOptionals(credentialsFile, grantType, scope string, expiration int) (*Config, error) {

	filepath := SetPathForConfigName(credentialsFile)
	v.Set("grant_type", grantType)
	v.Set("scope", scope)
	v.Set("expiration", expiration)
	err := v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	return &Config{v}, nil
}

func InitLocalConfigureConfig(credentialsFile string) *Config {

	if credentialsFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(credentialsFile)
	}

	if err := v.MergeInConfig(); err != nil {
		return &Config{v}
	}

	return &Config{v}
}

func WriteToken(configurationName, token string) (*Config, error) {
	configFilepath := SetPathForConfigName(configurationName)
	filePath := SetPathForConfigName(".cli")

	viper.SetConfigFile(configFilepath)
	viper.MergeInConfig()
	viper.Set("token", token)

	err := viper.WriteConfigAs(filePath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	err = viper.WriteConfigAs(configFilepath)
	if err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	return &Config{viper.GetViper()}, nil
}

func SetViperMock() {
	viper.GetViper().Set("account_id", "account_id")
	viper.GetViper().Set("account_environment_id", "account_environment_id")
	viper.GetViper().Set("client_id", "client_id")
	viper.GetViper().Set("client_secret", "client_secret")
	viper.GetViper().Set("token", "token")
	viper.GetViper().Set("output_format", "json")
}
