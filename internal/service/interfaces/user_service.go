package interfaces

import (
	"context"

	dto "github.com/faizalnurrozi/go-starter-kit/internal/dto/request"
	response "github.com/faizalnurrozi/go-starter-kit/internal/dto/response"
)

type UserService interface {
	Create(ctx context.Context, req *dto.CreateUserRequest) (*response.UserResponse, error)
	GetByID(ctx context.Context, id uint) (*response.UserResponse, error)
	GetAll(ctx context.Context, limit, offset int) ([]*response.UserResponse, error)
	Update(ctx context.Context, id uint, req *dto.UpdateUserRequest) (*response.UserResponse, error)
	Delete(ctx context.Context, id uint) error
}
