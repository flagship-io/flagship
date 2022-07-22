/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "this edit right",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpManageUsers(DataRaw)
	},
}

func init() {

	editCmd.Flags().StringVarP(&UserEmail, "email", "e", "", "the email")
	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "a", "", "the raw data")

	editCmd.Flags().StringVarP(&UserRole, "role", "r", "", "the role")

	UserCmd.AddCommand(editCmd)
}
