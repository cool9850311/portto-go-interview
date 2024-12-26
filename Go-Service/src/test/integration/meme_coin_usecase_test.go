package integration

import (
	"Go-Service/src/main/application/DTO"
	"Go-Service/src/main/application/usecase"
	"Go-Service/src/test/integration/mock_data"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
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
	setup.mockMemeCoinRepository.On("Create", mock.Anything, mock.Anything).Return(nil)
	err := setup.usecase.Create(context.Background(), &DTO)
	assert.NoError(t, err)

}
