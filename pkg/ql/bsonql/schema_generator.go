package bsonql

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/tierklinik-dobersberg/apis/pkg/ql"
)

type FieldNameResolverFunc func(field reflect.StructField) string

func TagNameResolver(tagName string) FieldNameResolverFunc {
	return func(field reflect.StructField) string {
		tag := field.Tag.Get(tagName)
		if tag == "" {
			return ""
		}

		parts := strings.Split(tag, ",")
		if parts[0] == "" || parts[0] == "-" {
			return ""
		}

		return parts[0]
	}
}

var (
	BSONTagNameResolver = TagNameResolver("bson")
	JSONTagNameResolver = TagNameResolver("json")
	YAMLTagNameResolver = TagNameResolver("yaml")
)

func SchemaFromModel(model any, nameResolver FieldNameResolverFunc) (ql.FieldList, error) {
	val := reflect.ValueOf(model).Type()

	switch val.Kind() {
	case reflect.Ptr, reflect.Interface:
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct type but got %s", val.String())
	}

	spec, err := generateSchema("", val, nameResolver)
	if err != nil {
		return nil, err
	}

	return spec.Children.(ql.FieldList), nil
}

func generateSchema(name string, val reflect.Type, nameResolver FieldNameResolverFunc) (*ql.FieldSpec, error) {
	switch val.Kind() {
	case reflect.Ptr, reflect.Interface:
		val = val.Elem()
	}

	spec := &ql.FieldSpec{
		Name: name,
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int8, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		spec.TypeResolver = ql.IntType()

	case reflect.Float32, reflect.Float64:
		spec.TypeResolver = ql.FloatType()

	case reflect.String:
		// no type resolver needed

	case reflect.Struct:
		list := ql.FieldList{}
		for idx := 0; idx < val.NumField(); idx++ {
			field := val.Field(idx)

			if !field.IsExported() {
				continue
			}

			fieldName := field.Name
			if nameResolver != nil {
				fieldName = nameResolver(field)
			}

			if fieldName == "" {
				continue
			}

			if name != "" {
				fieldName = fmt.Sprintf("%s.%s", name, field.Name)
			}

			fieldSpec, err := generateSchema(fieldName, field.Type, nameResolver)
			if err != nil {
				return spec, err
			}

			if fieldSpec != nil {
				list = append(list, fieldSpec)
			}
		}

		spec.Children = list

	default:
		return nil, fmt.Errorf("unsupported value type %s", val.String())
	}

	return spec, nil
}
