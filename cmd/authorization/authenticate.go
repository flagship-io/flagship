/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"log"

	"github.com/flagship-io/flagship/utils/config"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	grantType  string
	scope      string
	expiration int
)

// AuthenticateCmd represents the authenticate command
var AuthenticateCmd = &cobra.Command{
	Use:   "authenticate [--grant-type=<grant-type>] [--scope=<scope>] [--expiration=<expiration>]",
	Short: "Generate your access token",
	Long:  `Generate the access token based on your credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		config.SetOptionalsDefault(grantType, scope, expiration)
		token, err := httprequest.HTTPCreateToken(viper.GetString("client_id"), viper.GetString("client_secret"), grantType, scope, expiration)
		if err != nil {
			log.Fatalf("%s", err)
			return
		}
		if token == "" {
			log.Fatal("client_id or client_secret not valid")
			return
		} else {
			log.Println("Token generated successfully")
		}
		config.WriteToken(config.CredentialsFile, token)
	},
}

func init() {

	AuthenticateCmd.Flags().StringVarP(&grantType, "grant-type", "", config.GrantType, "grant type of the token, DEFAULT value is client_credentials")
	AuthenticateCmd.Flags().StringVarP(&scope, "scope", "", config.Scope, "scope of the token, DEFAULT value is *")
	AuthenticateCmd.Flags().IntVarP(&expiration, "expiration", "", config.Expiration, "expiration time in second of the token, DEFAULT value is 86400")

}
