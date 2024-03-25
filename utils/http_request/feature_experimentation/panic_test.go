package feature_experimentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var panicRequester = PanicRequester{}

func TestHTTPUpdatePanic(t *testing.T) {

	_, err := panicRequester.HTTPUpdatePanic("active")

	assert.Nil(t, err)
}
