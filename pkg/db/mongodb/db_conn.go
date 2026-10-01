package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Options struct {
	URI         string
	AppName     string
	MinPoolSize uint64
	MaxPoolSize uint64
}

func Connect(ctx context.Context, opts Options) (*mongo.Client, error) {
	mongoOpts := options.Client().
		SetAppName(opts.AppName).
		ApplyURI(opts.URI).
		SetMinPoolSize(opts.MinPoolSize).
		SetMaxPoolSize(opts.MaxPoolSize)

	mongoClient, err := mongo.Connect(mongoOpts)
	if err != nil {
		return nil, fmt.Errorf("error connecting to mongo: %w", err)
	}

	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("error pinging mongo: %w", err)
	}

	return mongoClient, nil
}
