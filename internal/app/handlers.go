package app

import (
	strategiesHTTP "github.com/Hello256World/trader-backend_monorepo/internal/adapters/http/strategy"
	"github.com/Hello256World/trader-backend_monorepo/internal/adapters/mongo"
	"github.com/Hello256World/trader-backend_monorepo/internal/usecases/strategy"
)

func (a *application) mapHandlers() {
	strategiesColl := a.mongoClient.Database(a.conf.MongoConfig.Database).Collection("strategies")
	strategiesRepo := mongo.NewStrategiesRepository(strategiesColl)
	strategiesSvc := strategy.NewService(strategiesRepo)
	handlers := strategiesHTTP.NewHandlers(strategiesSvc)

	strategiesHTTP.RegisterRoutes(a.httpEngine, handlers)
}
