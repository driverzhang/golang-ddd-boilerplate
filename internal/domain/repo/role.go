package repo

import (
	"context"
	"golang-ddd-boilerplate/internal/domain/entity"
)

// IRoleRepo I代表interface
type IRoleRepo interface {
	Get(ctx context.Context, roleID string) (entity.Entity, error)
}
