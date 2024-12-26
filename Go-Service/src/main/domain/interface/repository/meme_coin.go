package repository

import (
	"Go-Service/src/main/domain/entity"
	"context"
)

type MemeCoinRepository interface {
	Create(ctx context.Context, memeCoin *entity.MemeCoin) error
}
