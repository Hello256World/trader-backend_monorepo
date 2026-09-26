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
	strategyResponse string
}

func (su strategyCreateUC) Handle(c context.Context, req *CreateStrategyRequest) (*CreateStrategyResponse, error) {
	strategy, err := su.fromCreateStrategyRequestToStrategy(req)

	if err != nil {
		return nil, err
	}

	strategyID, err := su.repo.Insert(c, strategy)

	if err != nil {
		return nil, err
	}

	return &CreateStrategyResponse{strategyResponse: strategyID}, nil
}

func (su strategyCreateUC) fromCreateStrategyRequestToStrategy(req *CreateStrategyRequest) (*domain.Strategy, error) {
	if req == nil {
		return nil, apierrors.NewBadRequestError("invalid message")
	}

	if req.Name = strings.TrimSpace(req.Name); req.Name == "" {
		return nil, apierrors.NewBadRequestError("Invalid message")
	}

	return &domain.Strategy{Name: req.Name, Description: req.Descrption}, nil
}
