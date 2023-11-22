package dto

type Comment struct {
	CommentID uint64
	UserID    uint64 `json:"user_id"`
	Message   string `json:"message"`
	PostID    uint64 `json:"post_id"`
}
