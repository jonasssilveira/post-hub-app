package entity

type Migrations interface {
	GetType() string
	GetID() string
	ToMessage() []byte
}
