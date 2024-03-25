/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package account

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

// getCmd represents the list command
var currentCmd = &cobra.Command{
	Use:   "get",
	Short: "get an auth credential",
	Long:  `list an auth credential from your system`,
	Run: func(cmd *cobra.Command, args []string) {

		var configurationYaml models.ConfigurationYaml_new
		var configuration models.Configuration_new

		yamlFile, err := os.ReadFile(config.SetPathForCredentials(utils.FEATURE_EXPERIMENTATION, Username))
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		// Unmarshal the YAML data into the struct
		err = yaml.Unmarshal(yamlFile, &configurationYaml)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		configuration.Username = configurationYaml.Username
		configuration.ClientID = configurationYaml.ClientID
		configuration.ClientSecret = configurationYaml.ClientSecret

		utils.FormatItem([]string{"Username", "ClientID", "ClientSecret"}, configuration, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	currentCmd.Flags().StringVarP(&Username, "username", "u", "", "username of the credentials you want to display")

	if err := currentCmd.MarkFlagRequired("username"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	AccountCmd.AddCommand(currentCmd)
}
