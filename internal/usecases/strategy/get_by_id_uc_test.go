package strategy

import (
	"context"
	"errors"
	"testing"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
	"github.com/Hello256World/trader-backend_monorepo/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStrategyRepository_GetByID(t *testing.T) {
	t.Parallel()

	t.Run("nil value pass error", func(t *testing.T) {
		t.Parallel()
		strategy := &getByIDUC{}

		res, err := strategy.Handle(context.Background(), nil)

		assert.EqualError(t, err, "Nil req passed")
		assert.Nil(t, res)
	})

	t.Run("error get strategy by id", func(t *testing.T) {
		t.Parallel()

		repoMock := mocks.NewStrategiesRepositoryMock(t)

		repoMock.On("GetByID", mock.Anything, mock.Anything).Return(nil, errors.New("This is the error")).Once()

		strategy := NewGetByIDUC(repoMock)

		res, err := strategy.Handle(context.Background(), &GetByIDRequest{})

		assert.Nil(t, res)
		assert.EqualError(t, err, "This is the error")
	})

	t.Run("successfully get strategy by id", func(t *testing.T) {
		t.Parallel()

		repoMock := mocks.NewStrategiesRepositoryMock(t)

		repoMock.On("GetByID", mock.Anything, mock.Anything).Return(&domain.Strategy{
			ID:   "strategy-id",
			Name: "my name",
		}, nil).Once()

		strategy := NewGetByIDUC(repoMock)

		res, err := strategy.Handle(context.Background(), &GetByIDRequest{ID: "strategy-id"})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.EqualValues(t, "strategy-id", res.Strategy.ID)
		assert.EqualValues(t, "my name", res.Strategy.Name)
		assert.EqualValues(t, "", res.Strategy.Description)
	})
}
