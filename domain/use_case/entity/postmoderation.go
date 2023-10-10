package entity

import (
	"github.com/google/uuid"
)

type PostModeration struct {
	PostID       uuid.UUID `json:"post_id"`
	ModerationID User      `json:"moderation_id"`
	UserID       User      `json:"user_id"`
	State        bool      `json:"state"`
}

func NewPostModeration(user, moderation User, state bool) *PostModeration {
	return &PostModeration{
		PostID:       uuid.New(),
		UserID:       user,
		ModerationID: moderation,
		State:        state,
	}
}
