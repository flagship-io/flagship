/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"log"

	"github.com/Chadiii/flagship/utils"
	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "this list project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPListProject()
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		utils.FormatItem([]string{"ID", "Name"}, body, viper.GetString("output_format"))
	},
}

func init() {
	ProjectCmd.AddCommand(listCmd)
}
