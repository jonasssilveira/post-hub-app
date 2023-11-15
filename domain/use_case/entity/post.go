package entity

type Post struct {
	PostID       uint64 `gorm:"primaryKey;autoIncrement"`
	Title        string `json:"title"`
	Message      string `json:"message"`
	UserID       uint64 `json:"user_id"`
	ModerationID uint64 `json:"moderation_id"`
}

func NewPost(title, message string, user User) *Post {
	return &Post{
		Title:   title,
		Message: message,
	}
}

func (post Post) GetType() string {
	return "post"
}

func (post Post) GetID() uint64 {
	return post.PostID
}

// TableName specifies the table name for the Post model
func (Post) TableName() string {
	return "post" // Set the table name to match the actual table name in your database
}
