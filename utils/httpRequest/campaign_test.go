package httprequest

import (
	"net/http"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type FlagKey struct {
	Color string
}

var variationTest = []models.Variation{
	{
		Name:       "My variation 1",
		Reference:  true,
		Allocation: 50,
		Modifications: models.Modification{
			Type: "string",
			Value: FlagKey{
				Color: "red",
			},
		},
	},
	{
		Name:       "My variation 2",
		Reference:  true,
		Allocation: 30,
		Modifications: models.Modification{
			Type: "string",
			Value: FlagKey{
				Color: "blue",
			},
		},
	},
}

var variationGroupsTest = []models.VariationGroup{
	{
		Name:       "variationGroupName",
		Variations: variationTest,
	},
}

func TestHTTPGetCampaign(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testCampaignID := "campaignID1"

	testCampaign := models.Campaign{
		ID:              testCampaignID,
		Name:            "newTestingCampaign",
		ProjectID:       "projectIDTest",
		Description:     "descriptionTest",
		Type:            "toggle",
		VariationGroups: variationGroupsTest,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/campaigns/"+testCampaignID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testCampaign)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPGetCampaign(testCampaignID)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "campaignID1", respBody.ID)
	assert.Equal(t, "newTestingCampaign", respBody.Name)
	assert.Equal(t, "projectIDTest", respBody.ProjectID)
	assert.Equal(t, "descriptionTest", respBody.Description)
	assert.Equal(t, "toggle", respBody.Type)
}

func TestHTTPListCampaign(t *testing.T) {

	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProjectList := []models.Project{
		{
			ID:   "1",
			Name: "newTestingProject1",
		},
		{
			ID:   "2",
			Name: "newTestingProject2",
		},
	}

	resp := HTTPListResponse[models.Project]{
		Items:             testProjectList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPListProject()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "1", respBody[0].ID)
	assert.Equal(t, "newTestingProject1", respBody[0].Name)

	assert.Equal(t, "2", respBody[1].ID)
	assert.Equal(t, "newTestingProject2", respBody[1].Name)
}

func TestHTTPCreateCampaign(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProject := models.Project{
		ID:   "1",
		Name: "newTestingProject",
	}

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testProject)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPCreateProject(testProject.Name)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"1\",\"name\":\"newTestingProject\"}"), respBody)
}

func TestHTTPEditCampaign(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProject := models.Project{
		ID:   "1",
		Name: "newTestingProject",
	}

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+testProject.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testProject)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPEditProject(testProject.ID, testProject.Name)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"1\",\"name\":\"newTestingProject\"}"), respBody)
}

func TestHTTPDeleteCampaign(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProject := models.Project{
		ID:   "1",
		Name: "newTestingProject",
	}

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+testProject.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)

	err := HTTPDeleteProject("1")

	assert.Nil(t, err)
}

func TestHTTPToggleCampaign(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProject := models.Project{
		ID:   "1",
		Name: "newTestingProject",
	}

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+testProject.ID+"/toggle",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, ""), nil
		},
	)

	err := HTTPToggleProject(testProject.ID, "active")

	assert.Nil(t, err)
}
