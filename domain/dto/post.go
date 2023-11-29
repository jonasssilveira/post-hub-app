package dto

import (
	entity2 "PostHubApp/domain/entity"
)

type PostDTO struct {
	entity2.Post `json:"post"`
	entity2.User `json:"user"`
	comment      []entity2.Comment `json:"comment"`
}

type Post struct {
	Title        string `json:"title"`
	Message      string `json:"message"`
	UserID       string `json:"user_id"`
	ModerationID string `json:"moderation_id"`
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

func (post PostDTO) GetID() string {
	return post.PostID
}
