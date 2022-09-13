package version

import (
	"fmt"
	"testing"

	"github.com/flagship-io/flagship/cmd/authorization"
	"github.com/flagship-io/flagship/utils"
	"github.com/magiconair/properties/assert"
)

func TestGetResult(t *testing.T) {
	output, _ := utils.ExecuteCommand(authorization.VersionCmd, "", "")
	fmt.Println(output)
	assert.Equal(t, "Flagship CLI version : v0.2.0", output)
}
