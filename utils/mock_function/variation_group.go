package mockfunction

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
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

var variations = []models.Variation{
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

	config.SetViper()

	campaignID := "campaignID"

	resp := utils.HTTPListResponse[models.VariationGroup]{
		Items:             TestVariationGroupList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+TestVariationGroup.Id,
		func(req *http.Request) (*http.Response, error) {
			mockResp, _ := httpmock.NewJsonResponse(200, TestVariationGroup)
			return mockResp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups",
		func(req *http.Request) (*http.Response, error) {
			mockResp, _ := httpmock.NewJsonResponse(200, resp)
			return mockResp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestVariationGroup)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+TestVariationGroup.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestVariationGroupEdit)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+TestVariationGroup.Id,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil
		},
	)

}
