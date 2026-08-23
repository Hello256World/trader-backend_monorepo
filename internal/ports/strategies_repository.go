package ports

import (
	"context"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
)

type StrategiesRepository interface {
	Insert(context.Context, *domain.Strategy) (string, error)
	GetByID(context.Context, string) (*domain.Strategy, error)
}
