package httprequest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPGetFlag(t *testing.T) {

	respBody, err := HTTPGetFlag("testFlagID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testFlagID", respBody.Id)
	assert.Equal(t, "testFlagName", respBody.Name)
}

func TestHTTPListFlag(t *testing.T) {

	respBody, err := HTTPListFlag()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testFlagID", respBody[0].Id)
	assert.Equal(t, "testFlagName", respBody[0].Name)

	assert.Equal(t, "testFlagID1", respBody[1].Id)
	assert.Equal(t, "testFlagName1", respBody[1].Name)
}

func TestHTTPFlagUsage(t *testing.T) {

	respBody, err := HTTPFlagUsage()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testFlagUsageID", respBody[0].Id)
	assert.Equal(t, "isVIP", respBody[0].FlagKey)
	assert.Equal(t, "flagship", respBody[0].Repository)
	assert.Equal(t, "https://github.com/flagship-io/flagship", respBody[0].FilePath)
	assert.Equal(t, "main", respBody[0].Branch)

}

func TestHTTPCreateFlag(t *testing.T) {

	data := "{\"name\":\"testFlagName\", \"type\":\"string\", \"description\":\"testFlagDescription\", \"source\":\"manual\"}"

	respBody, err := HTTPCreateFlag(data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testFlagID\",\"name\":\"testFlagName\",\"type\":\"string\",\"description\":\"testFlagDescription\",\"source\":\"manual\"}"), respBody)
}

func TestHTTPEditFlag(t *testing.T) {

	data := "{\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"manual\"}"

	respBody, err := HTTPEditFlag("testFlagID", data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testFlagID\",\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"manual\"}"), respBody)
}

func TestHTTPDeleteFlag(t *testing.T) {

	err := HTTPDeleteFlag("testFlagID")

	assert.Nil(t, err)
}
