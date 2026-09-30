package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hello256World/trader-backend_monorepo/internal/config"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Application interface {
	Run() error
}

type application struct {
	conf        *config.Config
	httpEngine  *gin.Engine
	mongoClient *mongo.Client
}

func NewApplication(config *config.Config, client *mongo.Client) Application {
	return &application{
		conf:        config,
		httpEngine:  gin.Default(),
		mongoClient: client,
	}
}

func (a *application) Run() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.conf.HTTPConfig.Port),
		Handler: a.httpEngine,
	}

	a.mapHandlers()

	errChan := make(chan error, 1)
	go func() {
		log.Printf("server listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-errChan:
		return fmt.Errorf("server failed to start: %w", err)
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
