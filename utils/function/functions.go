package function

import (
	"github.com/flagship-io/flagship/cmd/authorization"
	"github.com/spf13/cobra"
)

func RegenerateToken(cmd *cobra.Command, args []string) {
	authorization.AuthenticateCmd.Run(cmd, args)
}
