/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package token

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

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get the information related to your token",
	Long:  `Get the information related to your token`,
	Run: func(cmd *cobra.Command, args []string) {
		var accountYaml models.AccountYaml
		var credYaml models.AuthYaml

		// Home account
		credPath, err := config.CredentialPath(utils.WEB_EXPERIMENTATION, utils.HOME_CLI)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		homeFile, err := os.ReadFile(credPath)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		err = yaml.Unmarshal(homeFile, &accountYaml)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		// Current Cred
		credPathCurr, err := config.CredentialPath(utils.WEB_EXPERIMENTATION, accountYaml.CurrentUsedCredential)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		currentCredFile, err := os.ReadFile(credPathCurr)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		// Unmarshal the YAML data into the struct
		err = yaml.Unmarshal(currentCredFile, &credYaml)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		body := models.Token{
			ClientID:  credYaml.ClientID,
			AccountID: accountYaml.AccountID,
			Scope:     credYaml.Scope,
		}
		utils.FormatItem([]string{"ClientID", "AccountID", "Scope"}, body, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	TokenCmd.AddCommand(infoCmd)
}
