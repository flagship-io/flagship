package httprequest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/spf13/viper"
)

func HTTPDecisionApi(host, path, visitorId, visitorContext string) ([]byte, error) {

	var visitorContextMap map[string]any

	err := json.Unmarshal([]byte(visitorContext), &visitorContextMap)
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}

	decisionRequest := models.Visitor{
		VisitorId:      visitorId,
		VisitorContext: visitorContextMap,
	}
	decisionRequestJSON, err := json.Marshal(decisionRequest)
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}

	return HTTPRequestDecisionAPI(http.MethodPost, fmt.Sprintf("%s/v2/%s/%s", host, viper.GetString("account_environment_id"), path), decisionRequestJSON)
}
