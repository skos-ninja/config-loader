package parser

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const envTagName = "env"

type ErrEnvVariableNotFound struct {
	variable string
}

func (e ErrEnvVariableNotFound) Error() string {
	return fmt.Sprintf("Environment variable not found: %s", e.variable)
}

type EnvironmentParser struct {
}

// GetString returns an environment variable as a string
func (e EnvironmentParser) GetString(ctx context.Context, name string) (string, error) {
	if name != strings.ToUpper(name) {
		fmt.Printf("WARN: %v is not in upper case. It is recommended all env variables are upper case\n", name)
	}

	if value, ok := os.LookupEnv(name); ok {
		return value, nil
	}

	return "", ErrEnvVariableNotFound{name}
}

// GetInt returns an environment variable as an integer
func (e EnvironmentParser) GetInt(ctx context.Context, name string) (int64, error) {
	d, err := e.GetString(ctx, name)
	if err != nil {
		return 0, err
	}

	parseint, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return 0, err
	}

	return parseint, nil
}

// GetFloat returns an environment variable as a float
func (e EnvironmentParser) GetFloat(ctx context.Context, name string) (float64, error) {
	d, err := e.GetString(ctx, name)
	if err != nil {
		return 0, err
	}

	parsefloat, err := strconv.ParseFloat(d, 64)
	if err != nil {
		return 0, err
	}

	return parsefloat, nil
}

// GetBoolean returns an environment variable as a boolean
func (e EnvironmentParser) GetBoolean(ctx context.Context, name string) (bool, error) {
	d, err := e.GetString(ctx, name)
	if err != nil {
		return false, err
	}

	parsebool, err := strconv.ParseBool(d)
	if err != nil {
		return false, err
	}

	return parsebool, nil
}

func (e EnvironmentParser) GetStringSlice(ctx context.Context, name string) ([]string, error) {
	d, err := e.GetString(ctx, name)
	if err != nil {
		return nil, err
	}

	return strings.Split(d, ","), nil
}
