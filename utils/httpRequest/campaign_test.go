package httprequest

import (
	"net/http"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
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

func TestHTTPGetCampaign(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testCampaign := models.Campaign{
		ID:              "testCampaignID",
		Name:            "testCampaignName",
		ProjectID:       "testProjectID",
		Description:     "testCampaignDescription",
		Type:            "toggle",
		VariationGroups: variationGroupsTest,
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

	respBody, err := HTTPGetCampaign("testCampaignID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testCampaignID", respBody.ID)
	assert.Equal(t, "testCampaignName", respBody.Name)
	assert.Equal(t, "testProjectID", respBody.ProjectID)
	assert.Equal(t, "testCampaignDescription", respBody.Description)
	assert.Equal(t, "toggle", respBody.Type)
}

func TestHTTPListCampaign(t *testing.T) {

	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testCampaignlist := []models.Campaign{
		{
			ID:              "testCampaignID",
			Name:            "testCampaignName",
			ProjectID:       "testProjectID",
			Description:     "testCampaignDescription",
			Type:            "toggle",
			VariationGroups: variationGroupsTest,
		},
		{
			ID:              "testCampaignID1",
			Name:            "testCampaignName1",
			ProjectID:       "testProjectID1",
			Description:     "testCampaignDescription1",
			Type:            "toggle",
			VariationGroups: variationGroupsTest,
		},
	}

	resp := utils.HTTPListResponse[models.Campaign]{
		Items:             testCampaignlist,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPListCampaign()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testCampaignID", respBody[0].ID)
	assert.Equal(t, "testCampaignName", respBody[0].Name)
	assert.Equal(t, "testProjectID", respBody[0].ProjectID)
	assert.Equal(t, "testCampaignDescription", respBody[0].Description)
	assert.Equal(t, "toggle", respBody[0].Type)

	assert.Equal(t, "testCampaignID1", respBody[1].ID)
	assert.Equal(t, "testCampaignName1", respBody[1].Name)
	assert.Equal(t, "testProjectID1", respBody[1].ProjectID)
	assert.Equal(t, "testCampaignDescription1", respBody[1].Description)
	assert.Equal(t, "toggle", respBody[1].Type)
}

func TestHTTPCreateCampaign(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testCampaign := models.Campaign{
		ID:              "testCampaignID",
		Name:            "testCampaignName",
		ProjectID:       "testProjectID",
		Description:     "testCampaignDescription",
		Type:            "toggle",
		VariationGroups: variationGroupsTest,
		Scheduler:       scheduler,
	}

	dataCampaign := "{\"project_id\":\"testProjectID\",\"name\":\"testCampaignName\",\"description\":\"testCampaignDescription\",\"type\":\"toggle\",\"variation_groups\":[{\"name\":\"variationGroupName\",\"variations\":[{\"name\":\"My variation 1\",\"allocation\":50,\"reference\":true,\"modifications\":{\"value\":{\"color\":\"blue\"}}},{\"name\":\"My variation 2\",\"allocation\":50,\"reference\":false,\"modifications\":{\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"operator\":\"CONTAINS\",\"key\":\"isVIP\",\"value\":\"true\"}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}"

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testCampaign)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPCreateCampaign(dataCampaign)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testCampaignID\",\"project_id\":\"testProjectID\",\"name\":\"testCampaignName\",\"description\":\"testCampaignDescription\",\"type\":\"toggle\",\"status\":\"\",\"variation_groups\":[{\"id\":\"\",\"name\":\"variationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"blue\"}}},{\"id\":\"\",\"name\":\"My variation 2\",\"reference\":false,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}"), respBody)
}

func TestHTTPEditCampaign(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testCampaign := models.Campaign{
		ID:              "testCampaignID",
		Name:            "testCampaignName1",
		ProjectID:       "testProjectID1",
		Description:     "testCampaignDescription1",
		Type:            "toggle",
		VariationGroups: variationGroupsTest,
		Scheduler:       scheduler,
	}

	dataCampaign := "{\"project_id\":\"testProjectID1\",\"name\":\"testCampaignName1\",\"description\":\"testCampaignDescription1\",\"type\":\"toggle\",\"variation_groups\":[{\"name\":\"variationGroupName\",\"variations\":[{\"name\":\"My variation 1\",\"allocation\":50,\"reference\":true,\"modifications\":{\"value\":{\"color\":\"blue\"}}},{\"name\":\"My variation 2\",\"allocation\":50,\"reference\":false,\"modifications\":{\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"operator\":\"CONTAINS\",\"key\":\"isVIP\",\"value\":\"true\"}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}"

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+testCampaign.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testCampaign)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPEditCampaign("testCampaignID", dataCampaign)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testCampaignID\",\"project_id\":\"testProjectID1\",\"name\":\"testCampaignName1\",\"description\":\"testCampaignDescription1\",\"type\":\"toggle\",\"status\":\"\",\"variation_groups\":[{\"id\":\"\",\"name\":\"variationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"blue\"}}},{\"id\":\"\",\"name\":\"My variation 2\",\"reference\":false,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":{\"color\":\"red\"}}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}],\"scheduler\":{\"start_date\":\"2022-02-01 10:00:00\",\"stop_date\":\"2022-02-02 08:00:00\",\"timezone\":\"Europe/Paris\"}}"), respBody)
}

func TestHTTPDeleteCampaign(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testCampaign := models.Campaign{
		ID:              "testCampaignID",
		Name:            "testCampaignName",
		ProjectID:       "testProjectID",
		Description:     "testCampaignDescription",
		Type:            "toggle",
		VariationGroups: variationGroupsTest,
		Scheduler:       scheduler,
	}

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+testCampaign.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)

	err := HTTPDeleteCampaign("testCampaignID")

	assert.Nil(t, err)
}

func TestHTTPToggleCampaign(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testCampaign := models.Campaign{
		ID:              "testCampaignID",
		Name:            "testCampaignName",
		ProjectID:       "testProjectID",
		Description:     "testCampaignDescription",
		Type:            "toggle",
		VariationGroups: variationGroupsTest,
		Scheduler:       scheduler,
	}

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+testCampaign.ID+"/toggle",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, ""), nil
		},
	)

	err := HTTPToggleCampaign(testCampaign.ID, "active")

	assert.Nil(t, err)
}
