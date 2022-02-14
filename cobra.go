package config

import (
	"os"
	"strings"

	"github.com/skos-ninja/config-loader/pkg/context"
	"github.com/skos-ninja/config-loader/pkg/parser"

	"github.com/spf13/cobra"
)

var configFlag = ""

func Init(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&configFlag, "config", configFlag, "Set the json config data (Input types: file path, environment var name, flag name)")
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

	// Try to read the config json from a file
	s, _ := os.ReadFile(configFlag)
	err := setJSONConfig(string(s), config)
	if err != nil {
		return err
	}

	// Try to read the config json from an env
	d, _ := parser.EnvironmentParser{}.GetString(ctx, strings.ToUpper(configFlag))
	err = setJSONConfig(d, config)
	if err != nil {
		return err
	}

	// Try to read the config json from a flag
	flag, _ := parser.FlagParser{}.GetString(ctx, configFlag)
	err = setJSONConfig(flag, config)
	if err != nil {
		return err
	}

	// Perform parsing on each field
	err = parser.ParseStruct(ctx, config, false)
	if err != nil {
		return err
	}

	return nil
}
