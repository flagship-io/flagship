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

var variationGroupID = "variationGroupID"

func TestHTTPGetVariation(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testVariation := models.Variation{
		ID:         "testVariationID",
		Name:       "testVariationName",
		Reference:  true,
		Allocation: 50,
		Modifications: models.Modification{
			Type:  "string",
			Value: "isVIP",
		},
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+CampaignID+"/variation_groups/"+variationGroupID+"/variations/"+testVariation.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testVariation)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPGetVariation(CampaignID, variationGroupID, "testVariationID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testVariationID", respBody.ID)
	assert.Equal(t, "testVariationName", respBody.Name)
}

func TestHTTPListVariation(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testVariationList := []models.Variation{
		{
			ID:         "testVariationID",
			Name:       "testVariationName",
			Reference:  true,
			Allocation: 50,
			Modifications: models.Modification{
				Type:  "string",
				Value: "isVIP",
			},
		},
		{
			ID:         "testVariationID1",
			Name:       "testVariationName1",
			Reference:  false,
			Allocation: 80,
			Modifications: models.Modification{
				Type:  "string",
				Value: "isVIP1",
			},
		},
	}

	resp := utils.HTTPListResponse[models.Variation]{
		Items:             testVariationList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+CampaignID+"/variation_groups/"+variationGroupID+"/variations",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPListVariation(CampaignID, variationGroupID)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testVariationID", respBody[0].ID)
	assert.Equal(t, "testVariationName", respBody[0].Name)
	assert.Equal(t, true, respBody[0].Reference)
	assert.Equal(t, 50, respBody[0].Allocation)

	assert.Equal(t, "testVariationID1", respBody[1].ID)
	assert.Equal(t, "testVariationName1", respBody[1].Name)
	assert.Equal(t, false, respBody[1].Reference)
	assert.Equal(t, 80, respBody[1].Allocation)
}

func TestHTTPCreateVariation(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testVariation := models.Variation{
		ID:         "testVariationID",
		Name:       "testVariationName",
		Reference:  true,
		Allocation: 50,
		Modifications: models.Modification{
			Type:  "string",
			Value: "isVIP",
		},
	}

	data := "{\"name\":\"testVariationName\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}"

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+CampaignID+"/variation_groups/"+variationGroupID+"/variations",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testVariation)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPCreateVariation(CampaignID, variationGroupID, data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testVariationID\",\"name\":\"testVariationName\",\"reference\":true,\"allocation\":50,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP\"}}"), respBody)
}

func TestHTTPEditVariation(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testVariation := models.Variation{
		ID:         "testVariationID1",
		Name:       "testVariationName1",
		Reference:  false,
		Allocation: 80,
		Modifications: models.Modification{
			Type:  "string",
			Value: "isVIP1",
		},
	}

	data := "{\"name\":\"testVariationName1\",\"reference\":false,\"allocation\":80,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP1\"}}"

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+CampaignID+"/variation_groups/"+variationGroupID+"/variations/"+testVariation.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testVariation)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPEditVariation(CampaignID, variationGroupID, testVariation.ID, data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testVariationID1\",\"name\":\"testVariationName1\",\"reference\":false,\"allocation\":80,\"modifications\":{\"type\":\"string\",\"value\":\"isVIP1\"}}"), respBody)
}

func TestHTTPDeleteVariation(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testVariation := models.Variation{
		ID:         "testVariationID",
		Name:       "testVariationName",
		Reference:  true,
		Allocation: 50,
		Modifications: models.Modification{
			Type:  "string",
			Value: "isVIP",
		},
	}

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+CampaignID+"/variation_groups/"+variationGroupID+"/variations/"+testVariation.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)

	err := HTTPDeleteVariation(CampaignID, variationGroupID, "testVariationID")

	assert.Nil(t, err)
}
