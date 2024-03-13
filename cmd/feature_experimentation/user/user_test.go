package user

import (
	"encoding/json"
	"testing"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function/feature_experimentation"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.APIUser()
	m.Run()
}

var testUserList []models.User

func TestUserCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(UserCmd)
	assert.Contains(t, output, "Manage your users")
}

func TestUserHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(UserCmd, "--help")
	assert.Contains(t, output, "Manage your users")
}

func TestUserListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(UserCmd, "list")

	err := json.Unmarshal([]byte(output), &testUserList)

	assert.Nil(t, err)

	assert.Equal(t, mockfunction.TestUserList, testUserList)
}

func TestUserCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(UserCmd, "create")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\" not set")

	successOutput, _ := utils.ExecuteCommand(UserCmd, "create", "--data-raw='[{\"email\":\"example@abtasty.com\",\"role\":\"ADMIN\"},{\"email\":\"example1@abtasty.com\",\"role\":\"VIEWER\"}]'")

	assert.Equal(t, "users created\n", successOutput)
}

func TestUserEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(UserCmd, "edit")
	assert.Contains(t, failOutput, "Error: required flag(s) \"data-raw\" not set")

	successOutput, _ := utils.ExecuteCommand(UserCmd, "edit", "--data-raw='[{\"email\":\"example@abtasty.com\",\"role\":\"ADMIN\"},{\"email\":\"example1@abtasty.com\",\"role\":\"VIEWER\"}]'")

	assert.Equal(t, "users created\n", successOutput)
}

func TestUserDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(UserCmd, "delete")
	assert.Contains(t, failOutput, "Error: required flag(s) \"email\" not set")

	successOutput, _ := utils.ExecuteCommand(UserCmd, "delete", "--email=example@abtasty.com")
	assert.Equal(t, "Email deleted\n", successOutput)
}
