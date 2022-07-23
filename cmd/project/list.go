/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"log"

	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
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
		log.Printf("%s", body)
	},
}

func init() {
	ProjectCmd.AddCommand(listCmd)
}
