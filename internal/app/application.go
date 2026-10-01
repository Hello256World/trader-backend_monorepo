package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Hello256World/trader-backend_monorepo/internal/config"
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

func NewApplication(config *config.Config, client *mongo.Client) Application {
	a := &application{
		conf:        config,
		httpEngine:  gin.Default(),
		mongoClient: client,
	}

	a.mapHandlers()

	return a
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
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return fmt.Errorf("server failed to start: %w", err)
	case <-ctx.Done():
		log.Printf("received signal %s, shutting down", "")
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
