/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"log"
	"strconv"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	grantType  string
	scope      string
	expiration string
)

func setOptionalsDefault(grantType, scope string, expiration int) {
	viper.Set("grant_type", grantType)
	viper.Set("scope", scope)
	viper.Set("expiration", expiration)
}

// AuthenticateCmd represents the authenticate command
var AuthenticateCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "authenticate shot desc",
	Long:  `authenticate long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		exp, err := strconv.Atoi(expiration)
		if err != nil {
			log.Fatalf("%s", err)
			return
		}
		setOptionalsDefault(grantType, scope, exp)
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
		httprequest.WriteToken(token)
	},
}

func init() {

	AuthenticateCmd.Flags().StringVarP(&grantType, "grant-type", "", "client_credentials", "the grant type")
	AuthenticateCmd.Flags().StringVarP(&scope, "scope", "", "*", "the scope")
	AuthenticateCmd.Flags().StringVarP(&expiration, "expiration", "", "0", "the expiration")

}
