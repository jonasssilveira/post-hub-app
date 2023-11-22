package entity

type Post struct {
	PostID       uint64 `gorm:"primaryKey;autoIncrement"`
	Title        string
	Message      string
	UserID       uint64
	ModerationID *uint64
}

func NewPost(title, message string, user uint64) *Post {
	return &Post{
		Title:        title,
		Message:      message,
		UserID:       user,
		ModerationID: nil,
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
