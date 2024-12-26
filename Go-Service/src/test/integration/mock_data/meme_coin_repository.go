package mock_data

import (
	"Go-Service/src/main/domain/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

type MockMemeCoinRepository struct {
	mock.Mock
}

func (m *MockMemeCoinRepository) Create(ctx context.Context, memeCoin *entity.MemeCoin) (string, error) {
	args := m.Called(ctx, memeCoin)
	return args.String(0), args.Error(1)
}

func (m *MockMemeCoinRepository) GetByID(ctx context.Context, id string) (*entity.MemeCoin, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.MemeCoin), args.Error(1)
}

func (m *MockMemeCoinRepository) Update(ctx context.Context, description string, id string) error {
	args := m.Called(ctx, description, id)
	return args.Error(0)
}

func (m *MockMemeCoinRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockMemeCoinRepository) Poke(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
