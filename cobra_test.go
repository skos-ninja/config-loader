package config

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestLoad(t *testing.T) {
	type args struct {
		cmd    *cobra.Command
		config interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Load(tt.args.cmd, tt.args.config)
		})
	}
}
