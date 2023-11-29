package use_case

import (
	entity2 "PostHubApp/domain/entity"
	"PostHubApp/domain/repository"
	"context"
	"errors"
)

type ServicePost struct {
	db      repository.DB
	produce ServiceProducer
}

func NewServicePost(db repository.DB, producer ServiceProducer) *ServicePost {
	return &ServicePost{
		db:      db,
		produce: producer,
	}
}

func (service *ServicePost) Comment(ctx context.Context, comment entity2.Comment) error {

	post, err := service.db.GetPost(ctx, comment.PostID)

	if err != nil {
		return err
	}

	moderation, err := service.db.GetModedationFromPost(ctx, post.PostID)

	if &moderation != nil || moderation.State == false {
		return errors.New("it's not possible comment on a post banned")
	}
	service.db.MergeComment(ctx, &comment)
	return nil
}

func (service *ServicePost) SavePost(ctx context.Context, post entity2.Post) {
	service.produce.Produce(post)
	service.db.MergePost(ctx, post)
}
