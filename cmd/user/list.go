/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package user

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
	Short: "List all users",
	Long:  `List all users in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListUsers()
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"Email", "Role"}, body, viper.GetString("output_format"), cmd.OutOrStdout())
	},
}

func init() {
	UserCmd.AddCommand(listCmd)
}
