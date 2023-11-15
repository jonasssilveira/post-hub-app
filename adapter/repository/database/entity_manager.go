package database

import (
	"PostHubApp/domain/use_case/dto"
	"PostHubApp/domain/use_case/entity"
	"PostHubApp/domain/use_case/repository"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
)

type EntityManager struct {
	db *gorm.DB
}

func NewEntityManager(db *gorm.DB) *EntityManager {
	return &EntityManager{db: db}
}

func (em *EntityManager) Merge(ctx context.Context, migrated entity.Migrations) error {

	found, err := em.Get(ctx, migrated.GetID())

	if err != nil {
		return err
	}

	if found.GetID() == 0 {
		return em.db.Create(migrated).Error
	} else {
		return em.update(migrated, ctx)
	}

}

func (em *EntityManager) FindAll(ctx context.Context, dbOptions ...*repository.DBOptions) ([]entity.Migrations, error) {
	var migratedFound []entity.Migrations

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

func (em *EntityManager) Get(ctx context.Context, id uint64) (entity.Migrations, error) {
	migratedFound := dto.NewDTOPost()

	err := em.db.WithContext(ctx).Select("post.*, user.*"). // Select the columns you need
								Joins("LEFT JOIN user ON post.user_id = user.user_id").
								Where("post.post_id = ?", id).
								Find(&migratedFound).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return *migratedFound, nil
		} else {
			return *migratedFound, err
		}

	}
	return *migratedFound, nil
}

func (em *EntityManager) update(found entity.Migrations, ctx context.Context) error {
	var result *gorm.DB
	if result.Error != nil {
		log.Fatal("an error occurred to save data, error " + result.Error.Error())
		return errors.New("generic error for now")
	}

	if found != nil {
		result = em.db.WithContext(ctx).
			Table(found.GetType()).
			Create(&found)
		return nil
	} else {
		return errors.New("generic error for now")
	}

}
