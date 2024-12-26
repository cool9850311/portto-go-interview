package integration

import (
	"Go-Service/src/main/application/DTO"
	"Go-Service/src/main/application/usecase"
	"Go-Service/src/main/domain/entity"
	"Go-Service/src/test/integration/mock_data"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MemeCoinUsecaseTestSetup struct {
	mockLogger             *mock_data.MockLogger
	mockMemeCoinRepository *mock_data.MockMemeCoinRepository
	usecase                usecase.MemeCoinUsecase
}

func setup() MemeCoinUsecaseTestSetup {
	mockLogger := new(mock_data.MockLogger)
	mockMemeCoinRepository := new(mock_data.MockMemeCoinRepository)
	usecase := usecase.NewMemeCoinUsecase(mockLogger, mockMemeCoinRepository)
	return MemeCoinUsecaseTestSetup{
		mockLogger:             mockLogger,
		mockMemeCoinRepository: mockMemeCoinRepository,
		usecase:                *usecase,
	}
}

func TestMemeCoinCreate(t *testing.T) {
	setup := setup()
	DTO := DTO.CreateMemeCoinRequestDTO{
		Name:        "Test",
		Description: "Test",
	}
	setup.mockMemeCoinRepository.On("Create", mock.Anything, mock.Anything).Return("123", nil)
	_, err := setup.usecase.Create(context.Background(), &DTO)
	assert.NoError(t, err)
}

func TestMemeCoinGetByID(t *testing.T) {
	setup := setup()
	expectedMemeCoin := entity.MemeCoin{
		Name:            "Test",
		Description:     "Test",
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}
	setup.mockMemeCoinRepository.On("GetByID", mock.Anything, "123").Return(&expectedMemeCoin, nil)
	response, err := setup.usecase.GetByID(context.Background(), &DTO.GetMemeCoinRequestDTO{ID: "123"})
	assert.NoError(t, err)
	assert.Equal(t, expectedMemeCoin.Name, response.Name)
	assert.Equal(t, expectedMemeCoin.Description, response.Description)
}

func TestMemeCoinUpdate(t *testing.T) {
	setup := setup()
	setup.mockMemeCoinRepository.On("Update", mock.Anything, "Updated Description", "123").Return(nil)
	err := setup.usecase.Update(context.Background(), &DTO.UpdateMemeCoinRequestDTO{ID: "123", Description: "Updated Description"})
	assert.NoError(t, err)
}

func TestMemeCoinDelete(t *testing.T) {
	setup := setup()
	setup.mockMemeCoinRepository.On("Delete", mock.Anything, "123").Return(nil)
	err := setup.usecase.Delete(context.Background(), &DTO.DeleteMemeCoinRequestDTO{ID: "123"})
	assert.NoError(t, err)
}

func TestMemeCoinPoke(t *testing.T) {
	setup := setup()
	setup.mockMemeCoinRepository.On("Poke", mock.Anything, "123").Return(nil)
	err := setup.usecase.Poke(context.Background(), &DTO.PokeMemeCoinRequestDTO{ID: "123"})
	assert.NoError(t, err)
}
