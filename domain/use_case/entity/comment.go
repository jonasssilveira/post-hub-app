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

func (comment *Comment) GetType() string {
	return "comment"
}

func (comment *Comment) GetID() uuid.UUID {
	return comment.CommentID
}
