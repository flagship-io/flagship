/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"github.com/spf13/cobra"
)

var (
	UserEmail string
	UserRole  string
	DataRaw   string
)

// userCmd represents the user command
var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "user short desc",
	Long:  `user long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
