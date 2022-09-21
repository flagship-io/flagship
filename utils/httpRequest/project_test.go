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

func TestHTTPGetProject(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProjectID := "1"

	testProject := models.Project{
		ID:   "1",
		Name: "newTestingProject",
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+testProjectID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testProject)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPGetProject("1")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, testProjectID, respBody.ID)
	assert.Equal(t, "newTestingProject", respBody.Name)
}

func TestHTTPListProject(t *testing.T) {

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

func TestHTTPCreateProject(t *testing.T) {
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

func TestHTTPEditProject(t *testing.T) {
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

func TestHTTPDeleteProject(t *testing.T) {
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

func TestHTTPToggleProject(t *testing.T) {
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
