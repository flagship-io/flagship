/*
Copyright © 2022 Flagship Team flagship@abtasty.com
*/
package variation_global_code

import (
	"github.com/spf13/cobra"
)

var CampaignID int
var VariationID int

// VariationGlobalCodeCmd represents the variation global code command
var VariationGlobalCodeCmd = &cobra.Command{
	Use:     "variation-global-code [get]",
	Short:   "Get variable global code",
	Aliases: []string{"vgc"},
	Long:    `Get variable global code`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
