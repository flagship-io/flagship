/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "this edit right",
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

	editCmd.Flags().StringVarP(&UserEmail, "email", "e", "", "the email")
	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "a", "", "the raw data")

	editCmd.Flags().StringVarP(&UserRole, "role", "r", "", "the role")

	UserCmd.AddCommand(editCmd)
}
