/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package decision_api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/flagship-io/flagship/cmd/decision_api/self_hosted"
	"github.com/flagship-io/flagship/models"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

var (
	visitorId      string
	visitorContext []string
)

var decisionResponse models.DecisionAPIInfo

// DecisionCmd represents the decision command
var DecisionCmd = &cobra.Command{
	Use:   "decision-api",
	Short: "use of decision api",
	Long:  `use of the decison api in the CLI`,
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := httprequest.HTTPDecisionApi(visitorId)
		if err != nil {
			log.Fatalf("error occured: %s", err)
		}

		err = json.Unmarshal(resp, &decisionResponse)
		if err != nil {
			log.Fatalf("error occured: %s", err)
		}

		fmt.Println(string(resp))

	},
}

func init() {

	DecisionCmd.Flags().StringVarP(&visitorId, "visitor-id", "", "", "visitorId")

	if err := DecisionCmd.MarkFlagRequired("visitor-id"); err != nil {
		log.Fatalf("error occurred: %v", err)
	}

	DecisionCmd.AddCommand(self_hosted.DecisionSelfHostedCmd)
}
