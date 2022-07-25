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
	Short: "this configure client-id and client-secret and account-id",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if clientId == "" {
			clientId = viper.GetString("client_id")
		}
		if clientSecret == "" {
			clientSecret = viper.GetString("client_secret")
		}
		if accountId == "" {
			accountId = viper.GetString("account_id")
		}
		if accountEnvId == "" {
			accountEnvId = viper.GetString("account_environment_id")
		}

		if clientId == "" || clientSecret == "" || accountId == "" || accountEnvId == "" {

			fmt.Println("required client-id and client-secret and account-id and account-env-id")

		} else {
			writeCredentials(clientId, clientSecret, accountId, accountEnvId)
			fmt.Println("Credentials wrote successfully")
		}

	},
}

func init() {

	cobra.OnInitialize(initLocalConfig)

	ConfigureCmd.Flags().StringVarP(&clientId, "client-id", "i", "", "the client id")
	ConfigureCmd.Flags().StringVarP(&clientSecret, "client-secret", "s", "", "the client secret")
	ConfigureCmd.Flags().StringVarP(&accountId, "account-id", "a", "", "the account id")
	ConfigureCmd.Flags().StringVarP(&accountEnvId, "account-environment-id", "e", "", "the account env id")

	v.BindPFlag("client-id", ConfigureCmd.PersistentFlags().Lookup("client_id"))
	v.BindPFlag("client-secret", ConfigureCmd.PersistentFlags().Lookup("client_secret"))
	v.BindPFlag("account-id", ConfigureCmd.PersistentFlags().Lookup("account_id"))
	v.BindPFlag("account-environment-id", ConfigureCmd.PersistentFlags().Lookup("account_environment_id"))

	ConfigureCmd.PersistentFlags().StringVarP(&credentialsFile, "credentials-file", "", "", "config file (default is $HOME/.flagship/credentials.yaml)")

}

func initLocalConfig() {

	if credentialsFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(credentialsFile)
	}

	v.MergeInConfig()
}
