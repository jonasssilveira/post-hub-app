package entity

type Comment struct {
	CommentID uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserID    uint64 `gorm:"foreignKey:user_id"`
	Message   string
	PostID    uint64 `gorm:"size:36"`
}

func NewComment(user, post uint64, message string) *Comment {
	return &Comment{
		UserID:  user,
		PostID:  post,
		Message: message,
	}
}

func (comment *Comment) GetType() string {
	return "comment"
}

func (comment *Comment) GetID() uint64 {
	return comment.CommentID
}

// TableName specifies the table name for the Post model
func (Comment) TableName() string {
	return "comment" // Set the table name to match the actual table name in your database
}
