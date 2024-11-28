package ql

import (
	"strconv"
	"time"
)

type (
	ValueProvider interface {
		ProvideValues(prefix string) ([]string, error)
	}

	TypeResolver interface {
		ResolveType(value string) (any, error)
	}

	TypeResolverFunc func(string) (any, error)

	FieldSpec struct {
		// Name is the name of the field.
		Name string

		// Aliases holds field aliases.
		Aliases []string

		// ValueProvider is called to get possible values
		// for the field. The function is passed the current
		// user input.
		ValueProvider ValueProvider

		// Description holds the description for the field.
		Description string

		TypeResolver TypeResolver
	}
)

func (trf TypeResolverFunc) ResolveType(s string) (any, error) {
	return trf(s)
}

func BooleanType() TypeResolver {
	return TypeResolverFunc(func(s string) (any, error) {
		b, err := strconv.ParseBool(s)
		return b, err
	})
}

func IntType() TypeResolver {
	return TypeResolverFunc(func(s string) (any, error) {
		v, err := strconv.ParseInt(s, 10, 0)
		return v, err
	})
}

func FloatType() TypeResolver {
	return TypeResolverFunc(func(s string) (any, error) {
		f, err := strconv.ParseFloat(s, 0)
		return f, err
	})
}

func TimeType() TypeResolver {
	return TypeResolverFunc(func(s string) (any, error) {
		t, err := time.Parse(time.RFC3339, s)
		return t, err
	})
}

func DateType() TypeResolver {
	return TypeResolverFunc(func(s string) (any, error) {
		d, err := time.Parse(time.DateOnly, s)
		return d, err
	})
}

func TimeFormatType(format string) TypeResolver {
	return TypeResolverFunc(func(s string) (any, error) {
		t, err := time.Parse(format, s)
		return t, err
	})
}
