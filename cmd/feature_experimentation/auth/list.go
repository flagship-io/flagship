/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package auth

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
	Short: "list all auth",
	Long:  `list all auth from your system`,
	Run: func(cmd *cobra.Command, args []string) {

		var configurations []models.Configuration_new
		existingCredentials, err := config.GetUsernames(utils.FEATURE_EXPERIMENTATION)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		for _, fileName := range existingCredentials {
			if fileName != "" {
				var configurationYaml models.ConfigurationYaml_new
				var configuration models.Configuration_new
				yamlFile, err := os.ReadFile(config.SetPathForCredentials(utils.FEATURE_EXPERIMENTATION, fileName))
				if err != nil {
					log.Fatalf("error occurred: %s", err)
				}

				// Unmarshal the YAML data into the struct
				err = yaml.Unmarshal(yamlFile, &configurationYaml)
				if err != nil {
					log.Fatalf("error occurred: %s", err)
				}
				if configurationYaml.Username != "" {
					configuration.Username = configurationYaml.Username
					configuration.ClientID = configurationYaml.ClientID
					configuration.ClientSecret = configurationYaml.ClientSecret
					configurations = append(configurations, configuration)
				}
			}
		}

		utils.FormatItem([]string{"Username", "ClientID", "ClientSecret"}, configurations, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {

	AuthCmd.AddCommand(listCmd)
}
