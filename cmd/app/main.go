package main

import (
	"context"
	"log"
	"time"

	"github.com/Hello256World/trader-backend_monorepo/internal/app"
	"github.com/Hello256World/trader-backend_monorepo/internal/config"
	"github.com/Hello256World/trader-backend_monorepo/pkg/db/mongo"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error getting config : %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.NewClient(ctx, conf.MongoConfig)

	if err != nil {
		log.Fatalf("error connecting to mongo : %v", err)
	}

	application := app.NewApplication(conf, client)

	if err := application.Run(); err != nil {
		log.Fatalf("error running server : %v", err)
	}
}
