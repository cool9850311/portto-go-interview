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

func (u *MemeCoinUsecase) Create(ctx context.Context, req *DTO.CreateMemeCoinRequestDTO) (string, error) {
	u.logger.Info(ctx, "Creating meme coin")
	memeCoin := entity.MemeCoin{
		Name:            req.Name,
		Description:     req.Description,
		CreatedAt:       time.Now(),
		PopularityScore: 0,
	}
	id, err := u.memeCoinRepository.Create(ctx, &memeCoin)
	if err != nil {
		u.logger.Error(ctx, "Error creating meme coin")
		return "", err
	}
	u.logger.Info(ctx, "Meme coin "+req.Name+" created successfully")
	return id, nil
}

func (u *MemeCoinUsecase) GetByID(ctx context.Context, req *DTO.GetMemeCoinRequestDTO) (DTO.GetMemeCoinResponseDTO, error) {
	memeCoin, err := u.memeCoinRepository.GetByID(ctx, req.ID)
	if err != nil {
		u.logger.Error(ctx, "Error getting meme coin by ID")
		return DTO.GetMemeCoinResponseDTO{}, err
	}
	return DTO.GetMemeCoinResponseDTO{
		Name:            memeCoin.Name,
		Description:     memeCoin.Description,
		CreatedAt:       memeCoin.CreatedAt,
		PopularityScore: memeCoin.PopularityScore,
	}, nil
}

func (u *MemeCoinUsecase) Update(ctx context.Context, req *DTO.UpdateMemeCoinRequestDTO) error {
	err := u.memeCoinRepository.Update(ctx, req.Description, req.ID)
	if err != nil {
		u.logger.Error(ctx, "Error updating meme coin")
		return err
	}
	return nil
}

func (u *MemeCoinUsecase) Delete(ctx context.Context, req *DTO.DeleteMemeCoinRequestDTO) error {
	err := u.memeCoinRepository.Delete(ctx, req.ID)
	if err != nil {
		u.logger.Error(ctx, "Error deleting meme coin")
		return err
	}
	return nil
}

func (u *MemeCoinUsecase) Poke(ctx context.Context, req *DTO.PokeMemeCoinRequestDTO) error {
	err := u.memeCoinRepository.Poke(ctx, req.ID)
	if err != nil {
		u.logger.Error(ctx, "Error poking meme coin")
		return err
	}
	return nil
}
