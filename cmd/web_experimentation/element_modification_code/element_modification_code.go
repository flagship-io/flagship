/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package element_modification_code

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

// ElementModificationCodeCmd represents the variation global code command
var ElementModificationCodeCmd = &cobra.Command{
	Use:     "element-modification-code [get]",
	Short:   "Get element modification code",
	Aliases: []string{"emc"},
	Long:    `Get element modification code`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initGlobalCodeDir)
	ElementModificationCodeCmd.PersistentFlags().StringVarP(&WorkingDir, "working-dir", "", utils.DefaultGlobalCodeWorkingDir(), "Directory where the file will be generated and pushed from")

}

func initConfig() {
	v := viper.New()

	homeDir, _ := os.UserHomeDir()

	v.BindPFlag("working_dir", ElementModificationCodeCmd.PersistentFlags().Lookup("working-dir"))

	v.SetConfigFile(homeDir + "/.flagship/credentials/" + utils.WEB_EXPERIMENTATION + "/.cli.yaml")
	v.MergeInConfig()

	err := v.WriteConfig()
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}
	viper.MergeConfigMap(v.AllSettings())
}

func initGlobalCodeDir() {
	config.CheckWorkingDirectory(viper.GetString("working_dir"))
}
