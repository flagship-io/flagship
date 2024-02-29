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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configurations",
	Long:  `List all configurations`,
	Run: func(cmd *cobra.Command, args []string) {

		var configurations []models.Configuration
		existingConfigurationsName, nil := config.GetConfigurationsName()

		for _, fileName := range existingConfigurationsName {
			if fileName != "" && fileName != ".cli" {
				var configurationYaml models.ConfigurationYaml
				var configuration models.Configuration
				yamlFile, err := os.ReadFile(config.SetPathForConfigName(fileName))
				if err != nil {
					log.Fatalf("error occurred: %s", err)
				}

				// Unmarshal the YAML data into the struct
				err = yaml.Unmarshal(yamlFile, &configurationYaml)
				if err != nil {
					log.Fatalf("error occurred: %s", err)
				}
				if configurationYaml.Name != "" {
					configuration.Name = configurationYaml.Name
					configuration.ClientID = configurationYaml.ClientID
					configuration.ClientSecret = configurationYaml.ClientSecret
					configuration.AccountID = configurationYaml.AccountID
					configuration.AccountEnvironmentID = configurationYaml.AccountEnvironmentID
					configurations = append(configurations, configuration)
				}
			}
		}

		utils.FormatItem([]string{"Name", "ClientID", "ClientSecret", "AccountID", "AccountEnvironmentID"}, configurations, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	ConfigurationCmd.AddCommand(listCmd)
}
