/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"fmt"
	"log"

	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "this get project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		body, err := httprequest.HTTPGetProject(ProjectId)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Printf("%s", body)
	},
}

func init() {

	getCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "get project by project id")

	if err := getCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
	ProjectCmd.AddCommand(getCmd)

}
