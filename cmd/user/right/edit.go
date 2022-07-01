/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package right

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	editEmail string
	editRole  string
)

func editRight(email, role string) string {
	return "edit right \n email: " + editEmail + "\n with role: " + editRole + "\n the account id: " + viper.GetViper().GetString("account_id")
}

// createCmd represents the create command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "this edit right",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(editRight(email, role))
	},
}

func init() {

	editCmd.Flags().StringVarP(&editEmail, "email", "e", "", "the email")

	if err := editCmd.MarkFlagRequired("email"); err != nil {
		fmt.Println(err)
	}

	editCmd.Flags().StringVarP(&editRole, "role", "r", "", "the role")

	if err := editCmd.MarkFlagRequired("role"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	RightCmd.AddCommand(editCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
