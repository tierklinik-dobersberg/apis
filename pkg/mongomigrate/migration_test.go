package mongomigrate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tierklinik-dobersberg/apis/pkg/mongomigrate"
	"github.com/tierklinik-dobersberg/apis/pkg/mongotest"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestSingleMigration(t *testing.T) {
	ctx, cli := mongotest.Start(t)

	migrator := mongomigrate.NewMigrator(cli.Database("test"), "")

	called := false

	migrator.Register(mongomigrate.Migration{
		Version:     1,
		Description: "create a single object",
		Database:    "test",
		Up: mongomigrate.MigrateFunc(func(sc mongo.SessionContext, d *mongo.Database) error {
			called = true

			_, err := d.Collection("test").InsertOne(ctx, bson.M{"document": true})
			if err != nil {
				return err
			}
			return nil
		}),
	})

	err := migrator.Run(ctx)
	require.NoError(t, err)
	require.True(t, called)

	doc := cli.Database("test").Collection("test").FindOne(ctx, bson.M{"document": true})
	require.NoError(t, doc.Err())

	var res struct {
		Document bool `bson:"document"`
	}

	require.NoError(t, doc.Decode(&res))
	require.True(t, res.Document)
}
