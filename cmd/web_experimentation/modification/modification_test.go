package modification

import (
	"encoding/json"
	"strconv"
	"testing"

	models "github.com/flagship-io/flagship/models/web_experimentation"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/http_request"
	mockfunction "github.com/flagship-io/flagship/utils/mock_function"
	mockfunction_we "github.com/flagship-io/flagship/utils/mock_function/web_experimentation"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockfunction.SetMock(&http_request.ResourceRequester)
	mockfunction_we.APIModification()

	m.Run()
}

var testModification []models.Modification
var testModificationList []models.Modification

func TestModificationCommand(t *testing.T) {
	failOutput, _ := utils.ExecuteCommand(ModificationCmd)
	assert.Contains(t, failOutput, "Error: required flag(s) \"campaign-id\" not set")

	successOutput, _ := utils.ExecuteCommand(ModificationCmd, "--campaign-id="+strconv.Itoa(100000))
	assert.Contains(t, successOutput, "Manage your modifications")
}

func TestModificationHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(ModificationCmd, "--help")
	assert.Contains(t, output, "Manage your modifications")
}

func TestModificationGetCommand(t *testing.T) {

	failOutput, err := utils.ExecuteCommand(ModificationCmd, "get", "--campaign-id="+strconv.Itoa(100000))
	assert.Contains(t, failOutput, "Error: required flag(s) \"id\" not set")

	successOutput, _ := utils.ExecuteCommand(ModificationCmd, "get", "--id="+strconv.Itoa(120003), "--campaign-id="+strconv.Itoa(100000))

	err = json.Unmarshal([]byte(successOutput), &testModification)

	assert.Nil(t, err)

	assert.Equal(t, []models.Modification{mockfunction_we.TestElementModification}, testModification)
}

func TestModificationListCommand(t *testing.T) {

	output, err := utils.ExecuteCommand(ModificationCmd, "list", "--campaign-id="+strconv.Itoa(100000))
	err = json.Unmarshal([]byte(output), &testModificationList)

	assert.Nil(t, err)
	assert.Equal(t, []models.Modification{mockfunction_we.TestModificationsJS, mockfunction_we.TestModificationsCSS}, testModificationList)
}
