package bsonql

import (
	"context"
	"fmt"
	"strings"

	"github.com/tierklinik-dobersberg/apis/pkg/ql"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddValueProviders(col *mongo.Collection, fields ql.FieldList) {
	for _, field := range fields {
		spec, ok := field.(*ql.FieldSpec)
		if !ok || spec.ValueProvider != nil {
			continue
		}

		spec.ValueProvider = ql.ValueProviderFunc(func(ctx context.Context, prefix string) ([]string, error) {
			values, err := col.Distinct(ctx, spec.Name, nil)
			if err != nil {
				return nil, err
			}

			res := make([]string, 0, len(values))
			for _, val := range values {
				var stringValue string

				if s, ok := val.(string); ok {
					stringValue = s
				} else {
					stringValue = fmt.Sprintf("%v", val)
				}

				if prefix != "" && !strings.HasPrefix(strings.ToLower(stringValue), strings.ToLower(prefix)) {
					continue
				}

				res = append(res, stringValue)
			}

			return res, nil
		})
	}
}
