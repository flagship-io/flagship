/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package auth

import (
	"fmt"
	"log"
	"slices"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/spf13/cobra"
)

var (
	credentialsFile string
)

func checkSingleFlag(bool1, bool2 bool) bool {
	count := 0
	if bool1 {
		count++
	}
	if bool2 {
		count++
	}

	return count == 1
}

// createCmd represents the create command
var loginCmd = &cobra.Command{
	Use:   "login [--credential-file] | [-u <username> | --username=<username>] [-i <clientID> | --client-id=<clientID>] [-s <clientSecret> | --client-secret=<clientSecret>]",
	Short: "login",
	Long:  `login`,
	Run: func(cmd *cobra.Command, args []string) {
		if !checkSingleFlag(credentialsFile != "", Username != "") {
			log.Fatalf("error occurred: %s", "1 flag is required. (browser, username, credential-file, email)")
		}

		if credentialsFile != "" {
			v := config.ReadCredentialsFromFile(credentialsFile)
			authenticationResponse, err := common.HTTPCreateTokenFE(v.GetString("client_id"), v.GetString("client_secret"), v.GetString("account_id"))
			if err != nil {
				log.Fatalf("%s", err)
				return
			}
			config.CreateAuthFile(utils.FEATURE_EXPERIMENTATION, v.GetString("username"), v.GetString("client_id"), v.GetString("client_secret"), authenticationResponse)
			config.SelectAuth(utils.FEATURE_EXPERIMENTATION, v.GetString("username"))
			fmt.Fprintln(cmd.OutOrStdout(), "Credential created successfully")
			return
		}

		if Username != "" {
			existingCredentials, err := config.GetUsernames(utils.FEATURE_EXPERIMENTATION)
			if err != nil {
				log.Fatalf("error occurred: %s", err)
			}
			if slices.Contains(existingCredentials, Username) {
				config.SelectAuth(utils.FEATURE_EXPERIMENTATION, Username)
				config.SetAccountID(utils.FEATURE_EXPERIMENTATION, AccountId)

				fmt.Fprintln(cmd.OutOrStdout(), "Auth changed successfully to "+Username)
				return
			}

			if ClientID == "" && ClientSecret == "" && AccountId == "" {
				fmt.Fprintln(cmd.OutOrStderr(), "Error while login, required fields (username, client ID, client secret, account id)")
				return
			}
			authenticationResponse, err := common.HTTPCreateTokenFE(ClientID, ClientSecret, AccountId)
			if err != nil {
				log.Fatalf("%s", err)
				return
			}

			if authenticationResponse.AccessToken == "" {
				log.Fatal("client_id or client_secret not valid")
			}
			config.CreateAuthFile(utils.FEATURE_EXPERIMENTATION, Username, ClientID, ClientSecret, authenticationResponse)
			config.SelectAuth(utils.FEATURE_EXPERIMENTATION, Username)
			config.SetAccountID(utils.FEATURE_EXPERIMENTATION, AccountId)

			fmt.Fprintln(cmd.OutOrStdout(), "Credential created successfully")
		}

	},
}

func init() {

	loginCmd.Flags().StringVarP(&Username, "username", "u", "", "auth username")
	loginCmd.Flags().StringVarP(&ClientID, "client-id", "i", "", "client ID of an auth")
	loginCmd.Flags().StringVarP(&ClientSecret, "client-secret", "s", "", "client secret of an auth")
	loginCmd.Flags().StringVarP(&AccountId, "account-id", "a", "", "account id of an auth")

	loginCmd.Flags().StringVarP(&credentialsFile, "credential-file", "p", "", "config file to create")
	AuthCmd.AddCommand(loginCmd)
}
