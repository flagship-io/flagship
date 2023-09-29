/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package configuration

import (
	"log"
	"os"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// getCmd represents get command
var getCmd = &cobra.Command{
	Use:   "get [-n <name> | --name=<name>]",
	Short: "Get a configuration",
	Long:  `Get a configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		var configurationYaml models.ConfigurationYaml
		var configuration models.Configuration

		yamlFile, err := os.ReadFile(config.SetPathForConfigName(ConfigurationName))
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		// Unmarshal the YAML data into the struct
		err = yaml.Unmarshal(yamlFile, &configurationYaml)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		configuration.Name = configurationYaml.Name
		configuration.ClientID = configurationYaml.ClientID
		configuration.ClientSecret = configurationYaml.ClientSecret
		configuration.AccountID = configurationYaml.AccountID
		configuration.AccountEnvironmentID = configurationYaml.AccountEnvironmentID

		utils.FormatItem([]string{"Name", "ClientID", "ClientSecret", "AccountID", "AccountEnvironmentID"}, configuration, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	getCmd.Flags().StringVarP(&ConfigurationName, "name", "n", "", "name of the configuration you want to display")

	if err := getCmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
	ConfigurationCmd.AddCommand(getCmd)
}
