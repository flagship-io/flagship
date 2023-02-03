/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package analyse

import (
	"fmt"
	"log"

	"github.com/flagship-io/codebase-analyzer/pkg/config"
	"github.com/flagship-io/codebase-analyzer/pkg/handler"
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

type FlagData struct {
	Id               string   `json:",omitempty"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Description      string   `json:"description"`
	Source           string   `json:"source"`
	DefaultValue     string   `json:",omitempty"`
	PredefinedValues []string `json:",omitempty"`
}

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
					//not working
					body, err := httprequest.HTTPCreateFlag(result.FlagKey)
					if err != nil {
						log.Fatalf("error occurred: %v", err)
					}
					fmt.Println(body)
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
