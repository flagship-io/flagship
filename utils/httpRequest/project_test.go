package httprequest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPGetProject(t *testing.T) {

	respBody, err := HTTPGetProject("testProjectID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testProjectID", respBody.ID)
	assert.Equal(t, "testProjectName", respBody.Name)
}

func TestHTTPListProject(t *testing.T) {

	respBody, err := HTTPListProject()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testProjectID", respBody[0].ID)
	assert.Equal(t, "testProjectName", respBody[0].Name)

	assert.Equal(t, "testProjectID1", respBody[1].ID)
	assert.Equal(t, "testProjectName1", respBody[1].Name)
}

func TestHTTPCreateProject(t *testing.T) {

	respBody, err := HTTPCreateProject("testProjectName")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testProjectID\",\"name\":\"testProjectName\"}"), respBody)
}

func TestHTTPEditProject(t *testing.T) {

	respBody, err := HTTPEditProject("testProjectID", "testProjectName1")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testProjectID\",\"name\":\"testProjectName1\"}"), respBody)
}

func TestHTTPDeleteProject(t *testing.T) {

	err := HTTPDeleteProject("testProjectID")

	assert.Nil(t, err)
}

func TestHTTPToggleProject(t *testing.T) {

	err := HTTPToggleProject("testProjectID", "active")

	assert.Nil(t, err)
}
