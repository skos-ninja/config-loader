package main

import (
	"fmt"

	"github.com/skos-ninja/config-loader"
	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:  "example",
		RunE: runE,
	}
	cfg = &exampleConfig{
		Env:  "default",
		Flag: "default",
	}
)

func init() {
	config.Init(cmd)
	cmd.Flags().String("CONFIG_FLAG", cfg.Flag, "Config flag")
}

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

type exampleConfig struct {
	Env  string `env:"CONFIG_ENV"`
	Flag string `flag:"CONFIG_FLAG"`
}

func runE(cmd *cobra.Command, args []string) error {
	config.Load(cmd, cfg)

	fmt.Printf("env:  %s\n", cfg.Env)
	fmt.Printf("flag: %s\n", cfg.Flag)
	return nil
}
