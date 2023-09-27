/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package configuration

import "github.com/spf13/cobra"

var (
	ConfigurationName         string
	ConfigurationClientID     string
	ConfigurationClientSecret string
	ConfigurationAccountID    string
	ConfigurationAccountEnvID string
)

// ConfigurationCmd represents the configuration command
var ConfigurationCmd = &cobra.Command{
	Use:     "configuration [create|edit|get|list|delete|use]",
	Aliases: []string{"conf"},
	Short:   "Manage your CLI configurations",
	Long:    `Manage your CLI configurations in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
