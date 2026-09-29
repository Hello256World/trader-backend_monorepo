package strategy

import (
	"context"

	"github.com/Hello256World/trader-backend_monorepo/internal/ports"
)

type Service interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	GetByID(context.Context, *GetByIDRequest) (*GetByIDResponse, error)
}

type service struct {
	createUC  CreateUC
	getByIDUC GetByIDUC
}

func NewService(repo ports.StrategiesRepository) Service {
	return &service{
		createUC:  NewCreateUC(repo),
		getByIDUC: NewGetByIDUC(repo),
	}
}

func (s *service) Create(c context.Context, req *CreateRequest) (*CreateResponse, error) {
	return s.createUC.Handle(c, req)
}

func (s *service) GetByID(c context.Context, req *GetByIDRequest) (*GetByIDResponse, error) {
	return s.getByIDUC.Handle(c, req)
}
