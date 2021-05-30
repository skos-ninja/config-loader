package config

import (
	"strings"

	"github.com/skos-ninja/config-loader/pkg/context"
	"github.com/skos-ninja/config-loader/pkg/parser"

	"github.com/spf13/cobra"
)

const configFlag = "config"

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
	ctx := context.GetContextWithCmd(cmd)

	e := parser.EnvironmentParser{}
	d, _ := e.GetString(ctx, strings.ToUpper(configFlag))
	err := setJSONConfig(d, config)
	if err != nil {
		return err
	}

	f := parser.FlagParser{}
	flag, err := f.GetString(ctx, strings.ToUpper(configFlag))
	if flag != "" {
		err = setJSONConfig(flag, config)
		if err != nil {
			return err
		}
	}

	err = parser.ParseStruct(ctx, config, false)
	if err != nil {
		return err
	}

	return nil
}
