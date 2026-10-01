package strategy

import (
	"context"
	"strings"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
	"github.com/Hello256World/trader-backend_monorepo/internal/ports"
	"github.com/Hello256World/trader-backend_monorepo/pkg/apierrors"
)

type CreateUC interface {
	Handle(context.Context, *CreateRequest) (*CreateResponse, error)
}

type createUC struct {
	repo ports.StrategiesRepository
}

func NewCreateUC(repo ports.StrategiesRepository) CreateUC {
	return &createUC{
		repo: repo,
	}
}

type CreateRequest struct {
	Name, Description string
}

type CreateResponse struct {
	StrategyID string
}

func (uc createUC) Handle(c context.Context, req *CreateRequest) (*CreateResponse, error) {
	strategy, err := uc.fromCreateStrategyRequestToStrategy(req)

	if err != nil {
		return nil, err
	}

	strategyID, err := uc.repo.Insert(c, strategy)

	if err != nil {
		return nil, err
	}

	return &CreateResponse{StrategyID: strategyID}, nil
}

func (uc createUC) fromCreateStrategyRequestToStrategy(req *CreateRequest) (*domain.Strategy, error) {
	if req == nil {
		return nil, apierrors.NewBadRequestError("invalid message")
	}

	if req.Name = strings.TrimSpace(req.Name); req.Name == "" {
		return nil, apierrors.NewBadRequestError("Invalid name")
	}

	return &domain.Strategy{Name: req.Name, Description: req.Description}, nil
}
