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

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get current running auth credential for web experimentation",
	Long:  `Get current running auth credential for web experimentation`,
	Run: func(cmd *cobra.Command, args []string) {
		var accountYaml models.AccountYaml
		var account models.AccountJSON

		yamlFile, err := os.ReadFile(config.CredentialPath(utils.WEB_EXPERIMENTATION, utils.HOME_CLI))
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		// Unmarshal the YAML data into the struct
		err = yaml.Unmarshal(yamlFile, &accountYaml)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		account.CurrentUsedCredential = accountYaml.CurrentUsedCredential

		utils.FormatItem([]string{"CurrentUsedCredential"}, account, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {
	AuthCmd.AddCommand(currentCmd)
}
