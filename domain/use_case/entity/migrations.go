package entity

type Migrations interface {
	GetType() string
	GetID() uint64
}
