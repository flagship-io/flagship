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

func APIVariationGroup() {

	config.SetViper()

	campaignID := "campaignID"

	testVariationGroup := models.VariationGroup{
		ID:         "testVariationGroupID",
		Name:       "testVariationGroupName",
		Variations: variations,
		Targeting:  targeting,
	}

	testVariationGroup1 := models.VariationGroup{
		ID:         "testVariationGroupID1",
		Name:       "testVariationGroupName1",
		Variations: variations,
		Targeting:  targeting,
	}

	testVariationGroupEdit := models.VariationGroup{
		ID:         "testVariationGroupID",
		Name:       "testVariationGroupName1",
		Variations: variations,
		Targeting:  targeting,
	}

	testVariationGroupList := []models.VariationGroup{
		testVariationGroup,
		testVariationGroup1,
	}

	resp := utils.HTTPListResponse[models.VariationGroup]{
		Items:             testVariationGroupList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+testVariationGroup.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testVariationGroup)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testVariationGroup)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+testVariationGroup.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testVariationGroupEdit)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+campaignID+"/variation_groups/"+testVariationGroup.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)

}
