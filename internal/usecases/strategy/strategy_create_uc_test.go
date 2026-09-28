package strategy

import (
	"context"
	"errors"
	"testing"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
	"github.com/Hello256World/trader-backend_monorepo/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_CreateStrategyUC_Handle(t *testing.T) {
	t.Parallel()

	t.Run("invalid request returns error", func(t *testing.T) {
		t.Parallel()
		strategyUC := strategyCreateUC{}

		res, err := strategyUC.Handle(context.Background(), &CreateStrategyRequest{})

		assert.Nil(t, res)
		assert.EqualError(t, err, "Invalid name")
	})

	t.Run("create strategy return error", func(t *testing.T) {
		t.Parallel()
		mockRepository := &mocks.ManualStrategyRepositoryMock{}

		mockRepository.InsertFn = func(ctx context.Context, s *domain.Strategy) (string, error) {
			return "", errors.New("we have and error")
		}

		strategyUC := strategyCreateUC{
			repo: mockRepository,
		}

		res, err := strategyUC.Handle(context.Background(), &CreateStrategyRequest{Name: "Alex"})

		assert.Nil(t, res)
		assert.EqualError(t, err, "we have and error")

	})

	t.Run("successfully create strategy", func(t *testing.T) {
		t.Parallel()
		mockRepository := &mocks.ManualStrategyRepositoryMock{}

		mockRepository.InsertFn = func(ctx context.Context, s *domain.Strategy) (string, error) {
			return "strategy-id", nil
		}

		strategyUC := strategyCreateUC{
			repo: mockRepository,
		}

		res, err := strategyUC.Handle(context.Background(), &CreateStrategyRequest{Name: "Alex"})

		assert.NotNil(t, res)
		assert.NoError(t, err)
		assert.Equal(t, res.StrategyID, "strategy-id")
	})
}
