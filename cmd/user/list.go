/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"log"

	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ListUser() string {
	return "list users:" + "\n account_env_id: " + viper.GetViper().GetString("account_environment_id")
}

// createCmd represents the create command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "this list users",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListUsers()
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("%s", body)
	},
}

func init() {
	UserCmd.AddCommand(listCmd)
}
