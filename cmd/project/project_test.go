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
	assert.Equal(t, "project created: {\"id\":\"testProjectID\",\"name\":\"testProjectName\"}\n", successOutput)
}

func TestEditCommand(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProject := models.Project{
		ID:   "testProjectID",
		Name: "testProjectName1",
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

	failOutput, _ := utils.ExecuteCommand(ProjectCmd, "edit")
	assert.Equal(t, "Error: required flag(s) \"id\", \"name\" not set\nUsage:\n  project edit [-i <project-id> | --id=<project-id>] [-n <name> | --name=<name>] [flags]\n\nFlags:\n  -h, --help          help for edit\n  -i, --id string     id of the project you want to edit\n  -n, --name string   name you want to set for the project\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(ProjectCmd, "edit", "--id="+testProject.ID, "--name=testProjectName1")
	assert.Equal(t, "project updated: {\"id\":\"testProjectID\",\"name\":\"testProjectName1\"}\n", successOutput)
}

func TestDeleteCommand(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProject := models.Project{
		ID:   "testProjectID",
		Name: "testProjectName",
	}

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+testProject.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(204, testProject)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	failOutput, _ := utils.ExecuteCommand(ProjectCmd, "delete")
	assert.Equal(t, "Error: required flag(s) \"id\" not set\nUsage:\n  project delete [-i <project-id> | --id=<project-id>] [flags]\n\nFlags:\n  -h, --help        help for delete\n  -i, --id string   id of the project you want to delete\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(ProjectCmd, "delete", "--id="+testProject.ID)
	assert.Equal(t, "Project deleted\n", successOutput)
}

func TestToggleCommand(t *testing.T) {
	config.ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testProject := models.Project{
		ID:   "testProjectID",
		Name: "testProjectName",
	}

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/projects/"+testProject.ID+"/toggle",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testProject)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	failOutput, _ := utils.ExecuteCommand(ProjectCmd, "toggle")
	assert.Equal(t, "Error: required flag(s) \"id\", \"status\" not set\nUsage:\n  project toggle [-i <project-id> | --id=<project-id>] [-s <status> | --status=<status>] [flags]\n\nFlags:\n  -h, --help            help for toggle\n  -i, --id string       id of the project you want to toggle\n  -s, --status string   status you want to set to the project. Only 3 values are possible: active, paused and interrupted\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(ProjectCmd, "toggle", "--id="+testProject.ID, "--status=active")
	assert.Equal(t, "project set to active\n", successOutput)
}
