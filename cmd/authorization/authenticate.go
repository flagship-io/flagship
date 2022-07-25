/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	grantType  string
	scope      string
	expiration string
)

func writeToken(token string) {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	filepath, _ := filepath.Abs(homeDir + "/.flagship/credentials.yaml")
	viper.SetConfigFile(filepath)
	viper.Set("token", token)
	dir_err := viper.WriteConfigAs(filepath)
	if dir_err != nil {
		fmt.Println(dir_err)
	}
}

// AuthenticateCmd represents the authenticate command
var AuthenticateCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "authenticate shot desc",
	Long:  `authenticate long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := httprequest.HTTPCreateToken(viper.GetString("client_id"), viper.GetString("client_secret"), grantType, scope, expiration)
		if err != nil {
			log.Fatalf("%s", err)
			return
		}

		if token == "" {
			fmt.Println("client_id or client_secret not valid")
			return
		} else {
			fmt.Println("Token generated successfully")
		}
		writeToken(token)
	},
}

func init() {

	AuthenticateCmd.Flags().StringVarP(&grantType, "grant-type", "", "client_credentials", "the grant type")
	AuthenticateCmd.Flags().StringVarP(&scope, "scope", "", "*", "the scope")
	AuthenticateCmd.Flags().StringVarP(&expiration, "expiration", "", "0", "the expiration")

}
