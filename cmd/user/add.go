/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "create",
	Short: "this create user with right",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPBatchUpdateUsers(DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("%s", body)
	},
}

func init() {

	addCmd.Flags().StringVarP(&UserEmail, "email", "e", "", "the email")
	addCmd.Flags().StringVarP(&DataRaw, "data-raw", "a", "", "the raw data")

	addCmd.Flags().StringVarP(&UserRole, "role", "r", "", "the role")

	UserCmd.AddCommand(addCmd)

}
