package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
	"github.com/Hello256World/trader-backend_monorepo/internal/ports"
	"github.com/Hello256World/trader-backend_monorepo/pkg/apierrors"
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
		return "", apierrors.NewBadRequestError(err.Error())
	}

	result, err := repo.db.InsertOne(c, dto)
	if err != nil {
		return "", apierrors.NewInternalServerError("Error Creating Strategy")
	}

	id, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return "", apierrors.NewInternalServerError("error getting inserted id")
	}

	return id.Hex(), nil
}

func (repo *strategiesRepository) GetByID(c context.Context, id string) (*domain.Strategy, error) {
	mongoId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, apierrors.NewBadRequestError(fmt.Sprintf("invalid Strategy ID: '%v'", id))
	}

	filter := bson.M{
		"_id": mongoId,
	}

	res := repo.db.FindOne(c, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, apierrors.NewNotFoundError(fmt.Sprintf("strategy '%v' not found", id))
		}

		return nil, apierrors.NewInternalServerError("error getting strategy by id")
	}

	var strategyDTO StrategyDTO

	if err := res.Decode(strategyDTO); err != nil {
		return nil, apierrors.NewInternalServerError(fmt.Sprintf("error getting strategy '%s'", id))
	}

	strategy := fromStrategyDTOToCore(strategyDTO)

	return &strategy, nil
}
