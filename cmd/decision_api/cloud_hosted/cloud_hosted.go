/*
Copyright © 2022 Flagship Team flagship@abtasty.com

*/

package cloud_hosted

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/flagship-io/flagship/models"
	httprequest "github.com/flagship-io/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

var decisionResponse models.DecisionAPIInfo

// DecisionCmd represents the decision command
var DecisionCloudHostedCmd = &cobra.Command{
	Use:   "call",
	Short: "use of decision api",
	Long:  `use of the decison api in the CLI`,
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := httprequest.HTTPDecisionApi("VisitorId")
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
