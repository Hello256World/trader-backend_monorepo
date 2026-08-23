package mongo

import (
	"errors"
	"fmt"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type StrategyDTO struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
}

func fromStrategyCoreToDTO(input *domain.Strategy) (*StrategyDTO, error) {
	if input == nil {
		return nil, errors.New("Invalid input Strategy")
	}

	id, err := bson.ObjectIDFromHex(input.ID)

	if err != nil {
		return nil, fmt.Errorf("invalid Strategy ID : %w", err)
	}

	dto := &StrategyDTO{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	}

	return dto, nil
}

func fromStrategyDTOToCore(input *StrategyDTO) domain.Strategy {
	result := domain.Strategy{
		ID:          input.ID.Hex(),
		Name:        input.Name,
		Description: input.Description,
	}

	return result
}
