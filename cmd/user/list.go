/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
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
	Short: "this list variation group",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpListUsers()
	},
}

func init() {
	UserCmd.AddCommand(listCmd)
}
