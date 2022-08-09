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
	Use:   "edit [-d <data-raw> | --data-raw=<data-raw>]",
	Short: "Edit a user with right",
	Long:  `Edit a user with right in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPBatchUpdateUsers(DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("%s", body)
	},
}

func init() {

	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to edit your user, check the doc for details")

	UserCmd.AddCommand(editCmd)
}
