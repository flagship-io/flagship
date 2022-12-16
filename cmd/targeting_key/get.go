/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package targetingkey

import (
	"log"

	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents get command
var getCmd = &cobra.Command{
	Use:   "get [-i <targeting-key-id> | --id=<targeting-key-id>]",
	Short: "Get a targeting key",
	Long:  `Get a targeting key in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPGetTargetingKey(TargetingKeyID)
		if err != nil {
			log.Fatalf("error occurred: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name", "Type", "Description"}, body, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	getCmd.Flags().StringVarP(&TargetingKeyID, "id", "i", "", "id of the targeting key you want to display")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	TargetingKeyCmd.AddCommand(getCmd)
}
