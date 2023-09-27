package httprequest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPGetGoal(t *testing.T) {

	respBody, err := HTTPGetGoal("testGoalID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testGoalID", respBody.Id)
	assert.Equal(t, "testGoalLabel", respBody.Label)
}

func TestHTTPListGoal(t *testing.T) {

	respBody, err := HTTPListGoal()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testGoalID", respBody[0].Id)
	assert.Equal(t, "testGoalLabel", respBody[0].Label)

	assert.Equal(t, "testGoalID1", respBody[1].Id)
	assert.Equal(t, "testGoalLabel1", respBody[1].Label)
}

func TestHTTPCreateGoal(t *testing.T) {

	data := "{\"label\":\"testGoalLabel\", \"type\":\"screenview\", \"operator\":\"contains\", \"value\":\"VIP\"}"

	respBody, err := HTTPCreateGoal(data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testGoalID\",\"label\":\"testGoalLabel\",\"type\":\"screenview\",\"operator\":\"contains\",\"value\":\"VIP\"}"), respBody)
}

func TestHTTPEditGoal(t *testing.T) {

	data := "{\"label\":\"testGoalLabel\", \"type\":\"screenview\", \"operator\":\"contains\", \"value\":\"VIP\"}"

	respBody, err := HTTPEditGoal("testGoalID", data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testGoalID\",\"label\":\"testGoalLabel1\",\"type\":\"screenview\",\"operator\":\"contains\",\"value\":\"VIP1\"}"), respBody)
}

func TestHTTPDeleteGoal(t *testing.T) {

	err := HTTPDeleteGoal("testGoalID")

	assert.Nil(t, err)
}
