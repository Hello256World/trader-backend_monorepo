package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hello256World/trader-backend_monorepo/internal/app"
	"github.com/Hello256World/trader-backend_monorepo/internal/config"
	"github.com/Hello256World/trader-backend_monorepo/pkg/db/mongodb"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	conf, err := config.GetConfig()
	if err != nil {
		return fmt.Errorf("error getting config : %w", err)
	}

	startupCtx, cancelStratup := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelStratup()

	client, err := mongodb.Connect(startupCtx, mongodb.Options{
		URI:         conf.MongoConfig.GetConnectionURI(),
		AppName:     conf.MongoConfig.AppName,
		MinPoolSize: uint64(conf.MongoConfig.MinPoolSize),
		MaxPoolSize: uint64(conf.MongoConfig.MaxPoolSize),
	})

	if err != nil {
		return fmt.Errorf("error connecting to mongo : %v", err)
	}

	application := app.NewApplication(conf, client)

	runCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := application.Run(runCtx); err != nil {
		return fmt.Errorf("error running server : %v", err)
	}

	return nil
}
