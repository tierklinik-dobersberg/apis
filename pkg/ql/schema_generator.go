package ql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	ErrSkipField = errors.New("skip field")
)

type (
	FieldNameResolverFunc func(parent string, field reflect.StructField) string
)

func TagNameResolver(tagName string) FieldNameResolverFunc {
	return func(parent string, field reflect.StructField) string {
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

func SchemaFromModel(model any, nameResolver FieldNameResolverFunc) (FieldList, error) {
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

	return spec.Children.(FieldList), nil
}

func generateSchema(name string, val reflect.Type, nameResolver FieldNameResolverFunc) (*FieldSpec, error) {
	switch val.Kind() {
	case reflect.Ptr, reflect.Interface:
		val = val.Elem()
	}

	spec := &FieldSpec{
		Name: name,
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int8, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		spec.TypeResolver = IntType()

	case reflect.Float32, reflect.Float64:
		spec.TypeResolver = FloatType()

	case reflect.String:
		// no type resolver needed

	case reflect.Struct:
		list := FieldList{}
		for idx := 0; idx < val.NumField(); idx++ {
			field := val.Field(idx)

			if !field.IsExported() {
				continue
			}

			fieldName := field.Name

			if nameResolver != nil {
				fieldName = nameResolver(name, field)
			} else if name != "" {
				fieldName = fmt.Sprintf("%s.%s", name, field.Name)
			}

			if fieldName == "" {
				continue
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
	}

	return spec, nil
}
