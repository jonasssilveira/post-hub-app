package database

import (
	"PostHubApp/domain/use_case/entity"
	"context"
	"gorm.io/gorm"
	"log"
)

type EntityManager[T entity.Migrations] struct {
	db *gorm.DB
}

func NewEntityManager[T entity.Migrations](db *gorm.DB) *EntityManager[T] {
	return &EntityManager[T]{db: db}
}

func (em *EntityManager[T]) Merge(ctx context.Context, migrated T) error {

	all, err := em.FindAll(ctx)

	if err != nil {
		return err
	}
	var result *gorm.DB
	if all != nil {
		result = em.db.WithContext(ctx).
			Table(migrated.GetType()).
			Create(&migrated)
	}

	if result.Error != nil {
		log.Fatal("an error ocured to save data, error " + result.Error.Error())
	}

}

func (em *EntityManager[T]) FindAll(ctx context.Context) ([]T, error) {
	var migratedFound []T

	updateResult := em.db.WithContext(ctx).Find(&migratedFound)

	if updateResult.Error != nil {
		return migratedFound, updateResult.Error
	}

	if len(migratedFound) == 0 {
		log.Println("the table is empty")
		return nil, nil
	}
	return migratedFound, nil

}
