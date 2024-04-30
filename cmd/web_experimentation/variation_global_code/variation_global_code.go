/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package variation_global_code

import (
	"log"
	"os"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var WorkingDir string
var CampaignID string
var VariationID string
var CreateFile bool
var Override bool

// VariationGlobalCodeCmd represents the variation global code command
var VariationGlobalCodeCmd = &cobra.Command{
	Use:     "variation-global-code [get-js | get-css]",
	Short:   "Get variation global code",
	Aliases: []string{"vgc"},
	Long:    `Get variation global code`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initGlobalCodeDir)
	VariationGlobalCodeCmd.PersistentFlags().StringVarP(&WorkingDir, "working-dir", "", utils.DefaultGlobalCodeWorkingDir(), "Directory where the file will be generated and pushed from")
}

func initConfig() {
	v := viper.New()

	homeDir, _ := os.UserHomeDir()

	v.BindPFlag("working_dir", VariationGlobalCodeCmd.PersistentFlags().Lookup("working-dir"))

	v.SetConfigFile(homeDir + "/.flagship/credentials/" + utils.WEB_EXPERIMENTATION + "/.cli.yaml")
	v.MergeInConfig()

	err := v.WriteConfig()
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}
	viper.MergeConfigMap(v.AllSettings())
}

func initGlobalCodeDir() {
	_, err := config.CheckWorkingDirectory(viper.GetString("working_dir"))
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}
}
