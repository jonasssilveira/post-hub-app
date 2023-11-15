package entity

type Comment struct {
	CommentID uint64 `json:"comment_id" gorm:"primaryKey;autoIncrement:true"`
	User      uint64 `json:"user" gorm:"foreignKey:UserID"`
	Message   string `json:"message"`
	PostID    uint   `json:"postID" gorm:"size:36"`
}

func NewComment(user User, message string) *Comment {
	return &Comment{
		User:    user.UserID,
		Message: message,
	}
}

func (comment *Comment) GetType() string {
	return "comment"
}

func (comment *Comment) GetID() uint64 {
	return comment.CommentID
}
