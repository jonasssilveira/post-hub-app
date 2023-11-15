package dto

import (
	"PostHubApp/domain/use_case/entity"
)

type post struct {
	entity.Post
	entity.User
}

func NewDTOPost() *post {
	return &post{}
}

func (post post) GetType() string {
	return "post"
}

func (post post) GetID() uint64 {
	return post.PostID
}
