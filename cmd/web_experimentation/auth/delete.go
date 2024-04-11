/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/spf13/cobra"
)

// deleteCmd represents delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [-u <username> | --username=<username>]",
	Short: "Delete an auth",
	Long:  `Delete an auth`,
	Run: func(cmd *cobra.Command, args []string) {

		config.CheckABTastyHomeDirectory()

		if err := os.Remove(config.CredentialPath(utils.WEB_EXPERIMENTATION, Username)); err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		fmt.Fprintln(cmd.OutOrStdout(), "Credential deleted successfully")
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&Username, "username", "u", "", "username of the credentials you want to delete")

	if err := deleteCmd.MarkFlagRequired("username"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	AuthCmd.AddCommand(deleteCmd)
}