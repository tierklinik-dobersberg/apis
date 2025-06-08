package service

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseCreator[T any] interface {
	ConfigureDatabase(context.Context) (T, error)
}

type NoDatabase struct{}

func (NoDatabase) ConfigureDatabase(context.Context) (any, error) {
	return nil, nil
}

type MongoConfig struct {
	URL      string `env:"URL, required"`
	Database string `env:"DATABASE, required"`
}

func (mc MongoConfig) ConfigureDatabase(ctx context.Context) (*mongo.Database, error) {
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(mc.URL))
	if err != nil {
		return nil, err
	}

	return cli.Database(mc.Database), nil
}
