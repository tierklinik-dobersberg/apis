package ql

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	commonv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
	"github.com/tierklinik-dobersberg/apis/pkg/timeutil"
)

type (
	ValueProvider interface {
		ProvideValues(ctx context.Context, prefix string) ([]string, error)
	}

	ValueProviderFunc func(context.Context, string) ([]string, error)

	TypeResolver interface {
		ResolveType(value string) (any, error)
	}

	TypeResolverFunc func(string) (any, error)

	CombinedTypeResolver struct {
		Resolvers []TypeResolver
	}

	FieldProvider interface {
		LookupField(name string) *FieldSpec
	}

	FieldList []FieldProvider

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

		// TypeResolver is called to get the concrete typed value for this field.
		// It is only called if Children is unset.
		TypeResolver TypeResolver

		Children FieldProvider

		// Data might be used to attach addtional data to a FieldSpec.
		Data map[string]any
	}
)

func (vpf ValueProviderFunc) ProvideValues(ctx context.Context, prefix string) ([]string, error) {
	return vpf(ctx, prefix)
}

func (fl FieldList) LookupField(lit string) *FieldSpec {
	for _, spec := range fl {
		if match := spec.LookupField(lit); match != nil {
			return match
		}
	}

	return nil
}

func (spec FieldSpec) LookupField(name string) *FieldSpec {
	l := strings.ToLower(name)

	if strings.ToLower(spec.Name) == l {
		return &spec
	}

	for _, a := range spec.Aliases {
		if strings.ToLower(a) == l {
			return &spec
		}
	}

	if spec.Children != nil {
		return spec.Children.LookupField(name)
	}

	return nil
}

func (cr *CombinedTypeResolver) ResolveType(s string) (any, error) {
	merr := new(multierror.Error)

	for _, r := range cr.Resolvers {
		res, err := r.ResolveType(s)
		if err == nil {
			return res, nil
		}

		merr.Errors = append(merr.Errors, err)
	}

	return nil, merr.ErrorOrNil()
}

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

func DateOnlyType() TypeResolver {
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

func TimeStartKeywordType(loc *time.Location) TypeResolver {
	return TypeResolverFunc(func(s string) (any, error) {
		t, err := timeutil.ResolveTime(s, time.Now())
		if err == nil {
			return t, nil
		}
		return timeutil.ParseStartInLocation(s, loc)
	})
}

func TimeEndKeywordType(loc *time.Location) TypeResolver {
	return TypeResolverFunc(func(s string) (any, error) {
		t, err := timeutil.ResolveTime(s, time.Now())
		if err == nil {
			return t, nil
		}

		return timeutil.ParseEndInLocation(s, loc)
	})
}

func DayTimeType(loc *time.Location) TypeResolver {
	return TypeResolverFunc(func(s string) (any, error) {
		return commonv1.ParseDayTime(s)
	})
}
