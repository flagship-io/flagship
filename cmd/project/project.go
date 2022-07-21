/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package project

import (
	"github.com/spf13/cobra"
)

var (
	ProjectId     string
	ProjectName   string
	ProjectStatus string
)

// ProjectCmd represents the project command
var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "project short desc",
	Long:  `project long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
