/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	clientId        string
	clientSecret    string
	accountId       string
	credentialsFile string
	accountEnvId    string
)

var v = viper.New()

func configure(clientId, clientSecret, accountId, accountEnvId string) string {
	return "login with client_id: " + clientId + ", client_secret: " + clientSecret + ", account id: " + accountId + ", account env id: " + accountEnvId
}

func writeCredentials(clientId, clientSecret, accountId, accountEnvId string) {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if _, err := os.Stat(homeDir + "/.flagship"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(homeDir+"/.flagship", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	filepath, _ := filepath.Abs(homeDir + "/.flagship/credentials.yaml")
	v.SetConfigFile(filepath)
	v.Set("client_id", clientId)
	v.Set("client_secret", clientSecret)
	v.Set("account_id", accountId)
	v.Set("account_environment_id", accountEnvId)
	dir_err := v.WriteConfigAs(filepath)
	if dir_err != nil {
		fmt.Println(dir_err)
	}

}

// loginCmd represents the login command
var ConfigureCmd = &cobra.Command{
	Use:   "configure",
	Short: "this configure client_id and client_secret and account_id",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if clientId == "" {
			clientId = viper.GetViper().GetString("client_id")
		}
		if clientSecret == "" {
			clientSecret = viper.GetViper().GetString("client_secret")
		}
		if accountId == "" {
			accountId = viper.GetViper().GetString("account_id")
		}
		if accountEnvId == "" {
			accountEnvId = viper.GetViper().GetString("account_environment_id")
		}

		if clientId == "" || clientSecret == "" || accountId == "" {

			fmt.Println("required client_id and client_secret")

		} else {
			fmt.Println(configure(clientId, clientSecret, accountId, accountEnvId))
			//httprequest.HttpCreateToken(loginClientId, loginClientSecret, "*", "client_credentials")
			writeCredentials(clientId, clientSecret, accountId, accountEnvId)
		}

	},
}

func init() {

	cobra.OnInitialize(initLocalConfig)

	ConfigureCmd.Flags().StringVarP(&clientId, "client_id", "i", "", "the client id")
	ConfigureCmd.Flags().StringVarP(&clientSecret, "client_secret", "s", "", "the client secret")
	ConfigureCmd.Flags().StringVarP(&accountId, "account_id", "", "", "the account id")
	ConfigureCmd.Flags().StringVarP(&accountEnvId, "account_environment_id", "", "", "the account env id")

	v.BindPFlag("client_id", ConfigureCmd.PersistentFlags().Lookup("client_id"))
	v.BindPFlag("client_secret", ConfigureCmd.PersistentFlags().Lookup("client_secret"))
	v.BindPFlag("account_id", ConfigureCmd.PersistentFlags().Lookup("account_id"))
	v.BindPFlag("account_environment_id", ConfigureCmd.PersistentFlags().Lookup("account_environment_id"))

	ConfigureCmd.PersistentFlags().StringVarP(&credentialsFile, "credentials_file", "", "", "config file (default is $HOME/.flagship/credentials.yaml)")

}

func initLocalConfig() {

	if credentialsFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(credentialsFile)
	}

	v.MergeInConfig()
}
