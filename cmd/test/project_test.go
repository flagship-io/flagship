package test

import (
	"testing"

	"github.com/flagship-io/flagship/cmd/project"
	"github.com/flagship-io/flagship/utils"
	"github.com/stretchr/testify/assert"
)

func TestProjectCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(project.ProjectCmd)
	assert.Equal(t, "Manage your projects\n\nUsage:\n  project [create|edit|get|list|delete|toggle] [flags]\n  project [command]\n\nAvailable Commands:\n  completion  Generate the autocompletion script for the specified shell\n  create      Create a project\n  delete      Delete a project\n  edit        Edit a project\n  get         Get a project\n  help        Help about any command\n  list        List all projects\n  toggle      Toggle a project\n\nFlags:\n  -h, --help   help for project\n\nUse \"project [command] --help\" for more information about a command.\n", output)
}

func TestProjectHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(project.ProjectCmd, "--help")
	assert.Equal(t, "Manage your projects\n\nUsage:\n  project [create|edit|get|list|delete|toggle] [flags]\n  project [command]\n\nAvailable Commands:\n  completion  Generate the autocompletion script for the specified shell\n  create      Create a project\n  delete      Delete a project\n  edit        Edit a project\n  get         Get a project\n  help        Help about any command\n  list        List all projects\n  toggle      Toggle a project\n\nFlags:\n  -h, --help   help for project\n\nUse \"project [command] --help\" for more information about a command.\n", output)
}

func TestProjectGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "get")
	assert.Equal(t, "Error: required flag(s) \"id\" not set\nUsage:\n  project get [-i <project-id> | --id=<project-id>] [flags]\n\nFlags:\n  -h, --help        help for get\n  -i, --id string   id of the project you want to display\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "get", "--id=testProjectID")
	assert.Equal(t, "{\"id\":\"testProjectID\",\"name\":\"testProjectName\"}\n", successOutput)
}

func TestProjectListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(project.ProjectCmd, "list")
	assert.Equal(t, "[{\"id\":\"testProjectID\",\"name\":\"testProjectName\"},{\"id\":\"testProjectID1\",\"name\":\"testProjectName1\"}]\n", output)
}

func TestProjectCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "create")
	assert.Equal(t, "Error: required flag(s) \"name\" not set\nUsage:\n  project create [-n <name> | --name=<name>] [flags]\n\nFlags:\n  -h, --help          help for create\n  -n, --name string   name of the project you want to create\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "create", "--name=testProjectName")
	assert.Equal(t, "project created: {\"id\":\"testProjectID\",\"name\":\"testProjectName\"}\n", successOutput)
}

func TestProjectEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "edit")
	assert.Equal(t, "Error: required flag(s) \"id\", \"name\" not set\nUsage:\n  project edit [-i <project-id> | --id=<project-id>] [-n <name> | --name=<name>] [flags]\n\nFlags:\n  -h, --help          help for edit\n  -i, --id string     id of the project you want to edit\n  -n, --name string   name you want to set for the project\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "edit", "--id=testProjectID", "--name=testProjectName1")
	assert.Equal(t, "project updated: {\"id\":\"testProjectID\",\"name\":\"testProjectName1\"}\n", successOutput)
}

func TestProjectDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "delete")
	assert.Equal(t, "Error: required flag(s) \"id\" not set\nUsage:\n  project delete [-i <project-id> | --id=<project-id>] [flags]\n\nFlags:\n  -h, --help        help for delete\n  -i, --id string   id of the project you want to delete\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "delete", "--id=testProjectID")
	assert.Equal(t, "Project deleted\n", successOutput)
}

func TestProjectToggleCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "toggle")
	assert.Equal(t, "Error: required flag(s) \"id\", \"status\" not set\nUsage:\n  project toggle [-i <project-id> | --id=<project-id>] [-s <status> | --status=<status>] [flags]\n\nFlags:\n  -h, --help            help for toggle\n  -i, --id string       id of the project you want to toggle\n  -s, --status string   status you want to set to the project. Only 3 values are possible: active, paused and interrupted\n\n", failOutput)

	failOutput1, _ := utils.ExecuteCommand(project.ProjectCmd, "toggle", "--id=testProjectID", "--status=notKnown")
	assert.Equal(t, "Status can only have 3 values: active or paused or interrupted\n", failOutput1)

	successOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "toggle", "--id=testProjectID", "--status=active")
	assert.Equal(t, "project set to active\n", successOutput)
}
