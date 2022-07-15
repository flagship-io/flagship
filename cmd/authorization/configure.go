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
var configureCmd = &cobra.Command{
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

	configureCmd.Flags().StringVarP(&clientId, "client_id", "i", "", "the client id")
	configureCmd.Flags().StringVarP(&clientSecret, "client_secret", "s", "", "the client secret")
	configureCmd.Flags().StringVarP(&accountId, "account_id", "", "", "the account id")
	configureCmd.Flags().StringVarP(&accountEnvId, "account_environment_id", "", "", "the account env id")

	v.BindPFlag("client_id", configureCmd.PersistentFlags().Lookup("client_id"))
	v.BindPFlag("client_secret", configureCmd.PersistentFlags().Lookup("client_secret"))
	v.BindPFlag("account_id", configureCmd.PersistentFlags().Lookup("account_id"))
	v.BindPFlag("account_environment_id", configureCmd.PersistentFlags().Lookup("account_environment_id"))

	configureCmd.PersistentFlags().StringVarP(&credentialsFile, "credentials_file", "", "", "config file (default is $HOME/.flagship/credentials.yaml)")

	// Here you will define your flags and configuration settings.
	AuthorizationCmd.AddCommand(configureCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initLocalConfig() {

	if credentialsFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(credentialsFile)
	}

	v.MergeInConfig()
}
