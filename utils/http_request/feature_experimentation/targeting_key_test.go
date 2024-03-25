package feature_experimentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var targetingKeyRequester = TargetingKeyRequester{}

func TestHTTPGetTargetingKey(t *testing.T) {

	respBody, err := targetingKeyRequester.HTTPGetTargetingKey("testTargetingKeyID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testTargetingKeyID", respBody.Id)
	assert.Equal(t, "testTargetingKeyName", respBody.Name)
}

func TestHTTPListTargetingKey(t *testing.T) {

	respBody, err := targetingKeyRequester.HTTPListTargetingKey()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testTargetingKeyID", respBody[0].Id)
	assert.Equal(t, "testTargetingKeyName", respBody[0].Name)

	assert.Equal(t, "testTargetingKeyID1", respBody[1].Id)
	assert.Equal(t, "testTargetingKeyName1", respBody[1].Name)
}

func TestHTTPCreateTargetingKey(t *testing.T) {

	data := "{\"name\":\"testTargetingKeyName\", \"type\":\"string\", \"description\":\"testTargetingKeyDescription\"}"

	respBody, err := targetingKeyRequester.HTTPCreateTargetingKey(data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testTargetingKeyID\",\"name\":\"testTargetingKeyName\",\"type\":\"string\",\"description\":\"testTargetingKeyDescription\"}"), respBody)
}

func TestHTTPEditTargetingKey(t *testing.T) {

	respBody, err := targetingKeyRequester.HTTPEditTargetingKey("testTargetingKeyID", "testTargetingKeyName")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testTargetingKeyID\",\"name\":\"testTargetingKeyName1\",\"type\":\"string\",\"description\":\"testTargetingKeyDescription1\"}"), respBody)
}

func TestHTTPDeleteTargetingKey(t *testing.T) {

	err := targetingKeyRequester.HTTPDeleteTargetingKey("testTargetingKeyID")

	assert.Nil(t, err)
}
