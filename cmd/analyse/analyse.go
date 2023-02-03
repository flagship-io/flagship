/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package analyse

import (
	"log"

	"github.com/flagship-io/codebase-analyzer/pkg/config"
	"github.com/flagship-io/codebase-analyzer/pkg/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	directory  string
	repoURL    string
	repoBranch string
)

// analyseCmd represents the analyse command
var AnalyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "Manage your flags",
	Long:  `Manage your flags in your account`,
	Run: func(cmd *cobra.Command, args []string) {
		err := handler.AnalyzeCode(&config.Config{
			FlagshipAPIURL:        "https://api.flagship.io",
			FlagshipAPIToken:      viper.GetString("token"),
			FlagshipAccountID:     viper.GetString("account_id"),
			FlagshipEnvironmentID: viper.GetString("account_environment_id"),
			Directory:             directory,
			RepositoryURL:         repoURL,
			RepositoryBranch:      repoBranch,
			NbLineCodeEdges:       1,
			FilesToExcludes:       []string{".git", ".github", ".vscode"},
		})

		if err != nil {
			log.Fatalf("error occured: %s", err)
		}
	},
}

func init() {
	AnalyseCmd.Flags().StringVarP(&directory, "directory", "", ".", "directory")
	AnalyseCmd.Flags().StringVarP(&repoURL, "repository-url", "", "https://gitlab.com/org/repo", "repository URL")
	AnalyseCmd.Flags().StringVarP(&repoBranch, "repository-branch", "", "main", "repository branch")
}
