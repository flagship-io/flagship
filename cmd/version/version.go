/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version will match the tag
var Version = "main"

// VersionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "CLI version",
	Long:  `CLI version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "Flagship CLI version: %s\n", Version)
	},
}
