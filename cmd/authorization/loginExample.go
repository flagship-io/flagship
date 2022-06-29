/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func loginEx(loginClientId, loginClientSecret string) string {
	return "login with id: " + loginClientId + ", secret: " + loginClientSecret
}

// loginCmd represents the login command
var loginExCmd = &cobra.Command{
	Use:   "loginEx",
	Short: "this authorization login",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(loginEx(loginClientId, loginClientSecret))
		fmt.Println(viper.GetViper().GetString("name1"))
		fmt.Println(viper.GetViper().GetString("age"))
		fmt.Println(viper.GetViper().GetString("Hacker"))
		fmt.Println(viper.GetViper().GetStringSlice("hobbies"))
		fmt.Println(viper.GetViper().GetStringMap("clothing"))
		fmt.Println(viper.GetViper().GetString("client_id"))
		fmt.Println(viper.GetViper().GetString("client_secret"))

		/* if viper.GetViper().GetString("client_id") == "8f469786-27da-4390-8c1f-3d1b3367a4bc" && viper.GetViper().GetString("client_secret") == "6c3238675e8f50f5fd90f5267301969ebe89382ac5cb965b85f0489bb0e45500" {
			viper.SetDefault("token", "eyJhbGciOiJIUzUxMiIsImtpZCI6ImMwbjQ4am41dGh2MDFrMGlqbW5nIiwidHlwIjoiSldUIn0.eyJhdWQiOiI4ZjQ2OTc4Ni0yN2RhLTQzOTAtOGMxZi0zZDFiMzM2N2E0YmMiLCJleHAiOjE2NTYxNDcxOTl9.1tccOfqgqHfsjrZAUy0r_tPCTDjNgaLLLV6Jo0rpn5H3vdf76odt1drV2-SMxicoOs3-iWn1-WTXcc3kOpgbCA")
		}

		err := viper.WriteConfigAs("token.yaml")
		if err != nil {
			fmt.Println(err)
		} */
	},
}

func init() {

	loginExCmd.Flags().StringVarP(&loginClientId, "id", "i", "", "the client id")
	loginExCmd.Flags().StringVarP(&loginClientSecret, "secret", "s", "", "the client secret")

	loginExCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.flagship-mock.yaml)")

	if err := loginExCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	AuthorizationCmd.AddCommand(loginExCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
