package strategy

import (
	"context"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
	"github.com/Hello256World/trader-backend_monorepo/internal/ports"
	"github.com/Hello256World/trader-backend_monorepo/pkg/apierrors"
)

type StrategyGetByIDUC interface {
	Handle(context.Context, *GetStrategyByIDRequest) (*GetStrategyByIDResponse, error)
}

type strategyGetByIDUC struct {
	repo ports.StrategiesRepository
}

func NewStrategyGetByIDUC(repo ports.StrategiesRepository) StrategyGetByIDUC {
	return &strategyGetByIDUC{
		repo: repo,
	}
}

type GetStrategyByIDRequest struct {
	ID string
}

type GetStrategyByIDResponse struct {
	Strategy *domain.Strategy
}

func (uc *strategyGetByIDUC) Handle(c context.Context, req *GetStrategyByIDRequest) (*GetStrategyByIDResponse, error) {
	if req == nil {
		return nil, apierrors.NewBadRequestError("Nil req passed")
	}

	strategy, err := uc.repo.GetByID(c, req.ID)
	if err != nil {
		return nil, err
	}

	return &GetStrategyByIDResponse{
		Strategy: strategy,
	}, nil
}
