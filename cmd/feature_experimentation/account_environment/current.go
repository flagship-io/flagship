/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package accountenvironment

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
	Long:  `current an auth credential`,
	Run: func(cmd *cobra.Command, args []string) {

		var accountYaml models.AccountYaml
		var account models.AccountJSON

		credPath, err := config.CredentialPath(utils.FEATURE_EXPERIMENTATION, utils.HOME_CLI)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		yamlFile, err := os.ReadFile(credPath)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		// Unmarshal the YAML data into the struct
		err = yaml.Unmarshal(yamlFile, &accountYaml)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		account.CurrentUsedCredential = accountYaml.CurrentUsedCredential
		account.AccountEnvironmentID = accountYaml.AccountEnvironmentID

		utils.FormatItem([]string{"CurrentUsedCredential", "AccountEnvironmentID"}, account, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	AccountEnvironmentCmd.AddCommand(currentCmd)
}
