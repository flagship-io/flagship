package feature_experimentation

import (
	"encoding/json"
	"log"
	"testing"

	models "github.com/flagship-io/flagship/models/feature_experimentation"
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/stretchr/testify/assert"
)

var projectRequester = ProjectRequester{&common.ResourceRequest{AccountID: "account_id", AccountEnvironmentID: "account_environment_id"}}

func TestHTTPGetProject(t *testing.T) {

	respBody, err := projectRequester.HTTPGetProject("testProjectID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testProjectID", respBody.Id)
	assert.Equal(t, "testProjectName", respBody.Name)
}

func TestHTTPListProject(t *testing.T) {

	respBody, err := projectRequester.HTTPListProject()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testProjectID", respBody[0].Id)
	assert.Equal(t, "testProjectName", respBody[0].Name)

	assert.Equal(t, "testProjectID1", respBody[1].Id)
	assert.Equal(t, "testProjectName1", respBody[1].Name)
}

func TestHTTPCreateProject(t *testing.T) {
	projectRequest := models.Project{
		Name: "testProjectName",
	}
	projectRequestJSON, err := json.Marshal(projectRequest)
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}
	respBody, err := projectRequester.HTTPCreateProject(projectRequestJSON)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testProjectID\",\"name\":\"testProjectName\"}"), respBody)
}

func TestHTTPEditProject(t *testing.T) {
	projectRequest := models.Project{
		Name: "testProjectName1",
	}
	projectRequestJSON, err := json.Marshal(projectRequest)
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}
	respBody, err := projectRequester.HTTPEditProject("testProjectID", projectRequestJSON)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testProjectID\",\"name\":\"testProjectName1\"}"), respBody)
}

func TestHTTPDeleteProject(t *testing.T) {

	err := projectRequester.HTTPDeleteProject("testProjectID")

	assert.Nil(t, err)
}

func TestHTTPSwitchProject(t *testing.T) {

	err := projectRequester.HTTPSwitchProject("testProjectID", "active")

	assert.Nil(t, err)
}
