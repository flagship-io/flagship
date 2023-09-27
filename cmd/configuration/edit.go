/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package configuration

import (
	"fmt"
	"log"
	"slices"

	"github.com/flagship-io/flagship/utils/config"
	"github.com/spf13/cobra"
)

var (
	newConfigurationName string
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [-n <name> | --name=<name>] [-i <clientID> | --client-id=<clientID>] [-s <clientSecret> | --client-secret=<clientSecret>] [-a <accountID> | --account-id=<account-id>] [-e <accountEnvironmentID> | --account-environment-id=<accountEnvironmentID>]",
	Short: "Edit a configuration",
	Long:  `Edit a configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		existingConfigurationsName, err := config.GetConfigurationsName()

		if err != nil {
			log.Fatalf("error occurred : %s", err)
		}

		if !slices.Contains(existingConfigurationsName, ConfigurationName) {
			fmt.Fprintln(cmd.OutOrStdout(), "Configuration doesn't exists")
			return

		}

		if slices.Contains(existingConfigurationsName, newConfigurationName) {
			fmt.Fprintln(cmd.OutOrStdout(), "Configuration already exists")
			return
		}

		conf := config.ReadCredentialsFromFile(config.SetPathForConfigName(ConfigurationName))

		if newConfigurationName == "" {
			newConfigurationName = conf.GetString("name")
		}
		if ConfigurationClientID == "" {
			ConfigurationClientID = conf.GetString("client_id")
		}
		if ConfigurationClientSecret == "" {
			ConfigurationClientSecret = conf.GetString("client_secret")
		}
		if ConfigurationAccountID == "" {
			ConfigurationAccountID = conf.GetString("account_id")
		}
		if ConfigurationAccountEnvID == "" {
			ConfigurationAccountEnvID = conf.GetString("account_environment_id")
		}

		if newConfigurationName == "" && ConfigurationClientID == "" && ConfigurationClientSecret == "" && ConfigurationAccountID == "" && ConfigurationAccountEnvID == "" {
			log.Fatal("required new name or client-id or client-secret or account-id or account-env-id")
			return
		}

		config.EditConfigurationFile(ConfigurationName, newConfigurationName, ConfigurationClientID, ConfigurationClientSecret, ConfigurationAccountID, ConfigurationAccountEnvID)

		fmt.Fprintln(cmd.OutOrStdout(), "Configuration edited successfully")
		return
	},
}

func init() {

	editCmd.Flags().StringVarP(&ConfigurationName, "name", "n", "", "name of the configuration you want to edit")

	editCmd.Flags().StringVarP(&newConfigurationName, "new-name", "", "", "new name for the configuration you want to edit")
	editCmd.Flags().StringVarP(&ConfigurationClientID, "client-id", "i", "", "client ID configuration")
	editCmd.Flags().StringVarP(&ConfigurationClientSecret, "client-secret", "s", "", "client secret of a configuration")
	editCmd.Flags().StringVarP(&ConfigurationAccountID, "account-id", "a", "", "account ID of a configuration")
	editCmd.Flags().StringVarP(&ConfigurationAccountEnvID, "account-environment-id", "e", "", "account environment ID of a configuration")

	if err := editCmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	ConfigurationCmd.AddCommand(editCmd)
}
