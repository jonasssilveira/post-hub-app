package repository

import (
	"PostHubApp/domain/use_case/entity"
)

type DB[T entity.Migrations] interface {
	Get(id string) T
	FindAll(dbOptions ...DBOptions) []T
	Merge(object T) error
}
