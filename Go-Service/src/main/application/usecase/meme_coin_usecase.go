package usecase

import (
	"Go-Service/src/main/application/DTO"
	"Go-Service/src/main/domain/entity"
	"Go-Service/src/main/domain/interface/logger"
	"Go-Service/src/main/domain/interface/repository"
	"context"
	"time"
)

type MemeCoinUsecase struct {
	logger             logger.Logger
	memeCoinRepository repository.MemeCoinRepository
}

func NewMemeCoinUsecase(logger logger.Logger, memeCoinRepository repository.MemeCoinRepository) *MemeCoinUsecase {
	return &MemeCoinUsecase{logger: logger, memeCoinRepository: memeCoinRepository}
}

func (u *MemeCoinUsecase) Create(ctx context.Context, req *DTO.CreateMemeCoinRequestDTO) error {
	u.logger.Info(ctx, "Creating meme coin")
	memeCoin := entity.MemeCoin{
		Name:            req.Name,
		Description:     req.Description,
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}
	err := u.memeCoinRepository.Create(ctx, &memeCoin)
	if err != nil {
		u.logger.Error(ctx, "Error creating meme coin")
		return err
	}
	u.logger.Info(ctx, "Meme coin "+req.Name+" created successfully")
	return nil
}
