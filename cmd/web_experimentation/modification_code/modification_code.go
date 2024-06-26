/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package modification_code

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
var ModificationID string
var CreateFile bool

// ModificationCodeCmd represents the variation global code command
var ModificationCodeCmd = &cobra.Command{
	Use:     "modification-code [get|push]",
	Short:   "Get modification code",
	Aliases: []string{"mc"},
	Long:    `Get modification code`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initGlobalCodeDir)
	ModificationCodeCmd.PersistentFlags().StringVarP(&WorkingDir, "working-dir", "", utils.DefaultGlobalCodeWorkingDir(), "Directory where the file will be generated and pushed from")
}

func initConfig() {
	v := viper.New()

	homeDir, _ := os.UserHomeDir()

	v.BindPFlag("working_dir", ModificationCodeCmd.PersistentFlags().Lookup("working-dir"))

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
