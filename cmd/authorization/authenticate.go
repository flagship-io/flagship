/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	httprequest "github.com/Chadiii/flagship-mock/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func authenticate(loginClientId, loginClientSecret string) string {
	return "login with client_id: " + loginClientId + ", client_secret: " + loginClientSecret
}

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

var AuthenticateCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "authenticate",
	Long:  `authenticate long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(authenticate(viper.GetViper().GetString("client_id"), viper.GetViper().GetString("client_secret")))
		token, err := httprequest.HttpCreateToken(viper.GetViper().GetString("client_id"), viper.GetViper().GetString("client_secret"), "*", "client_credentials")
		if err != nil {
			log.Fatalf("%s", err)
			return
		}
		//fmt.Println("token: " + token)

		if token == "" {
			fmt.Println("client_id or client_secret not valid")
			return
		} else {
			fmt.Println("Token generated successfully")
		}
		writeToken(token)
	},
}
