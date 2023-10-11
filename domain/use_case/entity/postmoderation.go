package entity

import (
	"github.com/google/uuid"
)

type PostModeration struct {
	PostModerationID uuid.UUID `json:"post_id"`
	ModerationID     User      `json:"moderation_id"`
	UserID           User      `json:"user_id"`
	State            bool      `json:"state"`
}

func NewPostModeration(user, moderation User, state bool) *PostModeration {
	return &PostModeration{
		PostModerationID: uuid.New(),
		UserID:           user,
		ModerationID:     moderation,
		State:            state,
	}
}

func (moderation *PostModeration) GetType() string {
	return "post_moderation"
}

func (moderation *PostModeration) GetID() uuid.UUID {
	return moderation.PostModerationID
}
