package entity

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
)

type Comment struct {
	CommentID string `gorm:"primaryKey;autoIncrement:true"`
	UserID    string `gorm:"foreignKey:user_id"`
	Message   string
	PostID    string `gorm:"size:36"`
}

func (comment *Comment) ToMessage() []byte {
	jsonData, err := json.Marshal(comment)
	if err != nil {
		log.Fatalf("Error encoding struct to JSON: %s", err)
	}
	return jsonData
}

func NewComment(user, post string, message string) Comment {
	return Comment{
		CommentID: uuid.New().String(),
		UserID:    user,
		PostID:    post,
		Message:   message,
	}
}

func (comment *Comment) GetType() string {
	return "comment"
}

func (comment *Comment) GetID() string {
	return comment.CommentID
}

// TableName specifies the table name for the Post model
func (Comment) TableName() string {
	return "comment" // Set the table name to match the actual table name in your database
}
