package config

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
)

const configFlag = "config"

type contextKey string

var cmdKey contextKey = "cmd"

func Init(cmd *cobra.Command) {
	cmd.PersistentFlags().String(configFlag, "", "Set the config data")
}

func Load(cmd *cobra.Command, config interface{}) error {
	return load(cmd, config)
}

func MustLoad(cmd *cobra.Command, config interface{}) {
	err := load(cmd, config)
	if err != nil {
		panic(err)
	}
}

func load(cmd *cobra.Command, config interface{}) error {
	ctx := getContextWithCmd(cmd)

	e := EnvironmentParser{}
	d, _ := e.GetString(ctx, strings.ToUpper(configFlag))
	err := setJSONConfig(d, config)
	if err != nil {
		return err
	}

	flag := cmd.Flag(configFlag)
	if flag != nil {
		val := flag.Value.String()
		err = setJSONConfig(val, config)
		if err != nil {
			return err
		}
	}

	err = ParseStruct(ctx, config, false)
	if err != nil {
		return err
	}

	return nil
}

func getContextWithCmd(cmd *cobra.Command) context.Context {
	return context.WithValue(cmd.Context(), cmdKey, cmd)
}

func getCmdFromContext(ctx context.Context) *cobra.Command {
	v := ctx.Value(cmdKey)
	if v != nil {
		return v.(*cobra.Command)
	}

	return nil
}
