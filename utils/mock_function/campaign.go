package mockfunction

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
)

var scheduler = models.Scheduler{
	StartDate: "2022-02-01 10:00:00",
	StopDate:  "2022-02-02 08:00:00",
	TimeZone:  "Europe/Paris",
}

var targetingGroupsTest = []models.TargetingGroup{
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

var targetingTest = models.Targeting{
	TargetingGroups: targetingGroupsTest,
}

var variationTest = []models.Variation{
	{
		Name:       "My variation 1",
		Reference:  true,
		Allocation: 50,
		Modifications: models.Modification{
			Type:  "string",
			Value: map[string]interface{}{"color": "blue"},
		},
	},
	{
		Name:       "My variation 2",
		Reference:  false,
		Allocation: 50,
		Modifications: models.Modification{
			Type:  "string",
			Value: map[string]interface{}{"color": "red"},
		},
	},
}

var variationGroupsTest = []models.VariationGroup{
	{
		Name:       "variationGroupName",
		Variations: variationTest,
		Targeting:  targetingTest,
	},
}

var TestCampaign = models.Campaign{
	Id:              "testCampaignID",
	Name:            "testCampaignName",
	ProjectId:       "testProjectID",
	Description:     "testCampaignDescription",
	Type:            "toggle",
	VariationGroups: variationGroupsTest,
	Scheduler:       scheduler,
}

var TestCampaign1 = models.Campaign{
	Id:              "testCampaignID1",
	Name:            "testCampaignName1",
	ProjectId:       "testProjectID1",
	Description:     "testCampaignDescription1",
	Type:            "toggle",
	VariationGroups: variationGroupsTest,
	Scheduler:       scheduler,
}

var TestCampaignEdit = models.Campaign{
	Id:              "testCampaignID",
	Name:            "testCampaignName1",
	ProjectId:       "testProjectID1",
	Description:     "testCampaignDescription1",
	Type:            "toggle",
	VariationGroups: variationGroupsTest,
	Scheduler:       scheduler,
}

var TestCampaignlist = []models.Campaign{
	TestCampaign,
	TestCampaign1,
}

func APICampaign() {
	config.SetViper()

	resp := utils.HTTPListResponse[models.Campaign]{
		Items:             TestCampaignlist,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+TestCampaign.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestCampaign)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, resp)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestCampaign)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+TestCampaign.Id,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestCampaignEdit)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+TestCampaign.Id,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+TestCampaign.Id+"/toggle",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, ""), nil
		},
	)
}
