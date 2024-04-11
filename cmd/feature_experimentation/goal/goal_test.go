package goal

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
	mockfunction_fe.APIGoal()

	m.Run()
}

var testGoal models.Goal
var testGoalList []models.Goal

func TestGoalCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(GoalCmd)
	assert.Contains(t, output, "Manage your goals")
}

func TestGoalHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(GoalCmd, "--help")
	assert.Contains(t, output, "Manage your goals")
}

func TestGoalGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(GoalCmd, "get")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(GoalCmd, "get", "--id=testGoalID")

	err := json.Unmarshal([]byte(successOutput), &testGoal)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestGoal, testGoal)
}

func TestGoalListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(GoalCmd, "list")

	err := json.Unmarshal([]byte(output), &testGoalList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestGoalList, testGoalList)
}

func TestGoalCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(GoalCmd, "create")

	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\" not set")

	successOutput, _ := utils.ExecuteCommand(GoalCmd, "create", "--data-raw='{\"id\":\"testGoalID\",\"label\":\"testGoalLabel\",\"type\":\"screenview\",\"operator\":\"contains\",\"value\":\"VIP\"}'")

	err := json.Unmarshal([]byte(successOutput), &testGoal)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestGoal, testGoal)
}

func TestGoalEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(GoalCmd, "edit")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\", \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(GoalCmd, "edit", "--id=testGoalID", "--data-raw='{\"id\":\"testGoalID\",\"label\":\"testGoalLabel1\",\"type\":\"screenview\",\"operator\":\"contains\",\"value\":\"VIP1\"}'")

	err := json.Unmarshal([]byte(successOutput), &testGoal)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction_fe.TestGoalEdit, testGoal)
}

func TestGoalDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(GoalCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(GoalCmd, "delete", "--id=testGoalID")
	assert.Equal(t, "Goal deleted\n", successOutput)
}
