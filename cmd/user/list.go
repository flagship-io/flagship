/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"log"

	"github.com/Chadiii/flagship/utils"
	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "this list users",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListUsers()
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"Email", "Role"}, body, viper.GetString("output_format"))
	},
}

func init() {
	UserCmd.AddCommand(listCmd)
}
