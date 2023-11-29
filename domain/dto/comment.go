package dto

type Comment struct {
	CommentID string
	UserID    string `json:"user_id"`
	Message   string `json:"message"`
	PostID    string `json:"post_id"`
}
