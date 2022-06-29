/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"os"

	"github.com/spf13/cobra"
)

// campaignCmd represents the campaign command
var AuthorizationCmd = &cobra.Command{
	Use:     "auth [login|logout|check]",
	Short:   "auth short desc",
	Aliases: []string{"au", "a"},
	Long:    `auth long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// campaignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// campaignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
