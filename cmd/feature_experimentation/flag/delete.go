/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package flag

import (
	"fmt"
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest/feature_experimentation"
	"github.com/spf13/cobra"
)

// deleteCmd represents delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [-i <flag-id> | --id=<flag-id>]",
	Short: "Delete a flag",
	Long:  `Delete a flag in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPDeleteFlag(FlagID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Flag deleted")

	},
}

func init() {
	deleteCmd.Flags().StringVarP(&FlagID, "id", "i", "", "id of the flag you want to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	FlagCmd.AddCommand(deleteCmd)
}
