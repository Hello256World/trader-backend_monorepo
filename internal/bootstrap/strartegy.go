package bootstrap

import (
	strategiesHTTP "github.com/Hello256World/trader-backend_monorepo/internal/adapters/http/strategy"
	mongoAdapters "github.com/Hello256World/trader-backend_monorepo/internal/adapters/mongo"
	"github.com/Hello256World/trader-backend_monorepo/internal/usecases/strategy"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func WireStrategy(engine *gin.Engine, db *mongo.Database) {
	strategiesColl := db.Collection("strategies")
	strategiesRepo := mongoAdapters.NewStrategiesRepository(strategiesColl)
	strategiesSvc := strategy.NewService(strategiesRepo)
	handlers := strategiesHTTP.NewHandlers(strategiesSvc)

	strategiesHTTP.RegisterRoutes(engine, handlers)
}
