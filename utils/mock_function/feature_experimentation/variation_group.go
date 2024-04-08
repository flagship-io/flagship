package feature_experimentation

import (
	"net/http"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/jarcoal/httpmock"
)

var targetingGroups = []models.TargetingGroup{
	{
		Targetings: []models.InnerTargeting{
			{
				Key:      "isVIP",
				Operator: "CONTAINS",
				Value:    true,
			},
		},
	},
}

var variations = []models.VariationFE{
	{
		Name:       "My variation 1",
		Reference:  true,
		Allocation: 50,
		Modifications: models.Modification{
			Type:  "string",
			Value: "isVIP",
		},
	},
}

var targeting = models.Targeting{
	TargetingGroups: targetingGroups,
}

var TestVariationGroup = models.VariationGroup{
	Id:         "testVariationGroupID",
	Name:       "testVariationGroupName",
	Variations: &variations,
	Targeting:  targeting,
}

var TestVariationGroup1 = models.VariationGroup{
	Id:         "testVariationGroupID1",
	Name:       "testVariationGroupName1",
	Variations: &variations,
	Targeting:  targeting,
}

var TestVariationGroupEdit = models.VariationGroup{
	Id:         "testVariationGroupID",
	Name:       "testVariationGroupName1",
	Variations: &variations,
	Targeting:  targeting,
}

var TestVariationGroupList = []models.VariationGroup{
	TestVariationGroup,
	TestVariationGroup1,
}

func APIVariationGroup() {

	campaignID := "campaignID"

	resp := utils.HTTPListResponseFE[models.VariationGroup]{
		Items:             TestVariationGroupList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/campaigns/"+campaignID+"/variation_groups/"+TestVariationGroup.Id,
		func(req *http.Request) (*http.Response, error) {
			mockResp, _ := httpmock.NewJsonResponse(200, TestVariationGroup)
			return mockResp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/campaigns/"+campaignID+"/variation_groups",
		func(req *http.Request) (*http.Response, error) {
			mockResp, _ := httpmock.NewJsonResponse(200, resp)
			return mockResp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/campaigns/"+campaignID+"/variation_groups",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestVariationGroup)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/campaigns/"+campaignID+"/variation_groups/"+TestVariationGroup.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestVariationGroupEdit)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.GetFeatureExperimentationHost()+"/v1/accounts/"+mockfunction.Auth.AccountID+"/account_environments/"+mockfunction.Auth.AccountEnvironmentID+"/campaigns/"+campaignID+"/variation_groups/"+TestVariationGroup.Id,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil
		},
	)

}
