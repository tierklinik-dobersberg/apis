package mongotest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Start(t *testing.T) (context.Context, *mongo.Client) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	mongoC, err := mongodb.Run(ctx, "mongo:latest",
		mongodb.WithUsername("root"),
		mongodb.WithPassword("example"),
		mongodb.WithReplicaSet("rs0"),
	)
	require.NoError(t, err)

	ep, err := mongoC.Endpoint(ctx, "")
	require.NoError(t, err)

	ep = fmt.Sprintf("mongodb://root:example@%s/?replicaSet=rs0", ep)

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(ep).SetServerSelectionTimeout(time.Second*5))
	require.NoError(t, err)

	require.NoError(t, cli.Ping(ctx, nil))

	return ctx, cli
}
