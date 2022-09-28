package test

import (
	"encoding/json"
	"testing"

	"github.com/flagship-io/flagship/cmd/project"
	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	"github.com/stretchr/testify/assert"
)

var testProject models.Project
var testProjectList []models.Project

func TestProjectCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(project.ProjectCmd)
	assert.Contains(t, output, "Manage your projects")
}

func TestProjectHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(project.ProjectCmd, "--help")
	assert.Contains(t, output, "Manage your projects")
}

func TestProjectGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "get", "--id=testProjectID")

	err := json.Unmarshal([]byte(successOutput), &testProject)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestProject, testProject)
}

func TestProjectListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(project.ProjectCmd, "list")

	err := json.Unmarshal([]byte(output), &testProjectList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestProjectList, testProjectList)
}

func TestProjectCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "create")
	assert.Contains(t, failOutput, "Error: required flag(s) \"name\" not set")

	successOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "create", "--name=testProjectName")

	err := json.Unmarshal([]byte(successOutput), &testProject)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestProject, testProject)
}

func TestProjectEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "edit")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\", \"name\" not set")

	successOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "edit", "--id=testProjectID", "--name=testProjectName1")

	err := json.Unmarshal([]byte(successOutput), &testProject)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestProjectEdit, testProject)
}

func TestProjectDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "delete", "--id=testProjectID")
	assert.Equal(t, "Project deleted\n", successOutput)
}

func TestProjectToggleCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "toggle")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\", \"status\" not set")

	failOutput1, _ := utils.ExecuteCommand(project.ProjectCmd, "toggle", "--id=testProjectID", "--status=notKnown")
	assert.Equal(t, "Status can only have 3 values: active or paused or interrupted\n", failOutput1)

	successOutput, _ := utils.ExecuteCommand(project.ProjectCmd, "toggle", "--id=testProjectID", "--status=active")
	assert.Equal(t, "project set to active\n", successOutput)
}
