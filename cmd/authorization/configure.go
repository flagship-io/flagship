/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package authorization

import (
	"fmt"
	"log"

	"github.com/flagship-io/flagship-cli/utils/config"
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

// loginCmd represents the login command
var ConfigureCmd = &cobra.Command{
	Use:   "configure [-i <client-id> | --client-id=<client-id>] [-s <client-secret> | --client-secret=<client-secret>] [-a <account-id> | --account-id=<account-id>] [-e <account-env-id> | --account-environment-id=<account-env-id>]",
	Short: "Set the credentials of your flagship account",
	Long: `Set your credentials (account_id, account_env_id, client_id, client_secret) 
	by managing the file credentials.yaml ($HOME/.flagship/credentials.yaml)
	
	The command create credentials file at $HOME/.flagship/credentials.yaml and write value of flag as key-value pairs.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteOptionals(config.CredentialsFile, config.GrantType, config.Scope, config.Expiration)

		if credentialsFile != "" {
			viper.SetConfigFile(credentialsFile)
			viper.ReadInConfig()
		}

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
			config.WriteCredentials(config.CredentialsFile, clientId, clientSecret, accountId, accountEnvId)
			fmt.Fprintln(cmd.OutOrStdout(), "Credentials wrote successfully")
		}

	},
}

func init() {

	cobra.OnInitialize(initLocalConfig)

	ConfigureCmd.Flags().StringVarP(&clientId, "client-id", "i", "", "client id of your flagship account")
	ConfigureCmd.Flags().StringVarP(&clientSecret, "client-secret", "s", "", "client secret of your flagship account")
	ConfigureCmd.Flags().StringVarP(&accountId, "account-id", "a", "", "account id of your flagship account")
	ConfigureCmd.Flags().StringVarP(&accountEnvId, "account-environment-id", "e", "", "account environment id of managed environment")

	config.Binder(ConfigureCmd)

	ConfigureCmd.PersistentFlags().StringVarP(&credentialsFile, "credentials-file", "", "", "config file (default is $HOME/.flagship/credentials.yaml)")

}

func initLocalConfig() {
	config.InitLocalConfigureConfig(credentialsFile)
}
