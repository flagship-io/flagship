package mockfunction

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
)

var TestVariation = models.Variation{
	ID:         "testVariationID",
	Name:       "testVariationName",
	Reference:  true,
	Allocation: 50,
	Modifications: models.Modification{
		Type:  "string",
		Value: "isVIP",
	},
}

var TestVariation1 = models.Variation{
	ID:         "testVariationID1",
	Name:       "testVariationName1",
	Reference:  false,
	Allocation: 80,
	Modifications: models.Modification{
		Type:  "string",
		Value: "isVIP1",
	},
}

var TestVariationEdit = models.Variation{
	ID:         "testVariationID",
	Name:       "testVariationName1",
	Reference:  false,
	Allocation: 80,
	Modifications: models.Modification{
		Type:  "string",
		Value: "isVIP1",
	},
}

var TestVariationList = []models.Variation{
	TestVariation,
	TestVariation1,
}

func APIVariation() {

	config.SetViper()

	campaignID := "campaignID"
	variationGroupID := "variationGroupID"

	resp := utils.HTTPListResponse[models.Variation]{
		Items:             TestVariationList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+TestVariation.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, TestVariation)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, TestVariation)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+TestVariation.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, TestVariationEdit)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+TestVariation.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)

}
