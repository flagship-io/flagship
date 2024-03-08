/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package user

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest/feature_experimentation"
	"github.com/spf13/cobra"
)

// createCmd represents the add command
var createCmd = &cobra.Command{
	Use:   "create [-d <data-raw> | --data <data-raw>]",
	Short: "Create a user with right",
	Long:  `Create a user with right in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := httprequest.HTTPBatchUpdateUsers(DataRaw)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "users created")
	},
}

func init() {
	createCmd.Flags().StringVarP(&DataRaw, "data-raw", "d", "", "raw data contains all the info to create your user with right, check the doc for details")

	if err := createCmd.MarkFlagRequired("data-raw"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	UserCmd.AddCommand(createCmd)
}
