package strategy

import "github.com/Hello256World/trader-backend_monorepo/internal/domain"

func fromStrategyCoreToHTTP(strategy *domain.Strategy) *Strategy {
	if strategy == nil {
		return nil
	}

	return &Strategy{
		ID:          strategy.ID,
		Name:        strategy.Name,
		Description: strategy.Description,
	}
}
