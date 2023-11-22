package dto

import (
	"PostHubApp/domain/use_case/entity"
)

type PostDTO struct {
	entity.Post `json:"post"`
	entity.User `json:"user"`
	comment     []entity.Comment `json:"comment"`
}

type Post struct {
	Title        string `json:"title"`
	Message      string `json:"message"`
	UserID       uint64 `json:"user_id"`
	ModerationID uint64 `json:"moderation_id"`
}

func NewDTOPost() *PostDTO {
	return &PostDTO{}
}

// TableName specifies the table name for the Post model
func (PostDTO) TableName() string {
	return "post" // Set the table name to match the actual table name in your database
}
func (post PostDTO) GetType() string {
	return "post"
}

func (post PostDTO) GetID() uint64 {
	return post.PostID
}
