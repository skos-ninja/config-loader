package parser

import (
	"context"
	"errors"
	"fmt"

	c "github.com/skos-ninja/config-loader/pkg/context"

	"github.com/spf13/pflag"
)

const flagTagName = "flag"

var ErrNotUsingCobraCtx = errors.New("flag parser called without cobra context")

type ErrFlagNotFound struct {
	flag string
}

func (e ErrFlagNotFound) Error() string {
	return fmt.Sprintf("Flag not found: %s", e.flag)
}

var ErrFlagsNotFound = errors.New("flags not found in cobra context")

type FlagParser struct {
}

func (p FlagParser) getFlags(ctx context.Context) (*pflag.FlagSet, error) {
	cmd := c.GetCmdFromContext(ctx)
	if cmd == nil {
		return nil, ErrNotUsingCobraCtx
	}

	flags := cmd.Flags()
	if flags == nil {
		return nil, ErrFlagsNotFound
	}

	return flags, nil
}

// GetString returns a environment variable as a string
func (p FlagParser) GetString(ctx context.Context, name string) (string, error) {
	flags, err := p.getFlags(ctx)
	if err != nil {
		return "", err
	}

	if flags.Changed(name) {
		return flags.GetString(name)
	}

	return "", ErrFlagNotFound{name}
}

// GetInt returns a environment variable as an integer
func (p FlagParser) GetInt(ctx context.Context, name string) (int64, error) {
	flags, err := p.getFlags(ctx)
	if err != nil {
		return 0, err
	}

	if flags.Changed(name) {
		return flags.GetInt64(name)
	}

	return 0, ErrFlagNotFound{name}
}

// GetFloat returns a flag variable as a float
func (p FlagParser) GetFloat(ctx context.Context, name string) (float64, error) {
	flags, err := p.getFlags(ctx)
	if err != nil {
		return 0, err
	}

	if flags.Changed(name) {
		return flags.GetFloat64(name)
	}

	return 0, ErrFlagNotFound{name}
}

// GetBoolean returns a flag variable as a boolean
func (p FlagParser) GetBoolean(ctx context.Context, name string) (bool, error) {
	flags, err := p.getFlags(ctx)
	if err != nil {
		return false, err
	}

	if flags.Changed(name) {
		return flags.GetBool(name)
	}

	return false, ErrFlagNotFound{name}
}

// GetStringSlice returns a flag variable as a string slice
func (p FlagParser) GetStringSlice(ctx context.Context, name string) ([]string, error) {
	flags, err := p.getFlags(ctx)
	if err != nil {
		return nil, err
	}

	if flags.Changed(name) {
		return flags.GetStringSlice(name)
	}

	return nil, ErrFlagNotFound{name}
}
