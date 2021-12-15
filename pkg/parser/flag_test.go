package parser

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	c "github.com/skos-ninja/config-loader/pkg/context"

	"github.com/spf13/cobra"
)

type flag struct {
	name       string
	value      string
	kind       reflect.Kind
	forceEmpty bool
}

var f = FlagParser{}

func TestGetFlagNotCobra(t *testing.T) {
	ctx := context.Background()
	_, err := f.GetString(ctx, "test")
	if err == nil || err != ErrNotUsingCobraCtx {
		t.Errorf("expected() = %v, want %v", err, ErrNotUsingCobraCtx)
	}
}

func TestGetFlagNotFound(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	ctx := c.GetContextWithCmd(cmd)
	_, err := f.GetString(ctx, "test")
	expectErr := ErrFlagNotFound{flag: "test"}
	if err == nil || err != expectErr {
		t.Errorf("expected() = %v, want %v", err, expectErr)
	}
}

func TestGetFlagString(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	ctx := c.GetContextWithCmd(cmd)
	tests := []struct {
		name    string
		args    flag
		want    string
		wantErr bool
	}{
		{
			name: "Valid String: test",
			args: flag{
				name:  "valid-string",
				value: "test",
				kind:  reflect.String,
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "Valid Empty String",
			args: flag{
				name:       "valid-empty-string",
				value:      "",
				forceEmpty: true,
				kind:       reflect.String,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Invalid Empty String",
			args: flag{
				name:  "invalid-empty-string",
				value: "",
				kind:  reflect.String,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setFlag(cmd, tt.args)
			got, err := f.GetString(ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFlagInt(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	ctx := c.GetContextWithCmd(cmd)
	tests := []struct {
		name    string
		args    flag
		want    int64
		wantErr bool
	}{
		{
			name: "Valid Int: 1",
			args: flag{
				name:  "valid-int",
				value: "1",
				kind:  reflect.Int64,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Valid Int: -1",
			args: flag{
				name:  "valid-negative-int",
				value: "-1",
				kind:  reflect.Int64,
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Invalid Int: missing",
			args: flag{
				name:  "invalid-missing",
				value: "",
				kind:  reflect.Int64,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setFlag(cmd, tt.args)
			got, err := f.GetInt(ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFlagFloat(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	ctx := c.GetContextWithCmd(cmd)
	tests := []struct {
		name    string
		e       EnvironmentParser
		args    flag
		want    float64
		wantErr bool
	}{
		{
			name: "Valid Float: 1",
			args: flag{
				name:  "valid-float",
				value: "1",
				kind:  reflect.Float64,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Valid Float: -1",
			args: flag{
				name:  "valid-negative-float",
				value: "-1",
				kind:  reflect.Float64,
			},
			want:    -1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setFlag(cmd, tt.args)
			got, err := f.GetFloat(ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFlagBoolean(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	ctx := c.GetContextWithCmd(cmd)
	tests := []struct {
		name    string
		args    flag
		want    bool
		wantErr bool
	}{
		{
			name: "Valid Boolean: true",
			args: flag{
				name:  "valid-true",
				value: "true",
				kind:  reflect.Bool,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Valid Capital Boolean: true",
			args: flag{
				name:  "valid-capital-true",
				value: "TRUE",
				kind:  reflect.Bool,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Valid Boolean: false",
			args: flag{
				name:  "valid-false",
				value: "false",
				kind:  reflect.Bool,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Valid Capital Boolean: false",
			args: flag{
				name:  "valid-capital-false",
				value: "FALSE",
				kind:  reflect.Bool,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Invalid Boolean: missing",
			args: flag{
				name:  "invalid-missing",
				value: "",
				kind:  reflect.Bool,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setFlag(cmd, tt.args)
			got, err := f.GetBoolean(ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoolean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBoolean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setFlag(cmd *cobra.Command, flag flag) {
	switch flag.kind {
	case reflect.String:
		cmd.Flags().String(flag.name, "", "")
	case reflect.Int64:
		cmd.Flags().Int64(flag.name, 0, "")
	case reflect.Float64:
		cmd.Flags().Float64(flag.name, 0, "")
	case reflect.Bool:
		cmd.Flags().Bool(flag.name, false, "")
	default:
		panic(fmt.Sprintf("%v unsupported type: %v", flag.name, flag.kind))
	}

	if flag.value != "" || flag.forceEmpty {
		err := cmd.Flags().Set(flag.name, flag.value)
		if err != nil {
			panic(err)
		}
		cmd.Flags().Lookup(flag.name).Changed = true
	}
}
