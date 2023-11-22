package dto

type Post struct {
	Title        string `json:"title"`
	Message      string `json:"message"`
	UserID       uint64 `json:"user_id"`
	ModerationID uint64 `json:"moderation_id"`
}
