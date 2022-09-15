/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package targeting_key

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
	Short: "List all targeting keys",
	Long:  `List all targeting keys in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListTargetingKey()
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name", "Type", "Description"}, body, viper.GetString("output_format"))
	},
}

func init() {
	TargetingKeyCmd.AddCommand(listCmd)
}