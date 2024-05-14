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
	accountId       string
)

// createCmd represents the create command
var loginCmd = &cobra.Command{
	Use:   "login [--credential-file] | [-u <username> | --username=<username>] [-i <clientID> | --client-id=<clientID>] [-s <clientSecret> | --client-secret=<clientSecret>] [-a <accountID> | --account-id=<accountID>]",
	Short: "login",
	Long:  `login`,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.CheckSingleFlag(credentialsFile != "", Username != "") {
			log.Fatalf("error occurred: %s", "1 flag is required. (browser, username)")
		}

		if credentialsFile != "" {
			v, err := config.ReadCredentialsFromFile(credentialsFile)
			if err != nil {
				log.Fatalf("error occurred: %v", err)
			}

			if v.GetString("username") == "" || v.GetString("client_id") == "" || v.GetString("client_secret") == "" || v.GetString("account_id") == "" {
				fmt.Fprintln(cmd.OutOrStderr(), "Error while login, required fields (username, client ID, client secret, account id)")
				return
			}

			authenticationResponse, err := common.HTTPCreateTokenWE(v.GetString("client_id"), v.GetString("client_secret"), v.GetString("account_id"))
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "error occurred: %s", err)
				return
			}

			err = config.CreateAuthFile(utils.WEB_EXPERIMENTATION, v.GetString("username"), v.GetString("client_id"), v.GetString("client_secret"), authenticationResponse)
			if err != nil {
				log.Fatalf("error occurred: %v", err)
			}

			err = config.SelectAuth(utils.WEB_EXPERIMENTATION, v.GetString("username"))
			if err != nil {
				log.Fatalf("error occurred: %v", err)
			}

			err = config.SetAccountID(utils.WEB_EXPERIMENTATION, v.GetString("account_id"))
			if err != nil {
				log.Fatalf("error occurred: %s", err)
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Credential created successfully")
			return
		}

		if Username != "" {
			existingCredentials, err := config.GetUsernames(utils.WEB_EXPERIMENTATION)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "error occurred: %s", err)
				return
			}

			if slices.Contains(existingCredentials, Username) {
				if accountId != "" {
					err := config.SelectAuth(utils.WEB_EXPERIMENTATION, Username)
					if err != nil {
						log.Fatalf("error occurred: %v", err)
					}

					err = config.SetAccountID(utils.WEB_EXPERIMENTATION, accountId)
					if err != nil {
						log.Fatalf("error occurred: %s", err)
					}

					fmt.Fprintln(cmd.OutOrStdout(), "Credential changed successfully to "+Username)
					return
				}

				fmt.Fprintln(cmd.OutOrStderr(), "Error while login, required fields (account id)")
				return
			}

			if ClientID == "" || ClientSecret == "" || accountId == "" {
				fmt.Fprintln(cmd.OutOrStderr(), "Error while login, required fields (username, client ID, client secret, account id)")
				return
			}

			authenticationResponse, err := common.HTTPCreateTokenWE(ClientID, ClientSecret, accountId)
			if err != nil {
				fmt.Fprintln(cmd.OutOrStderr(), err)
				return
			}

			if authenticationResponse.AccessToken == "" {
				fmt.Fprintln(cmd.OutOrStderr(), "Error while login, client_id or client_secret not valid")
				return
			}

			err = config.CreateAuthFile(utils.WEB_EXPERIMENTATION, Username, ClientID, ClientSecret, authenticationResponse)
			if err != nil {
				log.Fatalf("error occurred: %v", err)
			}

			err = config.SelectAuth(utils.WEB_EXPERIMENTATION, Username)
			if err != nil {
				log.Fatalf("error occurred: %v", err)
			}

			err = config.SetAccountID(utils.WEB_EXPERIMENTATION, accountId)
			if err != nil {
				log.Fatalf("error occurred: %s", err)
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Credential created successfully")
		}
	},
}

func init() {

	loginCmd.Flags().StringVarP(&Username, "username", "u", "", "username")

	loginCmd.Flags().StringVarP(&ClientID, "client-id", "i", "", "client ID of an auth")
	loginCmd.Flags().StringVarP(&ClientSecret, "client-secret", "s", "", "client secret of an auth")
	loginCmd.Flags().StringVarP(&accountId, "account-id", "a", "", "account id of an auth")
	loginCmd.Flags().StringVarP(&credentialsFile, "credential-file", "p", "", "config file to create")

	AuthCmd.AddCommand(loginCmd)
}
