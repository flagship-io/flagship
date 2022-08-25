/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"log"

	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [-n <name> | --n=<name>]",
	Short: "Create a project",
	Long:  `Create a project in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		err := httprequest.HTTPCreateProject(ProjectName)
		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
		log.Println("Project created")
	},
}

func init() {

	createCmd.Flags().StringVarP(&ProjectName, "name", "n", "", "name of the project you want to create")

	if err := createCmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("error occured: %v", err)
	}

	ProjectCmd.AddCommand(createCmd)
}
