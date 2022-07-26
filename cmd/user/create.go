/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the add command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "this create user with right",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := httprequest.HTTPBatchUpdateUsers(DataRaw)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("Users created")
	},
}

func init() {
	createCmd.Flags().StringVarP(&DataRaw, "data-raw", "a", "", "the raw data")
	UserCmd.AddCommand(createCmd)
}
