package repository

import (
	"PostHubApp/domain/use_case/entity"
	"context"
)

type DB interface {
	Get(ctx context.Context, id uint64) (entity.Migrations, error)
	FindAll(ctx context.Context, dbOptions ...*DBOptions) ([]entity.Migrations, error)
	Merge(ctx context.Context, object entity.Migrations) error
}
