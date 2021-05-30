package context

import (
	"context"

	"github.com/spf13/cobra"
)

type contextKey string

var cmdKey contextKey = "cmd"

func GetContextWithCmd(cmd *cobra.Command) context.Context {
	ctx := cmd.Context()
	if ctx == nil {
		ctx = context.TODO()
	}
	return context.WithValue(ctx, cmdKey, cmd)
}

func GetCmdFromContext(ctx context.Context) *cobra.Command {
	v := ctx.Value(cmdKey)
	if v != nil {
		return v.(*cobra.Command)
	}

	return nil
}
