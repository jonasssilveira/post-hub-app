package entity

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
)

type Post struct {
	PostID       string `gorm:"primaryKey;"`
	Title        string `json:"title"`
	Message      string `json:"message"`
	UserID       string `json:"user_id"`
	ModerationID string `json:"moderation_id" gorm:"-"`
}

func (post Post) ToMessage() []byte {
	jsonData, err := json.Marshal(post)
	if err != nil {
		log.Fatalf("Error encoding struct to JSON: %s", err)
	}
	return jsonData
}

func NewPost(title, message string, user string) Post {
	return Post{
		PostID:  uuid.New().String(),
		Title:   title,
		Message: message,
		UserID:  user,
	}
}

func (post Post) GetType() string {
	return "post"
}

func (post Post) GetID() string {
	return post.PostID
}

// TableName specifies the table name for the Post model
func (Post) TableName() string {
	return "post" // Set the table name to match the actual table name in your database
}
