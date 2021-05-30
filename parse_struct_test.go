package config

import (
	"context"
	"strconv"
	"testing"
)

func TestParse(t *testing.T) {
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
