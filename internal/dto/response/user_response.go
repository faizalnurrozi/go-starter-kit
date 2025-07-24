package dto

import (
	"time"

	"github.com/faizalnurrozi/go-starter-kit/internal/entity"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewUserListResponse(users []entity.User) []*UserResponse {
	response := make([]*UserResponse, len(users))
	for i, user := range users {
		response[i] = NewUserResponse(&user)
	}
	return response
}
