package repository

import (
	"PostHubApp/domain/dto"
	entity2 "PostHubApp/domain/entity"
	"context"
)

type DB interface {
	GetPost(ctx context.Context, id string) (dto.PostDTO, error)
	GetModedationFromPost(ctx context.Context, id string) (entity2.PostModeration, error)
	GetComment(ctx context.Context, id string) (entity2.Comment, error)
	FindAllPost(ctx context.Context, dbOptions ...*DBOptions) ([]entity2.Post, error)
	FindAllComment(ctx context.Context, dbOptions ...*DBOptions) ([]entity2.Comment, error)
	MergePost(ctx context.Context, object entity2.Migrations) error
	MergeComment(ctx context.Context, object entity2.Migrations) error
}
