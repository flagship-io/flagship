/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"fmt"

	"github.com/spf13/cobra"
)

func logout() string {
	return "logout from session by deleting token"
}

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "this authorization logout",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(logout())
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	AuthorizationCmd.AddCommand(logoutCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
