package database

import (
	"PostHubApp/domain/dto"
	entity "PostHubApp/domain/entity"
	"PostHubApp/domain/repository"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
)

type EntityManager struct {
	db *gorm.DB
}

func (em *EntityManager) GetModedationFromPost(ctx context.Context, id string) (entity.PostModeration, error) {
	//TODO implement me
	panic("implement me")
}

func NewEntityManager(db *gorm.DB) *EntityManager {
	return &EntityManager{db: db}
}

func (em *EntityManager) MergePost(ctx context.Context, post entity.Migrations) error {

	found, err := em.GetPost(ctx, post.GetID())

	if err != nil {
		return err
	}

	if found.GetID() == "" {
		d := em.db.Create(post)

		return d.Error
	} else {
		return em.update(post, ctx)
	}

}

func (em *EntityManager) MergeComment(ctx context.Context, comment entity.Migrations) error {

	found, err := em.GetComment(ctx, comment.GetID())

	if err != nil {
		return err
	}

	if found.PostID == "" {
		return em.db.Create(comment).Error
	} else {
		return em.update(comment, ctx)
	}

}

func (em *EntityManager) FindAllComment(ctx context.Context, dbOptions ...*repository.DBOptions) ([]entity.Comment, error) {
	var commentFound []entity.Comment

	updateResult := em.db.WithContext(ctx).Find(&commentFound)

	if updateResult.Error != nil {
		return commentFound, updateResult.Error
	}

	if len(commentFound) == 0 {
		log.Println("the table is empty")
		return nil, nil
	}
	return commentFound, nil

}

func (em *EntityManager) FindAllPost(ctx context.Context, dbOptions ...*repository.DBOptions) ([]entity.Post, error) {
	var migratedFound []entity.Post

	updateResult := em.db.WithContext(ctx).Select("post.*, user.*, comment.*").
		Joins("JOIN user ON post.user_id = user.user_id").
		Joins("JOIN comment ON post.id = comment.post_id").
		Find(&migratedFound).Error

	if updateResult != nil {
		if errors.Is(updateResult, gorm.ErrRecordNotFound) {
			return migratedFound, nil
		} else {
			return migratedFound, updateResult
		}
	}

	if len(migratedFound) == 0 {
		log.Println("the table is empty")
		return nil, nil
	}
	return migratedFound, nil

}

func (em *EntityManager) GetPost(ctx context.Context, id string) (dto.PostDTO, error) {
	var migratedFound dto.PostDTO

	err := em.db.WithContext(ctx).Select("post.*, user.*").
		Joins("JOIN user ON post.user_id = user.user_id").
		Where("post.post_id = ?", id).
		Find(&migratedFound).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return migratedFound, nil
		} else {
			return migratedFound, err
		}
	}

	return migratedFound, nil
}
func (em *EntityManager) GetComment(ctx context.Context, id string) (entity.Comment, error) {
	var migratedFound entity.Comment

	err := em.db.WithContext(ctx).Select("post.*, user.*, comment.*").
		Joins("JOIN user ON comment.user_id = user.user_id").
		Joins("JOIN post ON comment.post_id = post.post_id").
		Where("comment.comment_id = ?", id).
		Find(&migratedFound).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return migratedFound, nil
		} else {
			return migratedFound, err
		}
	}
	return migratedFound, nil
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
