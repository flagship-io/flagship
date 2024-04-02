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
	Use:   "current",
	Short: "current an auth credential",
	Long:  `current an auth credential from your system`,
	Run: func(cmd *cobra.Command, args []string) {

		var configurationYaml models.AccountYaml
		var configuration models.AccountJSON

		yamlFile, err := os.ReadFile(config.CredentialPath(utils.WEB_EXPERIMENTATION, utils.HOME_CLI))
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		// Unmarshal the YAML data into the struct
		err = yaml.Unmarshal(yamlFile, &configurationYaml)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		configuration.CurrentUsedCredential = configurationYaml.CurrentUsedCredential
		configuration.AccountID = configurationYaml.AccountID
		configuration.AccountEnvironmentID = configurationYaml.AccountEnvironmentID

		utils.FormatItem([]string{"CurrentUsedCredential", "AccountID", "AccountEnvironmentID"}, configuration, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	AccountCmd.AddCommand(currentCmd)
}
