/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package user

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/http_request"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var deleteCmd = &cobra.Command{
	Use:   "delete [-e <email> | --email <email>]",
	Short: "Delete a user",
	Long:  `Delete a user`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.UserRequester.HTTPDeleteUsers(UserEmail)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Email deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&UserEmail, "email", "e", "", "email you want to delete")

	if err := deleteCmd.MarkFlagRequired("email"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	UserCmd.AddCommand(deleteCmd)
}
