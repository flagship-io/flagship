package project

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

func TestHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(ProjectCmd, "--help")
	assert.Equal(t, output, "Manage your projects\n\nUsage:\n  project [create|edit|get|list|delete|toggle] [flags]\n  project [command]\n\nAvailable Commands:\n  completion  Generate the autocompletion script for the specified shell\n  create      Create a project\n  delete      Delete a project\n  edit        Edit a project\n  get         Get a project\n  help        Help about any command\n  list        List all projects\n  toggle      Toggle a project\n\nFlags:\n  -h, --help   help for project\n\nUse \"project [command] --help\" for more information about a command.\n")

}

func TestGetCommand(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProject := models.Project{
		ID:   "testProjectID",
		Name: "testProjectName",
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+testProject.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testProject)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	failOutput, _ := utils.ExecuteCommand(ProjectCmd, "get")
	assert.Equal(t, "Error: required flag(s) \"id\" not set\nUsage:\n  project get [-i <project-id> | --id=<project-id>] [flags]\n\nFlags:\n  -h, --help        help for get\n  -i, --id string   id of the project you want to display\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(ProjectCmd, "get", "--id=testProjectID")
	assert.Equal(t, "{\"id\":\"testProjectID\",\"name\":\"testProjectName\"}\n", successOutput)
}

func TestListCommand(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProjectList := []models.Project{
		{
			ID:   "testProjectID",
			Name: "testProjectName",
		},
		{
			ID:   "testProjectID1",
			Name: "testProjectName1",
		},
	}

	resp := utils.HTTPListResponse[models.Project]{
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

	output, _ := utils.ExecuteCommand(ProjectCmd, "list")
	assert.Equal(t, "[{\"id\":\"testProjectID\",\"name\":\"testProjectName\"},{\"id\":\"testProjectID1\",\"name\":\"testProjectName1\"}]\n", output)
}

func TestCreateCommand(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProject := models.Project{
		ID:   "testProjectID",
		Name: "testProjectName",
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

	failOutput, _ := utils.ExecuteCommand(ProjectCmd, "create")
	assert.Equal(t, "Error: required flag(s) \"name\" not set\nUsage:\n  project create [-n <name> | --name=<name>] [flags]\n\nFlags:\n  -h, --help          help for create\n  -n, --name string   name of the project you want to create\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(ProjectCmd, "create", "--name=testProjectName")
	assert.Equal(t, "", successOutput)
}
