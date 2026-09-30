package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hello256World/trader-backend_monorepo/internal/config"
	mongodb "github.com/Hello256World/trader-backend_monorepo/pkg/db/mongo"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Application interface {
	Run(ctx context.Context) error
}

type application struct {
	conf        *config.Config
	httpEngine  *gin.Engine
	mongoClient *mongo.Client
}

func NewApplication(ctx context.Context) (Application, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	client, err := mongodb.NewClient(ctx, conf.MongoConfig)
	if err != nil {
		return nil, err
	}

	// strategiesColl := mongoClient.Database(conf.MongoConfig.Database).Collection("strategies")
	// strategiesRepo := mongoAdapters.NewStrategiesRepository(strategiesColl)
	// strategiesSvc := strategiesUC.NewService(strategiesRepo)
	// handlers := strategiesHTTP.NewHandlers(strategiesSvc)

	// engine := gin.Default()
	// strategiesHTTP.RegisterRoutes(engine, handlers)

	return &application{
		conf:        conf,
		httpEngine:  gin.Default(),
		mongoClient: client,
	}, nil
}

func (a *application) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.conf.HTTPConfig.Port),
		Handler: a.httpEngine,
	}

	errChan := make(chan error, 1)
	go func() {
		log.Printf("server listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// ErrServerClosed یعنی خودمون Shutdown صدا زدیم؛ خطای واقعی نیست
			errChan <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		return fmt.Errorf("server failed to start: %w", err)
	case <-ctx.Done():
		log.Println("context cancelled, shutting down")
	case sig := <-quit:
		log.Printf("received signal %s, shutting down", sig)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown failed: %w", err)
	}

	if err := a.mongoClient.Disconnect(shutdownCtx); err != nil {
		return fmt.Errorf("error disconnecting mongo: %w", err)
	}

	log.Println("server shut down gracefully")
	return nil
}
