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
	"github.com/spf13/viper"
)

var (
	credentialsFile string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [-n <name> | --name=<name>] [-i <clientID> | --client-id=<clientID>] [-s <clientSecret> | --client-secret=<clientSecret>] [-a <accountID> | --account-id=<account-id>] [-e <accountEnvironmentID> | --account-environment-id=<accountEnvironmentID>]",
	Short: "Create a configuration",
	Long:  `Create a configuration based on the credentials`,
	Run: func(cmd *cobra.Command, args []string) {

		existingConfigurationsName, err := config.GetConfigurationsName()
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		if credentialsFile != "" {
			config.ReadCredentialsFromFile(credentialsFile)
			if slices.Contains(existingConfigurationsName, viper.GetString("name")) {
				fmt.Fprintln(cmd.OutOrStdout(), "Configuration name already exists")
				return
			}
			config.CreateConfigurationFile(viper.GetString("name"), viper.GetString("client_id"), viper.GetString("client_secret"), viper.GetString("account_id"), viper.GetString("account_environment_id"))
			fmt.Fprintln(cmd.OutOrStdout(), "Configuration created successfully")
			return
		}
		if ConfigurationName != "" && ConfigurationClientID != "" && ConfigurationClientSecret != "" && ConfigurationAccountID != "" && ConfigurationAccountEnvID != "" {

			if slices.Contains(existingConfigurationsName, ConfigurationName) {
				fmt.Fprintln(cmd.OutOrStdout(), "Configuration name already exists")
				return
			}

			config.CreateConfigurationFile(ConfigurationName, ConfigurationClientID, ConfigurationClientSecret, ConfigurationAccountID, ConfigurationAccountEnvID)
			fmt.Fprintln(cmd.OutOrStdout(), "Configuration created successfully")
			return
		}

		fmt.Fprintln(cmd.OutOrStdout(), "Configuration not created, required fields (name, client ID, client secret, account ID, account environment ID)")

	},
}

func init() {

	createCmd.Flags().StringVarP(&ConfigurationName, "name", "n", "", "configuration name")
	createCmd.Flags().StringVarP(&ConfigurationClientID, "client-id", "i", "", "client ID of a configuration")
	createCmd.Flags().StringVarP(&ConfigurationClientSecret, "client-secret", "s", "", "client secret of a configuration")
	createCmd.Flags().StringVarP(&ConfigurationAccountID, "account-id", "a", "", "account ID of a configuration")
	createCmd.Flags().StringVarP(&ConfigurationAccountEnvID, "account-environment-id", "e", "", "account environment ID of a configuration")

	createCmd.Flags().StringVarP(&credentialsFile, "path", "p", "", "config file to create")

	ConfigurationCmd.AddCommand(createCmd)
}
