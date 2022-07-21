/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	email string
	role  string
)

func addRight(email, role string) string {
	return "add right \n email: " + email + "\n with role: " + role
}

// createCmd represents the create command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "this add user & right",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(addRight(email, role))
	},
}

func init() {

	addCmd.Flags().StringVarP(&email, "email", "e", "", "the email")

	if err := addCmd.MarkFlagRequired("email"); err != nil {
		fmt.Println(err)
	}

	addCmd.Flags().StringVarP(&role, "role", "r", "", "the role")

	if err := addCmd.MarkFlagRequired("role"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	UserCmd.AddCommand(addCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
