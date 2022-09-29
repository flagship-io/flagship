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

// createCmd represents the create command
var deleteCmd = &cobra.Command{
	Use:   "delete [-e <email> | --email <email>]",
	Short: "Delete a user",
	Long:  `Delete a user in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPDeleteUsers(UserEmail)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Email deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&UserEmail, "email", "e", "", "email you want to delete")

	if err := deleteCmd.MarkFlagRequired("email"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	UserCmd.AddCommand(deleteCmd)
}
