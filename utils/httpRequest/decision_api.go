package httprequest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/spf13/viper"
)

func HTTPDecisionApi(name string) ([]byte, error) {
	decisionRequest := models.Visitor{
		VisitorId:      "jflsdkjlsf",
		VisitorContext: map[string]any{"device": "firefox"},
	}
	decisionRequestJSON, err := json.Marshal(decisionRequest)
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}
	return HTTPRequestDecisionAPI(http.MethodPost, fmt.Sprintf("%s/v2/%s/campaigns", utils.GetDecisionAPIHost(), viper.GetString("env_id")), decisionRequestJSON)
}
