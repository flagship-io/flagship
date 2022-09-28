package test

import (
	"testing"

	"github.com/flagship-io/flagship/cmd/flag"
	"github.com/flagship-io/flagship/utils"
	"github.com/stretchr/testify/assert"
)

func TestFlagCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(flag.FlagCmd)
	assert.Equal(t, "Manage your flags in your account\n\nUsage:\n  flag [create|edit|get|list|delete|usage] [flags]\n  flag [command]\n\nAvailable Commands:\n  completion  Generate the autocompletion script for the specified shell\n  create      Create a flag\n  delete      Delete a flag\n  edit        Edit a flag\n  get         Get a flag\n  help        Help about any command\n  list        List all flags\n  usage       Manage flag usage statistics inside your codebase\n\nFlags:\n  -h, --help   help for flag\n\nUse \"flag [command] --help\" for more information about a command.\n", output)
}

func TestFlagHelpCommand(t *testing.T) {
	output, _ := utils.ExecuteCommand(flag.FlagCmd, "--help")
	assert.Equal(t, "Manage your flags in your account\n\nUsage:\n  flag [create|edit|get|list|delete|usage] [flags]\n  flag [command]\n\nAvailable Commands:\n  completion  Generate the autocompletion script for the specified shell\n  create      Create a flag\n  delete      Delete a flag\n  edit        Edit a flag\n  get         Get a flag\n  help        Help about any command\n  list        List all flags\n  usage       Manage flag usage statistics inside your codebase\n\nFlags:\n  -h, --help   help for flag\n\nUse \"flag [command] --help\" for more information about a command.\n", output)
}

func TestFlagGetCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "get")
	assert.Equal(t, "Error: required flag(s) \"id\" not set\nUsage:\n  flag get [-i <flag-id> | --id=<flag-id>] [flags]\n\nFlags:\n  -h, --help        help for get\n  -i, --id string   id of the flag you want to display\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "get", "--id=testFlagID")
	assert.Equal(t, "{\"id\":\"testFlagID\",\"name\":\"testFlagName\",\"type\":\"string\",\"description\":\"testFlagDescription\",\"source\":\"manual\"}\n", successOutput)
}

func TestFlagListCommand(t *testing.T) {

	output, _ := utils.ExecuteCommand(flag.FlagCmd, "list")
	assert.Equal(t, "[{\"id\":\"testFlagID\",\"name\":\"testFlagName\",\"type\":\"string\",\"description\":\"testFlagDescription\",\"source\":\"manual\"},{\"id\":\"testFlagID1\",\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"manual\"}]\n", output)
}

func TestFlagCreateCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "create")
	assert.Equal(t, "Error: required flag(s) \"data-raw\" not set\nUsage:\n  flag create [-d <data-raw> | --data-raw <data-raw>] [flags]\n\nFlags:\n  -d, --data-raw string   raw data contains all the info to create your flag, check the doc for details\n  -h, --help              help for create\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "create", "--data-raw='{\"name\":\"testFlagName\",\"type\":\"string\",\"description\":\"testFlagDescription\",\"source\":\"manual\"}'")
	assert.Equal(t, "flag created: {\"id\":\"testFlagID\",\"name\":\"testFlagName\",\"type\":\"string\",\"description\":\"testFlagDescription\",\"source\":\"manual\"}\n", successOutput)
}

func TestFlagEditCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "edit")
	assert.Equal(t, "Error: required flag(s) \"data-raw\", \"id\" not set\nUsage:\n  flag edit [-i <flag-id> | --id=<flag-id>] [-d <data-raw> | --data-raw <data-raw>] [flags]\n\nFlags:\n  -d, --data-raw string   raw data contains all the info to edit your flag, check the doc for details\n  -h, --help              help for edit\n  -i, --id string         id of the flag you want to edit\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "edit", "--id=testFlagID", "--data-raw={\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"manual\"}\n")
	assert.Equal(t, "flag updated: {\"id\":\"testFlagID\",\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"manual\"}\n", successOutput)
}

func TestFlagDeleteCommand(t *testing.T) {

	failOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "delete")
	assert.Equal(t, "Error: required flag(s) \"id\" not set\nUsage:\n  flag delete [-i <flag-id> | --id=<flag-id>] [flags]\n\nFlags:\n  -h, --help        help for delete\n  -i, --id string   id of the flag you want to delete\n\n", failOutput)

	successOutput, _ := utils.ExecuteCommand(flag.FlagCmd, "delete", "--id=testFlagID")
	assert.Equal(t, "Flag deleted\n", successOutput)
}
