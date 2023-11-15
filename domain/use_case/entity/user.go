package entity

type User struct {
	UserID   uint64 `gorm:"primaryKey;autoIncrement:true"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Posts    []Post // Define a one-to-many relationship
}

func NewUser(email, password, fullname string) *User {
	return &User{
		Email:    email,
		Password: password,
		FullName: fullname,
	}
}

func (user *User) GetType() string {
	return "user"
}

func (user *User) GetID() uint64 {
	return user.UserID
}
