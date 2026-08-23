package mongo

import (
	"context"
	"fmt"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
	"github.com/Hello256World/trader-backend_monorepo/internal/ports"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type strategiesRepository struct {
	db *mongo.Collection
}

func NewStrategiesRepository(db *mongo.Collection) ports.StrategiesRepository {
	return &strategiesRepository{db: db}
}

func (repo *strategiesRepository) Insert(c context.Context, strategy *domain.Strategy) (string, error) {
	dto, err := fromStrategyCoreToDTO(strategy)
	if err != nil {
		return "", err
	}

	result, err := repo.db.InsertOne(c, dto)
	if err != nil {
		return "", fmt.Errorf("Error insert strategy to db %w", err)
	}

	id, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return "", fmt.Errorf("Error converting inserted ID (Strategy) to ObjectID")
	}

	return id.Hex(), nil
}

func (repo *strategiesRepository) GetByID(c context.Context, id string) (*domain.Strategy, error) {
	return nil, nil
}
