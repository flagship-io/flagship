/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"fmt"
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
		_, err := httprequest.HTTPBatchUpdateUsers(DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "users created\n")
	},
}

func init() {

	editCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to edit your user, check the doc for details")

	if err := editCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	UserCmd.AddCommand(editCmd)
}
