package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

func NewUser(id uuid.UUID, name, email string) User {
	return User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
