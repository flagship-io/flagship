package project

import (
	"testing"

	"github.com/flagship-io/flagship/cmd/authorization"
	"github.com/flagship-io/flagship/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetResult(t *testing.T) {
	output, _ := utils.ExecuteCommand(authorization.VersionCmd, "", "")
	assert.Equal(t, output, "Flagship CLI version : v0.2.1")
}
