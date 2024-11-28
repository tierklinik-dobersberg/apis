package bsonql

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tierklinik-dobersberg/apis/pkg/ql"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBSONQL_Basic(t *testing.T) {
	query := "completed = true AND dueTime <= now"

	completedField := ql.FieldSpec{
		Name:         "completed",
		TypeResolver: ql.BooleanType(),
	}
	dueTimeField := ql.FieldSpec{
		Name: "dueTime",
	}

	q := &BSONQL{
		Schema: ql.FieldList{
			completedField,
			dueTimeField,
		},
	}

	res, err := q.Parse(query)
	require.NoError(t, err)

	expected := bson.M{
		"$and": bson.A{
			bson.M{
				"completed": bson.M{
					"$eq": true,
				},
			},
			bson.M{
				"dueTime": bson.M{
					"$lte": "now",
				},
			},
		},
	}

	require.Equal(t, expected, res)
}
