package entity

type PostModeration struct {
	PostModerationID string `gorm:"primaryKey;autoIncrement"`
	ModerationID     string `json:"moderation_id"`
	PostID           string `json:"post_id"`
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

func (moderation *PostModeration) GetID() string {
	return moderation.PostModerationID
}
