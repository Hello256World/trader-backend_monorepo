package mongo

import (
	"context"
	"fmt"

	"github.com/Hello256World/trader-backend_monorepo/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewClient(ctx context.Context, conf config.MongoConfig) (*mongo.Client, error) {
	mongoOpts := options.Client().
		SetAppName(conf.AppName).
		ApplyURI(conf.GetConnectionURI()).
		SetMinPoolSize(uint64(conf.MinPoolSize)).
		SetMaxPoolSize(uint64(conf.MaxPoolSize))

	mongoClient, err := mongo.Connect(mongoOpts)
	if err != nil {
		return nil, fmt.Errorf("error connecting to mongo: %w", err)
	}

	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("error pinging mongo: %w", err)
	}

	return mongoClient, nil
}
