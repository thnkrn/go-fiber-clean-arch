package handler

import (
	"github.com/google/uuid"
	"github.com/thnkrn/go-fiber-crud-clean-arch/pkg/domain"
)

type UserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func NewUserResponse(user domain.User) *UserResponse {
	response := UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return &response
}

func NewUsesrResponse(users []domain.User) *[]UserResponse {
	response := make([]UserResponse, len(users))

	for i, v := range users {
		response[i] = UserResponse{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		}
	}

	return &response
}
