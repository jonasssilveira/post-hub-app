package entity

type PostModeration struct {
	PostModerationID uint64 `gorm:"primaryKey;autoIncrement"`
	ModerationID     uint64 `json:"moderation_id"`
	PostID           uint64 `json:"post_id"`
	State            bool   `json:"state"`
}

func NewPostModeration(user, moderation User, state bool, post Post) *PostModeration {
	return &PostModeration{
		PostID:       post.PostID,
		ModerationID: moderation.UserID,
		State:        state,
	}
}

func (moderation *PostModeration) GetType() string {
	return "post_moderation"
}

func (moderation *PostModeration) GetID() uint64 {
	return moderation.PostModerationID
}
