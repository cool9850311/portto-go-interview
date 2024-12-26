package repository

import (
	"Go-Service/src/main/domain/entity"
	"context"
)

type MemeCoinRepository interface {
	Create(ctx context.Context, memeCoin *entity.MemeCoin) (string, error)
	GetByID(ctx context.Context, id string) (*entity.MemeCoin, error)
	Update(ctx context.Context, description string, id string) error
	Delete(ctx context.Context, id string) error
	Poke(ctx context.Context, id string) error
}
