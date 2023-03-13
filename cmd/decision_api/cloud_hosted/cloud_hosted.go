/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package cloud_hosted

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var (
	visitorId        string
	visitorContext   string
	decisionResponse models.DecisionAPIInfo
	path             string
	host             string
	extended         bool
)

// DecisionCmd represents the decision command
var DecisionCloudHostedCmd = &cobra.Command{
	Use:   "call",
	Short: "use of decision api",
	Long:  `use of the decison api in the CLI`,
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := httprequest.HTTPDecisionApi(host, path, visitorId, visitorContext)
		if err != nil {
			log.Fatalf("error occured: %s", err)
		}

		err = json.Unmarshal(resp, &decisionResponse)
		if err != nil {
			log.Fatalf("error occured: %s", err)
		}

		if extended {
			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()

			tbl := table.New("Project Id", "Project Name", "Campaign Id", "Campaign Name", "Type", "Status")
			tbl.WithHeaderFormatter(headerFmt).WithPadding(2)

			for _, campaign := range decisionResponse.CampaignInfos {
				campaignModel, _ := httprequest.HTTPGetCampaign(campaign.Id)
				projectModel, _ := httprequest.HTTPGetProject(campaignModel.ProjectID)

				tbl.AddRow(projectModel.ID, projectModel.Name, campaignModel.ID, campaignModel.Name, campaignModel.Type, campaignModel.Status)

			}
			tbl.AddRow(fmt.Sprintf("\nThis visitor is present in %d campaign(s).", len(decisionResponse.CampaignInfos)))
			tbl.Print()
			return
		}

		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", string(resp))

	},
}

func init() {
	DecisionCloudHostedCmd.Flags().StringVarP(&visitorId, "visitor-id", "", "", "visitorId")

	if err := DecisionCloudHostedCmd.MarkFlagRequired("visitor-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	DecisionCloudHostedCmd.Flags().StringVarP(&visitorContext, "visitor-context", "", "", "visitorId")

	if err := DecisionCloudHostedCmd.MarkFlagRequired("visitor-context"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	DecisionCloudHostedCmd.Flags().StringVarP(&path, "path", "", "campaigns", "path")
	DecisionCloudHostedCmd.Flags().StringVarP(&host, "host", "", utils.GetDecisionAPIHost(), "host")
	DecisionCloudHostedCmd.Flags().BoolVarP(&extended, "extended", "", false, "format")

}
