package entity

import (
	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	FullName string    `json:"full_name"`
}

func NewUser(email, password, fullname string) *User {
	return &User{
		UserID:   uuid.New(),
		Email:    email,
		Password: password,
		FullName: fullname,
	}
}

func (user *User) GetType() string {
	return "user"
}

func (user *User) GetID() uuid.UUID {
	return user.UserID
}
