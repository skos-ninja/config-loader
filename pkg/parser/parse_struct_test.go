package parser

import (
	"context"
	"strconv"
	"testing"

	c "github.com/skos-ninja/config-loader/pkg/context"

	"github.com/spf13/cobra"
)

func TestEnvParse(t *testing.T) {
	ctx := context.Background()

	type Nested struct {
		Field string `env:"test-nested"`
	}
	type Test struct {
		Str     string  `env:"test-string"`
		Inter   int     `env:"test-int"`
		Float   float64 `env:"test-float"`
		Boolean bool    `env:"test-boolean"`
		Nested  Nested
	}

	original := &Test{
		Str:     "",
		Inter:   0,
		Float:   0,
		Boolean: false,
		Nested: Nested{
			Field: "",
		},
	}
	expected := &Test{
		Str:     "test-value",
		Inter:   100,
		Float:   100.10,
		Boolean: true,
		Nested: Nested{
			Field: "test-string",
		},
	}
	setEnv(t, env{name: "test-string", value: expected.Str})
	setEnv(t, env{name: "test-int", value: strconv.Itoa(expected.Inter)})
	setEnv(t, env{name: "test-float", value: strconv.FormatFloat(expected.Float, 'g', -1, 64)})
	setEnv(t, env{name: "test-boolean", value: strconv.FormatBool(expected.Boolean)})
	setEnv(t, env{name: "test-nested", value: expected.Nested.Field})

	err := ParseStruct(ctx, original, true)
	if err != nil {
		t.Error(err)
	}

	if *original != *expected {
		t.Errorf("parseStruct() = %v, want %v", *original, *expected)
	}
}

func TestFlagParse(t *testing.T) {
	type Nested struct {
		Field string `flag:"test-nested"`
	}
	type Test struct {
		Str     string  `flag:"test-string"`
		Inter   int     `flag:"test-int"`
		Float   float64 `flag:"test-float"`
		Boolean bool    `flag:"test-boolean"`
		Nested  Nested
	}

	original := &Test{
		Str:     "",
		Inter:   0,
		Float:   0,
		Boolean: false,
		Nested: Nested{
			Field: "",
		},
	}
	expected := &Test{
		Str:     "test-value",
		Inter:   100,
		Float:   100.10,
		Boolean: true,
		Nested: Nested{
			Field: "test-string",
		},
	}
	cmd := &cobra.Command{Use: "test"}
	setFlag(cmd, flag{name: "test-string", value: expected.Str})
	setFlag(cmd, flag{name: "test-int", value: strconv.Itoa(expected.Inter)})
	setFlag(cmd, flag{name: "test-float", value: strconv.FormatFloat(expected.Float, 'g', -1, 64)})
	setFlag(cmd, flag{name: "test-boolean", value: strconv.FormatBool(expected.Boolean)})
	setFlag(cmd, flag{name: "test-nested", value: expected.Nested.Field})
	ctx := c.GetContextWithCmd(cmd)

	err := ParseStruct(ctx, original, true)
	if err != nil {
		t.Error(err)
	}

	if *original != *expected {
		t.Errorf("parseStruct() = %v, want %v", *original, *expected)
	}
}

func TestMultipleTags(t *testing.T) {
	type Test struct {
		Str string `env:"test-string" flag:"test-string"`
	}

	original := &Test{
		Str: "",
	}
	expected := &Test{
		Str: "test-value",
	}
	cmd := &cobra.Command{Use: "test"}
	setEnv(t, env{name: "test-string", value: expected.Str})
	setFlag(cmd, flag{name: "test-string", value: ""})
	ctx := c.GetContextWithCmd(cmd)

	err := ParseStruct(ctx, original, false)
	if err != nil {
		t.Error(err)
	}

	if *original != *expected {
		t.Errorf("parseStruct() = %v, want %v", *original, *expected)
	}
}
