/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package account_global_code

import (
	"log"
	"os"

	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var WorkingDir string

// AccountGlobalCodeCmd represents the account global code command
var AccountGlobalCodeCmd = &cobra.Command{
	Use:     "account-global-code [get]",
	Short:   "Get account global code",
	Aliases: []string{"agc"},
	Long:    `Get account global code`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initGlobalCodeDir)
	AccountGlobalCodeCmd.PersistentFlags().StringVarP(&WorkingDir, "working-dir", "", utils.DefaultGlobalCodeWorkingDir(), "Directory where the file will be generated and pushed from")

}

func initConfig() {
	v := viper.New()

	homeDir, _ := os.UserHomeDir()

	v.BindPFlag("working_dir", AccountGlobalCodeCmd.PersistentFlags().Lookup("working-dir"))

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
