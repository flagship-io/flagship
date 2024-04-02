package project

import (
	"encoding/json"
	"testing"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	mockfunction_fe "github.com/flagship-io/flagship/utils/mock_function/feature_experimentation"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.SetMock(&http_request.ResourceRequester)
	mockfunction_fe.APIProject()

	m.Run()
}

var testProject models.Project
var testProjectList []models.Project

func TestProjectCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(ProjectCmd)
	assert.Contains(t, output, "Manage your projects")
}

func TestProjectHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(ProjectCmd, "--help")
	assert.Contains(t, output, "Manage your projects")
}

func TestProjectGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(ProjectCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(ProjectCmd, "get", "--id=testProjectID")

	err := json.Unmarshal([]byte(successOutput), &testProject)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestProject, testProject)
}

func TestProjectListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(ProjectCmd, "list")

	err := json.Unmarshal([]byte(output), &testProjectList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestProjectList, testProjectList)
}

func TestProjectCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(ProjectCmd, "create")
	assert.Contains(t, failOutput, "Error: required flag(s) \"name\" not set")

	successOutput, _ := utils.ExecuteCommand(ProjectCmd, "create", "--name=testProjectName")

	err := json.Unmarshal([]byte(successOutput), &testProject)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestProject, testProject)
}

func TestProjectEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(ProjectCmd, "edit")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\", \"name\" not set")

	successOutput, _ := utils.ExecuteCommand(ProjectCmd, "edit", "--id=testProjectID", "--name=testProjectName1")

	err := json.Unmarshal([]byte(successOutput), &testProject)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestProjectEdit, testProject)
}

func TestProjectDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(ProjectCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(ProjectCmd, "delete", "--id=testProjectID")
	assert.Equal(t, "Project deleted\n", successOutput)
}

func TestProjectSwitchCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(ProjectCmd, "switch")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\", \"status\" not set")

	failOutput1, _ := utils.ExecuteCommand(ProjectCmd, "switch", "--id=testProjectID", "--status=notKnown")
	assert.Equal(t, "Status can only have 3 values: active or paused or interrupted\n", failOutput1)

	successOutput, _ := utils.ExecuteCommand(ProjectCmd, "switch", "--id=testProjectID", "--status=active")
	assert.Equal(t, "project status set to active\n", successOutput)
}
