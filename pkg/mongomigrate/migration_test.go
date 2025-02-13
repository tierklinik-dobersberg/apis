package mongomigrate_test

import (
	"errors"
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

func TestMultipleMigrationsInOrder(t *testing.T) {
	ctx, cli := mongotest.Start(t)

	db := cli.Database("test")

	called := 0

	migrator := mongomigrate.NewMigrator(db, "")
	migrator.Register(
		mongomigrate.Migration{
			Version:     2,
			Description: "Second migration",
			Database:    "test",
			Up: mongomigrate.MigrateFunc(func(sc mongo.SessionContext, d *mongo.Database) error {
				if called != 1 {
					t.Errorf("expected first migration to executed")
				}
				called++
				_, err := d.Collection("test").UpdateOne(ctx, bson.M{"document": true}, bson.M{"$set": bson.M{"foo": "bar"}})
				return err
			}),
		},

		mongomigrate.Migration{
			Version:     1,
			Description: "First migration",
			Database:    "test",
			Up: mongomigrate.MigrateFunc(func(sc mongo.SessionContext, d *mongo.Database) error {
				if called != 0 {
					t.Errorf("expected no migrations yet")
				}
				called++
				_, err := d.Collection("test").InsertOne(ctx, bson.M{"document": true})
				return err
			}),
		},
	)

	err := migrator.Run(ctx)
	require.NoError(t, err)

	require.Equal(t, called, 2)

	doc := cli.Database("test").Collection("test").FindOne(ctx, bson.M{"document": true})
	require.NoError(t, doc.Err())

	var res struct {
		Document bool   `bson:"document"`
		Foo      string `bson:"foo"`
	}

	require.NoError(t, doc.Decode(&res))
	require.True(t, res.Document)
	require.Equal(t, res.Foo, "bar")
}

func TestRollback(t *testing.T) {
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
			require.NoError(t, err)

			return errors.New("simulated")
		}),
	})

	err := migrator.Run(ctx)
	require.Error(t, err)
	require.True(t, called)

	doc := cli.Database("test").Collection("test").FindOne(ctx, bson.M{"document": true})
	require.Error(t, doc.Err())
}
