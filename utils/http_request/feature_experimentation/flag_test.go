package feature_experimentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var flagRequester = FlagRequester{}

func TestHTTPGetFlag(t *testing.T) {

	respBody, err := flagRequester.HTTPGetFlag("testFlagID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testFlagID", respBody.Id)
	assert.Equal(t, "testFlagName", respBody.Name)
}

func TestHTTPListFlag(t *testing.T) {

	respBody, err := flagRequester.HTTPListFlag()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testFlagID", respBody[0].Id)
	assert.Equal(t, "testFlagName", respBody[0].Name)

	assert.Equal(t, "testFlagID1", respBody[1].Id)
	assert.Equal(t, "testFlagName1", respBody[1].Name)
}

func TestHTTPCreateFlag(t *testing.T) {

	data := "{\"name\":\"testFlagName\", \"type\":\"string\", \"description\":\"testFlagDescription\", \"source\":\"cli\"}"

	respBody, err := flagRequester.HTTPCreateFlag(data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testFlagID\",\"name\":\"testFlagName\",\"type\":\"string\",\"description\":\"testFlagDescription\",\"source\":\"cli\"}"), respBody)
}

func TestHTTPEditFlag(t *testing.T) {

	data := "{\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"cli\"}"

	respBody, err := flagRequester.HTTPEditFlag("testFlagID", data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testFlagID\",\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"cli\"}"), respBody)
}

func TestHTTPDeleteFlag(t *testing.T) {

	err := flagRequester.HTTPDeleteFlag("testFlagID")

	assert.Nil(t, err)
}
