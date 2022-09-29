package httprequest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPUpdatePanic(t *testing.T) {

	_, err := HTTPUpdatePanic("active")

	assert.Nil(t, err)
}
