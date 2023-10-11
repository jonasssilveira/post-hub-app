package entity

import (
	"github.com/google/uuid"
)

type Migrations interface {
	Comment | Post | PostModeration | User
	GetType() string
	GetID() uuid.UUID
}
