/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package authorization

import (
	"fmt"

	"github.com/spf13/cobra"
)

const flagshipVersion = "v0.2.0"

// versionCmd represents the check command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "this show version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Flagship CLI version : %s", flagshipVersion)
	},
}
