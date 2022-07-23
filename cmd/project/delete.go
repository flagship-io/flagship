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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this delete project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPDeleteProject(ProjectId)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("project deleted")
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&ProjectId, "id", "i", "", "the project id")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}

	ProjectCmd.AddCommand(deleteCmd)
}
