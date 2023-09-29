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

// currentCmd represents current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get current configuration",
	Long:  `Get current configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		var configurationYaml models.ConfigurationYaml
		var configuration models.Configuration

		var v = viper.New()
		configFilepath := config.SetPathForConfigName(".cli")
		v.SetConfigFile(configFilepath)
		v.MergeInConfig()

		configurationName := v.GetString("current_used_configuration")

		yamlFile, err := os.ReadFile(config.SetPathForConfigName(configurationName))
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
	ConfigurationCmd.AddCommand(currentCmd)
}
