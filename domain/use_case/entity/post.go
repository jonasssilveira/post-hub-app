package entity

import (
	"github.com/google/uuid"
)

type Post struct {
	PostID   uuid.UUID `json:"post_id"`
	Title    string    `json:"title"`
	Message  string    `json:"message"`
	User     User      `json:"user"`
	Comments []Comment `json:"comments"`
}

func NewPost(title, message string, user User) *Post {
	return &Post{
		PostID:  uuid.New(),
		Title:   title,
		Message: message,
		User:    user,
	}
}

func (post *Post) AddComment(comment Comment) {
	post.Comments = append(post.Comments, comment)
}
