package bsonql

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/tierklinik-dobersberg/apis/pkg/ql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func TestWithMongo(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForExposedPort(),
	}

	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, mongoC)
	require.NoError(t, err)

	ep, err := mongoC.Endpoint(ctx, "")
	require.NoError(t, err)

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", ep)))
	require.NoError(t, err)

	err = cli.Ping(ctx, nil)
	require.NoError(t, err)

	col := cli.Database("test").Collection("foo")

	// insert some documents
	_, err = col.InsertMany(ctx, []any{
		Doc{
			Name: "foo",
		},
		Doc{
			Name: "foo",
			A:    "bar",
			I:    10,
		},
		Doc{
			Name: "foo",
			A:    "bar",
		},
		Doc{
			Name: "foo",
			I:    10,
		},
	})
	require.NoError(t, err)

	// create the query
	res, err := (&BSONQL{
		Schema: ql.FieldList{
			ql.FieldSpec{
				Name: "name",
			},
			ql.FieldSpec{
				Name: "a",
			},
			ql.FieldSpec{
				Name:         "i",
				TypeResolver: ql.IntType(),
			},
		},
	}).Parse("name = foo AND ( a = bar OR i = 10 )")
	require.NoError(t, err)

	findRes, err := col.Find(ctx, res)
	require.NoError(t, err)

	var result []Doc
	require.NoError(t, findRes.All(ctx, &result))

	require.Equal(t, []Doc{}, result)
}

type Doc struct {
	Name string `bson:"name"`
	A    string `bson:"a"`
	I    int    `bson:"i"`
}
