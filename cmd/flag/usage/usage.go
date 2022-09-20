/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package usage

import (
	"github.com/spf13/cobra"
)

// usageCmd represents the usage command
var UsageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Manage flag usage statistics inside your codebase",
	Long:  `Manage flag usage statistics inside your codebase in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
