/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "create",
	Short: "this create user with right",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpManageUsers(DataRaw)
	},
}

func init() {

	addCmd.Flags().StringVarP(&UserEmail, "email", "e", "", "the email")
	addCmd.Flags().StringVarP(&DataRaw, "data-raw", "a", "", "the raw data")

	addCmd.Flags().StringVarP(&UserRole, "role", "r", "", "the role")

	UserCmd.AddCommand(addCmd)

}
