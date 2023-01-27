package mockfunction

import (
	"net/http"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
)

var TestProject = models.Project{
	ID:   "testProjectID",
	Name: "testProjectName",
}

var TestProject1 = models.Project{
	ID:   "testProjectID1",
	Name: "testProjectName1",
}

var TestProjectEdit = models.Project{
	ID:   "testProjectID",
	Name: "testProjectName1",
}

var TestProjectList = []models.Project{
	TestProject,
	TestProject1,
}

func APIProject() {

	config.SetViper()

	resp := utils.HTTPListResponse[models.Project]{
		Items:             TestProjectList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+TestProject.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestProject)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("GET", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/projects",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, resp)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("POST", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/projects",
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestProject)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PATCH", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+TestProject.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, _ := httpmock.NewJsonResponse(200, TestProjectEdit)
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+TestProject.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)

	httpmock.RegisterResponder("PATCH", utils.GetHost()+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+TestProject.ID+"/toggle",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, ""), nil
		},
	)
}
