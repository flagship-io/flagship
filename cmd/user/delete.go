/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	deleteEmail string
)

func DeleteUser(email string) string {
	return "delete users: \n the email: " + email + "\n account_env_id: " + viper.GetViper().GetString("account_environment_id")
}

// createCmd represents the create command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this ldelete user",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(DeleteUser(deleteEmail))
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&deleteEmail, "email", "e", "", "the email")

	if err := deleteCmd.MarkFlagRequired("email"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	UserCmd.AddCommand(deleteCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
