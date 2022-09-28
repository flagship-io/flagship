package mockfunction

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
)

func APITargetingKey() {

	config.SetViper()

	testTargetingKey := models.TargetingKey{
		ID:          "testTargetingKeyID",
		Name:        "testTargetingKeyName",
		Type:        "string",
		Description: "testTargetingKeyDescription",
	}

	testTargetingKey1 := models.TargetingKey{
		ID:          "testTargetingKeyID1",
		Name:        "testTargetingKeyName1",
		Type:        "string",
		Description: "testTargetingKeyDescription1",
	}

	testTargetingKeyEdit := models.TargetingKey{
		ID:          "testTargetingKeyID",
		Name:        "testTargetingKeyName1",
		Type:        "string",
		Description: "testTargetingKeyDescription1",
	}

	testTargetingKeyList := []models.TargetingKey{
		testTargetingKey,
		testTargetingKey1,
	}

	resp := utils.HTTPListResponse[models.TargetingKey]{
		Items:             testTargetingKeyList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys/"+testTargetingKey.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testTargetingKey)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testTargetingKey)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys/"+testTargetingKey.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testTargetingKeyEdit)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/targeting_keys/"+testTargetingKey.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)
}
