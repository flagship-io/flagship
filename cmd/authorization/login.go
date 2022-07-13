/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship-mock/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	loginClientId     string
	loginClientSecret string
)

func login(loginClientId, loginClientSecret string) string {
	return "login with client_id: " + loginClientId + ", client_secret: " + loginClientSecret
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "this authorization login",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if loginClientId == "" {
			loginClientId = viper.GetViper().GetString("client_id")
		}
		if loginClientSecret == "" {
			loginClientSecret = viper.GetViper().GetString("client_secret")
		}

		if loginClientId == "" || loginClientSecret == "" {

			fmt.Println("required client_id and client_secret")

		} else {
			fmt.Println(login(loginClientId, loginClientSecret))
			httprequest.HttpToken(loginClientId, loginClientSecret, "*", "client_credentials")

		}

	},
}

func init() {

	loginCmd.Flags().StringVarP(&loginClientId, "client_id", "i", "", "the client id")
	loginCmd.Flags().StringVarP(&loginClientSecret, "client_secret", "s", "", "the client secret")

	// Here you will define your flags and configuration settings.
	AuthorizationCmd.AddCommand(loginCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
