package mocks

import (
	"context"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
)

type ManualStrategyRepositoryMock struct {
	InsertFn  func(context.Context, *domain.Strategy) (string, error)
	GetByIDFn func(context.Context, string) (*domain.Strategy, error)
}

func (m *ManualStrategyRepositoryMock) Insert(c context.Context, strategy *domain.Strategy) (string, error) {
	return m.InsertFn(c, strategy)
}

func (m *ManualStrategyRepositoryMock) GetByID(c context.Context, id string) (*domain.Strategy, error) {
	return m.GetByIDFn(c, id)
}
