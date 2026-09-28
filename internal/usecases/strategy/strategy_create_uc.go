package strategy

import (
	"context"
	"strings"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
	"github.com/Hello256World/trader-backend_monorepo/internal/ports"
	"github.com/Hello256World/trader-backend_monorepo/pkg/apierrors"
)

type StrategyCreateUC interface {
	Handle(c context.Context, req *CreateStrategyRequest) (*CreateStrategyResponse, error)
}

type strategyCreateUC struct {
	repo ports.StrategiesRepository
}

func NewStrategyCreateUC(repo ports.StrategiesRepository) StrategyCreateUC {
	return &strategyCreateUC{
		repo: repo,
	}
}

type CreateStrategyRequest struct {
	Name, Descrption string
}

type CreateStrategyResponse struct {
	StrategyID string
}

func (uc strategyCreateUC) Handle(c context.Context, req *CreateStrategyRequest) (*CreateStrategyResponse, error) {
	strategy, err := uc.fromCreateStrategyRequestToStrategy(req)

	if err != nil {
		return nil, err
	}

	strategyID, err := uc.repo.Insert(c, strategy)

	if err != nil {
		return nil, err
	}

	return &CreateStrategyResponse{StrategyID: strategyID}, nil
}

func (uc strategyCreateUC) fromCreateStrategyRequestToStrategy(req *CreateStrategyRequest) (*domain.Strategy, error) {
	if req == nil {
		return nil, apierrors.NewBadRequestError("invalid message")
	}

	if req.Name = strings.TrimSpace(req.Name); req.Name == "" {
		return nil, apierrors.NewBadRequestError("Invalid name")
	}

	return &domain.Strategy{Name: req.Name, Description: req.Descrption}, nil
}
