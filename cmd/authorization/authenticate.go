/*
Copyright © 2022 Flagship Team flagship@abtasty.com

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

func setOptionalsDefault(grantType, scope, expiration string) {
	exp, err := strconv.Atoi(expiration)
	if err != nil {
		log.Fatalf("%s", err)
	}
	viper.Set("grant_type", grantType)
	viper.Set("scope", scope)
	viper.Set("expiration", exp)
}

// AuthenticateCmd represents the authenticate command
var AuthenticateCmd = &cobra.Command{
	Use:   "authenticate [--grant-type=<grant-type>] [--scope=<scope>] [--expiration=<expiration>]",
	Short: "Generate your access token",
	Long:  `Generate the access token based on your credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		setOptionalsDefault(grantType, scope, expiration)
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

	AuthenticateCmd.Flags().StringVarP(&grantType, "grant-type", "", "client_credentials", "grant type of the token, DEFAULT value is client_credentials")
	AuthenticateCmd.Flags().StringVarP(&scope, "scope", "", "*", "scope of the token, DEFAULT value is *")
	AuthenticateCmd.Flags().StringVarP(&expiration, "expiration", "", "86400", "expiration time in second of the token, DEFAULT value is 86400")

}
