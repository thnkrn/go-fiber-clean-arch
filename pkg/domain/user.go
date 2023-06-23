package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	Name       string
	Email      string
	Versioning int
}

func NewUser(id uuid.UUID, name, email string) User {
	return User{
		ID:         id,
		Name:       name,
		Email:      email,
		Versioning: 1,
	}
}

func (u *User) Version() int {
	return u.Versioning
}

func (u *User) SetVersion(i int) {
	u.Versioning = i
}
