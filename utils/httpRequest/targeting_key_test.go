package httprequest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPGetTargetingKey(t *testing.T) {

	respBody, err := HTTPGetTargetingKey("testTargetingKeyID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testTargetingKeyID", respBody.Id)
	assert.Equal(t, "testTargetingKeyName", respBody.Name)
}

func TestHTTPListTargetingKey(t *testing.T) {

	respBody, err := HTTPListTargetingKey()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testTargetingKeyID", respBody[0].Id)
	assert.Equal(t, "testTargetingKeyName", respBody[0].Name)

	assert.Equal(t, "testTargetingKeyID1", respBody[1].Id)
	assert.Equal(t, "testTargetingKeyName1", respBody[1].Name)
}

func TestHTTPCreateTargetingKey(t *testing.T) {

	data := "{\"name\":\"testTargetingKeyName\", \"type\":\"string\", \"description\":\"testTargetingKeyDescription\"}"

	respBody, err := HTTPCreateTargetingKey(data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testTargetingKeyID\",\"name\":\"testTargetingKeyName\",\"type\":\"string\",\"description\":\"testTargetingKeyDescription\"}"), respBody)
}

func TestHTTPEditTargetingKey(t *testing.T) {

	respBody, err := HTTPEditTargetingKey("testTargetingKeyID", "testTargetingKeyName")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testTargetingKeyID\",\"name\":\"testTargetingKeyName1\",\"type\":\"string\",\"description\":\"testTargetingKeyDescription1\"}"), respBody)
}

func TestHTTPDeleteTargetingKey(t *testing.T) {

	err := HTTPDeleteTargetingKey("testTargetingKeyID")

	assert.Nil(t, err)
}
