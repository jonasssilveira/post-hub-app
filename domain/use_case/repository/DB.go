package repository

import (
	"PostHubApp/domain/use_case/dto"
	"PostHubApp/domain/use_case/entity"
	"context"
)

type DB interface {
	GetPost(ctx context.Context, id uint64) (dto.PostDTO, error)
	GetComment(ctx context.Context, id uint64) (entity.Comment, error)
	FindAllPost(ctx context.Context, dbOptions ...*DBOptions) ([]entity.Post, error)
	FindAllComment(ctx context.Context, dbOptions ...*DBOptions) ([]entity.Comment, error)
	MergePost(ctx context.Context, object entity.Migrations) error
	MergeComment(ctx context.Context, object entity.Migrations) error
}
