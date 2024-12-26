package mock_data

import (
	"Go-Service/src/main/domain/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

type MockMemeCoinRepository struct {
	mock.Mock
}

func (m *MockMemeCoinRepository) Create(ctx context.Context, memeCoin *entity.MemeCoin) error {
	args := m.Called(ctx, memeCoin)
	return args.Error(0)
}
