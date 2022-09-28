package mockfunction

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
)

type FlagKey struct {
	Color string `json:"color"`
}

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
			Type: "string",
			Value: FlagKey{
				Color: "blue",
			},
		},
	},
	{
		Name:       "My variation 2",
		Reference:  false,
		Allocation: 50,
		Modifications: models.Modification{
			Type: "string",
			Value: FlagKey{
				Color: "red",
			},
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

func APICampaign() {
	config.SetViper()

	testCampaign := models.Campaign{
		ID:              "testCampaignID",
		Name:            "testCampaignName",
		ProjectID:       "testProjectID",
		Description:     "testCampaignDescription",
		Type:            "toggle",
		VariationGroups: variationGroupsTest,
		Scheduler:       scheduler,
	}

	testCampaignEdit := models.Campaign{
		ID:              "testCampaignID",
		Name:            "testCampaignName1",
		ProjectID:       "testProjectID1",
		Description:     "testCampaignDescription1",
		Type:            "toggle",
		VariationGroups: variationGroupsTest,
		Scheduler:       scheduler,
	}

	testCampaign1 := models.Campaign{
		ID:              "testCampaignID1",
		Name:            "testCampaignName1",
		ProjectID:       "testProjectID1",
		Description:     "testCampaignDescription1",
		Type:            "toggle",
		VariationGroups: variationGroupsTest,
		Scheduler:       scheduler,
	}

	testCampaignlist := []models.Campaign{
		testCampaign,
		testCampaign1,
	}

	resp := utils.HTTPListResponse[models.Campaign]{
		Items:             testCampaignlist,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+testCampaign.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testCampaign)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testCampaign)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+testCampaign.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testCampaignEdit)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+testCampaign.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+testCampaign.ID+"/toggle",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, ""), nil
		},
	)
}
