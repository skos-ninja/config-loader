package parser

import (
	"context"
	"errors"
	"log"
	"reflect"
)

// ParseStruct takes a struct ptr and iterates through the fields and applies any field parsers
func ParseStruct(ctx context.Context, s interface{}, failOnParseError bool) error {
	rv := reflect.ValueOf(s)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("struct must be a pointer and not nil")
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
			case reflect.Slice:
				switch v.Field(i).Type().Elem().Kind() {
				case reflect.String:
					s, err := f.GetStringSlice(ctx, tag)
					if err != nil {
						if failOnParseError {
							return err
						}
						continue
					}
					v.Field(i).Set(reflect.ValueOf(s))
				default:
					log.Printf("WARNING: Unsupported slice type found in struct: %s\n", v.Field(i).Type().Elem().Kind())
				}
			case reflect.String:
				value, err := f.GetString(ctx, tag)
				if err != nil {
					if failOnParseError {
						return err
					}
					continue
				}
				v.Field(i).SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				value, err := f.GetInt(ctx, tag)
				if err != nil {
					if failOnParseError {
						return err
					}
					continue
				}
				v.Field(i).SetInt(value)
			case reflect.Float64, reflect.Float32:
				value, err := f.GetFloat(ctx, tag)
				if err != nil {
					if failOnParseError {
						return err
					}
					continue
				}
				v.Field(i).SetFloat(value)
			case reflect.Bool:
				value, err := f.GetBoolean(ctx, tag)
				if err != nil {
					if failOnParseError {
						return err
					}
					continue
				}
				v.Field(i).SetBool(value)
			default:
				log.Printf("WARNING: Unsupported type found in struct: %s\n", v.Type().Field(i).Name)
			}
		}
	}

	return nil
}
