package parser

import (
	"context"
	"os"
	"testing"
)

type env struct {
	name       string
	value      string
	forceEmpty bool
}

var e = EnvironmentParser{}

func TestGetEnvString(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		args    env
		want    string
		wantErr bool
	}{
		{
			name: "Valid String: test",
			args: env{
				name:  "valid-string",
				value: "test",
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "Valid Empty String",
			args: env{
				name:       "valid-empty-string",
				value:      "",
				forceEmpty: true,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Invalid Empty String",
			args: env{
				name:  "invalid-empty-string",
				value: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(t, tt.args)
			got, err := e.GetString(ctx, tt.args.name)
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

func TestGetEnvInt(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		args    env
		want    int64
		wantErr bool
	}{
		{
			name: "Valid Int: 1",
			args: env{
				name:  "valid-int",
				value: "1",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Valid Int: -1",
			args: env{
				name:  "valid-negative-int",
				value: "-1",
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Invalid Int: missing",
			args: env{
				name:  "invalid-missing",
				value: "",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Invalid Int: not-a-int",
			args: env{
				name:  "invalid-int",
				value: "not-a-int",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Invalid Int: not-a-whole-int",
			args: env{
				name:  "invalid-whole-int",
				value: "100.10",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(t, tt.args)
			got, err := e.GetInt(ctx, tt.args.name)
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

func TestGetFloat(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		e       EnvironmentParser
		args    env
		want    float64
		wantErr bool
	}{
		{
			name: "Valid Float: 1",
			args: env{
				name:  "valid-float",
				value: "1",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Valid Float: -1",
			args: env{
				name:  "valid-negative-float",
				value: "-1",
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "Invalid Float: missing",
			args: env{
				name:  "invalid-missing",
				value: "",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Invalid Float: not-a-float",
			args: env{
				name:  "invalid-float",
				value: "not-a-float",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(t, tt.args)
			got, err := e.GetFloat(ctx, tt.args.name)
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

func TestGetEnvBoolean(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		args    env
		want    bool
		wantErr bool
	}{
		{
			name: "Valid Boolean: true",
			args: env{
				name:  "valid-true",
				value: "true",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Valid Capital Boolean: true",
			args: env{
				name:  "valid-capital-true",
				value: "TRUE",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Valid Boolean: false",
			args: env{
				name:  "valid-false",
				value: "false",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Valid Capital Boolean: false",
			args: env{
				name:  "valid-capital-false",
				value: "FALSE",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Invalid Boolean: missing",
			args: env{
				name:  "invalid-missing",
				value: "",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Invalid Boolean: not-a-bool",
			args: env{
				name:  "invalid-bool",
				value: "not-a-bool",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(t, tt.args)
			got, err := e.GetBoolean(ctx, tt.args.name)
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

func setEnv(t *testing.T, env env) {
	// Ignore setting if the value is empty
	if env.value == "" && !env.forceEmpty {
		return
	}
	// Set the environment variable before fetching
	err := os.Setenv(env.name, env.value)
	if err != nil {
		t.Error(err)
		return
	}
}
