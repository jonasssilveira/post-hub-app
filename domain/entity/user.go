package entity

import (
	"github.com/google/uuid"
)

type User struct {
	UserID   string `gorm:"primaryKey;autoIncrement:true"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

func NewUser(email, password, fullname string) *User {
	return &User{
		UserID:   uuid.New().String(),
		Email:    email,
		Password: password,
		FullName: fullname,
	}
}

func (user *User) GetType() string {
	return "user"
}

func (user *User) GetID() string {
	return user.UserID
}
