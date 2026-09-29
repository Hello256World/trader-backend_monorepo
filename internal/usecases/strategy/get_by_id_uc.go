package strategy

import (
	"context"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
	"github.com/Hello256World/trader-backend_monorepo/internal/ports"
	"github.com/Hello256World/trader-backend_monorepo/pkg/apierrors"
)

type GetByIDUC interface {
	Handle(context.Context, *GetByIDRequest) (*GetByIDResponse, error)
}

type getByIDUC struct {
	repo ports.StrategiesRepository
}

func NewGetByIDUC(repo ports.StrategiesRepository) GetByIDUC {
	return &getByIDUC{
		repo: repo,
	}
}

type GetByIDRequest struct {
	ID string
}

type GetByIDResponse struct {
	Strategy *domain.Strategy
}

func (uc *getByIDUC) Handle(c context.Context, req *GetByIDRequest) (*GetByIDResponse, error) {
	if req == nil {
		return nil, apierrors.NewBadRequestError("Nil req passed")
	}

	strategy, err := uc.repo.GetByID(c, req.ID)
	if err != nil {
		return nil, err
	}

	return &GetByIDResponse{
		Strategy: strategy,
	}, nil
}
