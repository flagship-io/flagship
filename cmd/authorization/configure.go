/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"errors"
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
	err = v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

}

func writeOptionals(grantType, scope string, expiration int) {
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
	v.Set("grant_type", grantType)
	v.Set("scope", scope)
	v.Set("expiration", expiration)
	err = v.WriteConfigAs(filepath)
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}

}

// loginCmd represents the login command
var ConfigureCmd = &cobra.Command{
	Use:   "configure [-i <client-id> | --client-id=<client-id>] [-s <client-secret> | --client-secret=<client-secret>] [-a <account-id> | --account-id=<account-id>] [-e <account-env-id> | --account-environment-id=<account-env-id>]",
	Short: "Set the credentials of your flagship account",
	Long: `Set your credentials (account_id, account_env_id, client_id, client_secret) 
	by managing the file credentials.yaml ($HOME/.flagship/credentials.yaml)
	
	The command create credentials file at $HOME/.flagship/credentials.yaml and write value of flag as key-value pairs.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		writeOptionals("client_credentials", "*", 86400)
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
			log.Fatal("required client-id and client-secret and account-id and account-env-id")
		} else {
			writeCredentials(clientId, clientSecret, accountId, accountEnvId)
			log.Println("Credentials wrote successfully")
		}

	},
}

func init() {

	cobra.OnInitialize(initLocalConfig)

	ConfigureCmd.Flags().StringVarP(&clientId, "client-id", "i", "", "client id of your flagship account")
	ConfigureCmd.Flags().StringVarP(&clientSecret, "client-secret", "s", "", "client secret of your flagship account")
	ConfigureCmd.Flags().StringVarP(&accountId, "account-id", "a", "", "account id of your flagship account")
	ConfigureCmd.Flags().StringVarP(&accountEnvId, "account-environment-id", "e", "", "account environment id of managed environment")

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
