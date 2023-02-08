/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package analyse

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/flagship-io/codebase-analyzer/pkg/config"
	"github.com/flagship-io/codebase-analyzer/pkg/handler"
	"github.com/flagship-io/flagship/models"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	directory    string
	repoURL      string
	repoBranch   string
	withCreation bool
)

var flagResponse models.Flag

// analyseCmd represents the analyse command
var AnalyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "Manage your flags",
	Long:  `Manage your flags in your account`,
	Run: func(cmd *cobra.Command, args []string) {

		FSConfig := &config.Config{
			FlagshipAPIURL:        "https://api.flagship.io",
			FlagshipAPIToken:      viper.GetString("token"),
			FlagshipAccountID:     viper.GetString("account_id"),
			FlagshipEnvironmentID: viper.GetString("account_environment_id"),
			Directory:             directory,
			RepositoryURL:         repoURL,
			RepositoryBranch:      repoBranch,
			NbLineCodeEdges:       1,
			FilesToExcludes:       []string{".git", ".github", ".vscode"},
		}

		if withCreation {
			result, err := handler.ExtractFlagsInfo(FSConfig)
			if err != nil {
				log.Fatalf("error occured: %s", err)
			}
			for _, r := range result {
				for _, result := range r.Results {
					flagRequest := models.Flag{
						Name:        result.FlagKey,
						Type:        "string",
						Description: "flag created by CLI",
						Source:      "codebase_analyzer",
					}

					flagRequestJSON, err_ := json.Marshal(flagRequest)
					if err_ != nil {
						log.Fatalf("error occurred: %s", err)
					}

					body, err := httprequest.HTTPCreateFlag(string(flagRequestJSON))
					if err != nil {
						log.Fatalf("error occurred: %v", err)
					}

					err_r := json.Unmarshal(body, &flagResponse)

					if err_r != nil {
						log.Fatalf("error occurred: %v", err)
					}

					if flagResponse.Id != "" {
						log.WithFields(log.Fields{
							"id":           flagResponse.Id,
							"key":          flagResponse.Name,
							"type":         flagResponse.Type,
							"defaultValue": flagResponse.DefaultValue,
							"description":  flagResponse.Description,
							"source":       flagResponse.Source,
						}).Info("Created Flag")
					} else {
						log.WithFields(log.Fields{
							"key": flagRequest.Name,
						}).Error("Existing Flag")
					}

				}
			}
		}

		err := handler.AnalyzeCode(FSConfig)

		if err != nil {
			log.Fatalf("error occured: %s", err)
		}
	},
}

func init() {
	AnalyseCmd.Flags().StringVarP(&directory, "directory", "", ".", "directory")
	AnalyseCmd.Flags().StringVarP(&repoURL, "repository-url", "", "https://gitlab.com/org/repo", "repository URL")
	AnalyseCmd.Flags().StringVarP(&repoBranch, "repository-branch", "", "main", "repository branch")
	AnalyseCmd.Flags().BoolVarP(&withCreation, "with-creation", "", false, "analyse and create flag if not exist")
}
