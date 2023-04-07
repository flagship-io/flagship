/*
Copyright © 2022 Flagship Team flagship@abtasty.com

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
	Use:   "project [create|edit|get|list|delete|switch]",
	Short: "Manage your projects",
	Long:  `Manage your projects`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
