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

var CampaignID = "campaignID"

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

func TestHTTPGetVariationGroup(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testVariationGroup := models.VariationGroup{
		ID:         "testVariationGroupID",
		Name:       "testVariationGroupName",
		Variations: variations,
		Targeting:  targeting,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+CampaignID+"/variation_groups/"+testVariationGroup.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testVariationGroup)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPGetVariationGroup(CampaignID, "testVariationGroupID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testVariationGroupID", respBody.ID)
	assert.Equal(t, "testVariationGroupName", respBody.Name)
}

func TestHTTPListVariationGroup(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testVariationGroupList := []models.VariationGroup{
		{
			ID:         "testVariationGroupID",
			Name:       "testVariationGroupName",
			Variations: variations,
			Targeting:  targeting,
		},
		{
			ID:         "testVariationGroupID1",
			Name:       "testVariationGroupName1",
			Variations: variations,
			Targeting:  targeting,
		},
	}

	resp := utils.HTTPListResponse[models.VariationGroup]{
		Items:             testVariationGroupList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+CampaignID+"/variation_groups",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPListVariationGroup(CampaignID)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testVariationGroupID", respBody[0].ID)
	assert.Equal(t, "testVariationGroupName", respBody[0].Name)

	assert.Equal(t, "testVariationGroupID1", respBody[1].ID)
	assert.Equal(t, "testVariationGroupName1", respBody[1].Name)
}

func TestHTTPCreateVariationGroup(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testVariationGroup := models.VariationGroup{
		ID:         "testVariationGroupID",
		Name:       "testVariationGroupName",
		Variations: variations,
		Targeting:  targeting,
	}

	data := "{\"name\":\"testVariationGroupName\",\"variations\":[{\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"operator\":\"CONTAINS\",\"key\":\"isVIP\",\"value\":true}]}]}}"

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+CampaignID+"/variation_groups",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testVariationGroup)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPCreateVariationGroup(CampaignID, data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testVariationGroupID\",\"name\":\"testVariationGroupName\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}"), respBody)
}

func TestHTTPEditVariationGroup(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testVariationGroup := models.VariationGroup{
		ID:         "testVariationGroupID1",
		Name:       "testVariationGroupName1",
		Variations: variations,
		Targeting:  targeting,
	}

	data := "{\"name\":\"testVariationGroupName1\",\"variations\":[{\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"operator\":\"CONTAINS\",\"key\":\"isVIP\",\"value\":true}]}]}}"

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+CampaignID+"/variation_groups/"+testVariationGroup.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testVariationGroup)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPEditVariationGroup(CampaignID, testVariationGroup.ID, data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testVariationGroupID1\",\"name\":\"testVariationGroupName1\",\"variations\":[{\"id\":\"\",\"name\":\"My variation 1\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}],\"targeting\":{\"targeting_groups\":[{\"targetings\":[{\"key\":\"isVIP\",\"operator\":\"CONTAINS\",\"value\":true}]}]}}"), respBody)
}

func TestHTTPDeleteVariationGroup(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testVariationGroup := models.VariationGroup{
		ID:         "testVariationGroupID",
		Name:       "testVariationGroupName",
		Variations: variations,
		Targeting:  targeting,
	}

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+CampaignID+"/variation_groups/"+testVariationGroup.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)

	err := HTTPDeleteVariationGroup(CampaignID, "testVariationGroupID")

	assert.Nil(t, err)
}
