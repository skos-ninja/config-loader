package parser

import (
	"context"
	"testing"

	c "github.com/skos-ninja/config-loader/pkg/context"

	"github.com/spf13/cobra"
)

type flag struct {
	name       string
	value      string
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
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Invalid Empty String",
			args: flag{
				name:  "invalid-empty-string",
				value: "",
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
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Valid Int: -1",
			args: flag{
				name:  "valid-negative-int",
				value: "-1",
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Invalid Int: missing",
			args: flag{
				name:  "invalid-missing",
				value: "",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Invalid Int: not-a-int",
			args: flag{
				name:  "invalid-int",
				value: "not-a-int",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Invalid Int: not-a-whole-int",
			args: flag{
				name:  "invalid-whole-int",
				value: "100.10",
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
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Valid Float: -1",
			args: flag{
				name:  "valid-negative-float",
				value: "-1",
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Invalid Float: missing",
			args: flag{
				name:  "invalid-missing",
				value: "",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Invalid Float: not-a-float",
			args: flag{
				name:  "invalid-float",
				value: "not-a-float",
			},
			want:    0,
			wantErr: true,
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
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Valid Capital Boolean: true",
			args: flag{
				name:  "valid-capital-true",
				value: "TRUE",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Valid Boolean: false",
			args: flag{
				name:  "valid-false",
				value: "false",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Valid Capital Boolean: false",
			args: flag{
				name:  "valid-capital-false",
				value: "FALSE",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Invalid Boolean: missing",
			args: flag{
				name:  "invalid-missing",
				value: "",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Invalid Boolean: not-a-bool",
			args: flag{
				name:  "invalid-bool",
				value: "not-a-bool",
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
	cmd.Flags().String(flag.name, "", "")

	if flag.value != "" || flag.forceEmpty {
		f := cmd.Flag(flag.name)
		f.Value.Set(flag.value)
		f.Changed = true
	}
}
