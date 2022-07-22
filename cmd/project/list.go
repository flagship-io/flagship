/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "this list project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpListProject()
	},
}

func init() {
	ProjectCmd.AddCommand(listCmd)
}
