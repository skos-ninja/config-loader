package parser

import (
	"context"
)

// FieldParser is an interface for fetching a field value given a tag value
type FieldParser interface {
	GetString(ctx context.Context, tagValue string) (string, error)
	GetInt(ctx context.Context, tagValue string) (int64, error)
	GetFloat(ctx context.Context, tagValue string) (float64, error)
	GetBoolean(ctx context.Context, tagValue string) (bool, error)
	GetStringSlice(ctx context.Context, tagValue string) ([]string, error)
}

// FieldParsers is a collection of tags to parsers for the parser to use
var FieldParsers = map[string]FieldParser{
	envTagName:  EnvironmentParser{},
	flagTagName: FlagParser{},
}
