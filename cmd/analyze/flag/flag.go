/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/
package flag

import (
	"encoding/json"
	"log"

	"github.com/flagship-io/codebase-analyzer/pkg/config"
	"github.com/flagship-io/codebase-analyzer/pkg/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Directory         string
	RepoURL           string
	RepoBranch        string
	NbLineCodeEdges   int
	FilesToExcludes   string
	SearchCustomRegex string
)
var FSConfig *config.Config

// FlagCmd represents the flag command
var FlagCmd = &cobra.Command{
	Use:   "flag [create|list]",
	Short: "Manage analyzed flags",
	Long:  `Manage analyzed flags in your account`,
	PreRun: func(cmd *cobra.Command, args []string) {
		var FilesToExcludes_ []string

		err := json.Unmarshal([]byte(FilesToExcludes), &FilesToExcludes_)
		if err != nil {
			log.Fatalf("error occurred: %s", err)
		}

		FSConfig = &config.Config{
			FlagshipAPIURL:        "https://api.flagship.io",
			FlagshipAPIToken:      viper.GetString("token"),
			FlagshipAccountID:     viper.GetString("account_id"),
			FlagshipEnvironmentID: viper.GetString("account_environment_id"),
			Directory:             Directory,
			RepositoryURL:         RepoURL,
			RepositoryBranch:      RepoBranch,
			NbLineCodeEdges:       NbLineCodeEdges,
			FilesToExcludes:       FilesToExcludes_,
			SearchCustomRegex:     SearchCustomRegex,
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := handler.AnalyzeCode(FSConfig)
		if err != nil {
			log.Fatalf("error occured: %s", err)
		}
	},
}

func init() {
	FlagCmd.PersistentFlags().StringVarP(&Directory, "directory", "", ".", "directory to analyze")
	FlagCmd.PersistentFlags().StringVarP(&RepoURL, "repository-url", "", "https://gitlab.com/org/repo", "repository URL")
	FlagCmd.PersistentFlags().StringVarP(&RepoBranch, "repository-branch", "", "main", "repository branch")
	FlagCmd.PersistentFlags().IntVarP(&NbLineCodeEdges, "code-edge", "", 1, "nombre of line code edges")
	FlagCmd.PersistentFlags().StringVarP(&FilesToExcludes, "file-excludes", "", "[\".git\", \".github\", \".vscode\"]", "nombre of line code edges")
	FlagCmd.PersistentFlags().StringVarP(&SearchCustomRegex, "custom-regex", "", "", "custom regex")
}
