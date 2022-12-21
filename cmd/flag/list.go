/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package flag

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all flags",
	Long:  `List all flags in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListFlag()
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name", "Type", "Description", "Source"}, body, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	FlagCmd.AddCommand(listCmd)
}
