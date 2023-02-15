/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package analyze

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

// analyzeCmd represents the analyze command
var AnalyzeCmd = &cobra.Command{
	Use:   "analyze",
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

			_, errListFlag := httprequest.HTTPListFlag()
			if errListFlag != nil {
				log.Fatalf("error occurred: %v", errListFlag)
			}

			result, err := handler.ExtractFlagsInfo(FSConfig)
			if err != nil {
				log.Fatalf("error occured: %s", err)
			}
			for _, r := range result {
				for _, result := range r.Results {

					var flagRequest models.Flag
					var flagResponse models.Flag

					if result.FlagType == "unknown" {
						log.WithFields(log.Fields{
							"key": result.FlagKey,
						}).Error("Type unknown, Flag not created")
						continue
					}

					if result.FlagType == "boolean" {
						flagRequest = models.Flag{
							Name:        result.FlagKey,
							Type:        result.FlagType,
							Description: "flag created by CLI",
							Source:      "codebase_analyzer",
						}
					} else {
						flagRequest = models.Flag{
							Name:         result.FlagKey,
							Type:         result.FlagType,
							DefaultValue: result.FlagDefaultValue,
							Description:  "flag created by CLI",
							Source:       "codebase_analyzer",
						}
					}

					flagRequestJSON, err_ := json.Marshal(flagRequest)
					if err_ != nil {
						log.Fatalf("error occurred: %s", err)
					}

					createdFlag, errCreatedFlag := httprequest.HTTPCreateFlag(string(flagRequestJSON))

					if errCreatedFlag != nil {
						log.Fatalf("error occurred: %v", err)
					}

					err_json := json.Unmarshal(createdFlag, &flagResponse)

					if err_json != nil {
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
						}).Warn("Existing Flag")
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
	AnalyzeCmd.Flags().StringVarP(&directory, "directory", "", ".", "directory")
	AnalyzeCmd.Flags().StringVarP(&repoURL, "repository-url", "", "https://gitlab.com/org/repo", "repository URL")
	AnalyzeCmd.Flags().StringVarP(&repoBranch, "repository-branch", "", "main", "repository branch")
	AnalyzeCmd.Flags().BoolVarP(&withCreation, "with-creation", "", false, "analyze and create flag if not exist")
}
