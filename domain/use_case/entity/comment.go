package entity

import (
	"github.com/google/uuid"
)

type Comment struct {
	CommentID uuid.UUID `json:"comment_id"`
	User      User      `json:"user"`
	Message   string    `json:"message"`
}

func NewComment(user User, message string) *Comment {
	return &Comment{
		CommentID: uuid.New(),
		User:      user,
		Message:   message,
	}
}
