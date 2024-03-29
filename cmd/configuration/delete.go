/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package configuration

import (
	"fmt"
	"log"
	"os"

	"github.com/flagship-io/flagship/utils/config"
	"github.com/spf13/cobra"
)

// deleteCmd represents delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [-n <configuration-name> | --name=<configuration-name>]",
	Short: "Delete a configuration",
	Long:  `Delete a configuration`,
	Run: func(cmd *cobra.Command, args []string) {

		config.CheckFlagshipHomeDirectory()

		if err := os.Remove(config.SetPathForConfigName(ConfigurationName)); err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(cmd.OutOrStdout(), "Configuration deleted successfully")
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&ConfigurationName, "name", "n", "", "name of the configuration you want to delete")

	if err := deleteCmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	ConfigurationCmd.AddCommand(deleteCmd)
}
