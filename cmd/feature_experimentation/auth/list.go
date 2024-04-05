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

		var auths []models.Auth
		existingAuths, err := config.GetUsernames(utils.FEATURE_EXPERIMENTATION)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		for _, fileName := range existingAuths {
			if fileName != "" {
				var authYaml models.AuthYaml
				var auth models.Auth
				yamlFile, err := os.ReadFile(config.CredentialPath(utils.FEATURE_EXPERIMENTATION, fileName))
				if err != nil {
					log.Fatalf("error occurred: %s", err)
				}

				// Unmarshal the YAML data into the struct
				err = yaml.Unmarshal(yamlFile, &authYaml)
				if err != nil {
					log.Fatalf("error occurred: %s", err)
				}
				if authYaml.Username != "" {
					auth.Username = authYaml.Username
					auth.ClientID = authYaml.ClientID
					auth.ClientSecret = authYaml.ClientSecret
					auth.Token = authYaml.Token
					auth.RefreshToken = authYaml.RefreshToken
					auths = append(auths, auth)
				}
			}
		}

		utils.FormatItem([]string{"Username", "ClientID", "ClientSecret", "Token"}, auths, viper.GetString("output_format"), cmd.OutOrStdout())

	},
}

func init() {

	AuthCmd.AddCommand(listCmd)
}
