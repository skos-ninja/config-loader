package config

import (
	"context"
	"errors"
	"fmt"
	"strconv"
)

const flagTagName = "flag"

var ErrNotUsingCobraCtx = errors.New("flag parser called without cobra context")

type ErrFlagNotFound struct {
	flag string
}

func (e ErrFlagNotFound) Error() string {
	return fmt.Sprintf("Flag not found: %s", e.flag)
}

type FlagParser struct {
}

// GetString returns a environment variable as a string
func (p FlagParser) GetString(ctx context.Context, name string) (string, error) {
	cmd := getCmdFromContext(ctx)
	if cmd == nil {
		return "", ErrNotUsingCobraCtx
	}

	flag := cmd.Flag(name)
	if flag == nil {
		return "", ErrFlagNotFound{name}
	}

	if flag.Changed {
		return flag.Value.String(), nil
	}

	return "", ErrFlagNotFound{name}
}

// GetInt returns a environment variable as an integer
func (p FlagParser) GetInt(ctx context.Context, name string) (int64, error) {
	d, err := p.GetString(ctx, name)
	if err != nil {
		return 0, err
	}

	parseint, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return 0, err
	}

	return parseint, nil
}

// GetFloat returns a flag variable as a float
func (p FlagParser) GetFloat(ctx context.Context, name string) (float64, error) {
	d, err := p.GetString(ctx, name)
	if err != nil {
		return 0, err
	}

	parsefloat, err := strconv.ParseFloat(d, 64)
	if err != nil {
		return 0, err
	}

	return parsefloat, nil
}

// GetBoolean returns a flag variable as a boolean
func (p FlagParser) GetBoolean(ctx context.Context, name string) (bool, error) {
	d, err := p.GetString(ctx, name)
	if err != nil {
		return false, err
	}

	parsebool, err := strconv.ParseBool(d)
	if err != nil {
		return false, err
	}

	return parsebool, nil
}
