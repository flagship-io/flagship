package feature_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
)

var TestVariation = models.VariationFE{
	Id:         "testVariationID",
	Name:       "testVariationName",
	Reference:  true,
	Allocation: 50,
	Modifications: models.Modification{
		Type:  "string",
		Value: "isVIP",
	},
}

var TestVariation1 = models.VariationFE{
	Id:         "testVariationID1",
	Name:       "testVariationName1",
	Reference:  false,
	Allocation: 80,
	Modifications: models.Modification{
		Type:  "string",
		Value: "isVIP1",
	},
}

var TestVariationEdit = models.VariationFE{
	Id:         "testVariationID",
	Name:       "testVariationName1",
	Reference:  false,
	Allocation: 80,
	Modifications: models.Modification{
		Type:  "string",
		Value: "isVIP1",
	},
}

var TestVariationList = []models.VariationFE{
	TestVariation,
	TestVariation1,
}

func APIVariation() {

	campaignID := "campaignID"
	variationGroupID := "variationGroupID"

	resp := utils.HTTPListResponseFE[models.VariationFE]{
		Items:             TestVariationList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+TestVariation.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestVariation)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, resp)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestVariation)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+TestVariation.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestVariationEdit)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/campaigns/"+campaignID+"/variation_groups/"+variationGroupID+"/variations/"+TestVariation.Id,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil
		},
	)

}
