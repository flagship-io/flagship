/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"github.com/spf13/cobra"
)

var AuthorizationCmd = &cobra.Command{
	Use:     "auth [login|logout|check]",
	Short:   "auth short desc",
	Aliases: []string{"au"},
	Long:    `auth long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
