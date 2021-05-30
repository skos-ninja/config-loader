package config

import (
	"context"
	"errors"
	"log"
	"reflect"
)

// FieldParser is an interface for fetching a field value given a tag value
type FieldParser interface {
	GetString(ctx context.Context, tagValue string) (string, error)
	GetInt(ctx context.Context, tagValue string) (int64, error)
	GetFloat(ctx context.Context, tagValue string) (float64, error)
	GetBoolean(ctx context.Context, tagValue string) (bool, error)
}

// FieldParsers is a collection of tags to parsers for the parser to use
var FieldParsers = map[string]FieldParser{
	envTagName:  EnvironmentParser{},
	flagTagName: FlagParser{},
}

// ParseStruct takes a struct ptr and iterates through the fields and applies any field parsers
func ParseStruct(ctx context.Context, s interface{}, failOnParseError bool) error {
	rv := reflect.ValueOf(s)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("Struct must be a pointer and not nil")
	}

	v := rv.Elem()

	for i := 0; i < v.NumField(); i++ {
		kind := v.Field(i).Kind()

		if kind == reflect.Struct {
			inter := v.Field(i).Addr().Interface()
			err := ParseStruct(ctx, inter, failOnParseError)
			if err != nil {
				return err
			}
			continue
		}

		for k, f := range FieldParsers {
			tag := v.Type().Field(i).Tag.Get(k)

			// Skip if tag is not defined or ignored
			if tag == "" || tag == "-" {
				continue
			}

			switch kind {
			case reflect.String:
				value, err := f.GetString(ctx, tag)
				if err != nil {
					if failOnParseError {
						return err
					}
					continue
				}
				if value != "" {
					v.Field(i).SetString(value)
				} else {
					break
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value, err := f.GetInt(ctx, tag)
				if err != nil {
					if failOnParseError {
						return err
					}
					continue
				}
				if value != 0 {
					v.Field(i).SetInt(value)
				} else {
					break
				}
			case reflect.Float64, reflect.Float32:
				value, err := f.GetFloat(ctx, tag)
				if err != nil {
					if failOnParseError {
						return err
					}
					continue
				}
				if value != 0 {
					v.Field(i).SetFloat(value)
				} else {
					break
				}
			case reflect.Bool:
				value, err := f.GetBoolean(ctx, tag)
				if err != nil {
					if failOnParseError {
						return err
					}
					continue
				}
				if value {
					v.Field(i).SetBool(value)
				} else {
					break
				}
			default:
				log.Printf("WARNING: Unsupported type found in struct: %s\n", v.Type().Field(i).Name)
			}
		}
	}

	return nil
}
