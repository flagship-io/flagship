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

func login(loginClientId, loginClientSecret string) string {
	return "login with client_id: " + loginClientId + ", client_secret: " + loginClientSecret
}

func writeToken1(token string) {
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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "this authorization login",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(login(viper.GetViper().GetString("client_id"), viper.GetViper().GetString("client_secret")))
		token, err := httprequest.HTTPCreateToken(viper.GetViper().GetString("client_id"), viper.GetViper().GetString("client_secret"), "client_credentials", "*", "0")
		if err != nil {
			log.Fatalf("%s", err)
			return
		}
		fmt.Println("token: " + token)

		if token == "" {
			fmt.Println("required valid client_id and client_secret")
			return
		}
		writeToken1(token)

	},
}

func init() {

	// Here you will define your flags and configuration settings.
	AuthorizationCmd.AddCommand(loginCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
