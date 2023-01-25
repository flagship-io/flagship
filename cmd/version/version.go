/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version will match the tag
var Version = "v0.3.6"

// VersionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "version short desc",
	Long:  `version long desc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "Flagship CLI version: %s\n", Version)
	},
}
