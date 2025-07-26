package handlers

import (
	"context"
	"time"

	dto "github.com/faizalnurrozi/go-starter-kit/internal/dto/request"
	"github.com/faizalnurrozi/go-starter-kit/internal/service/interfaces"
	pb "github.com/faizalnurrozi/go-starter-kit/proto/user"
)

type userHandler struct {
	pb.UnimplementedUserServiceServer
	userService interfaces.UserService
}

// NewUserHandler returns an implementation of pb.UserServiceServer
func NewUserHandler(userService interfaces.UserService) pb.UserServiceServer {
	return &userHandler{
		userService: userService,
	}
}

// CreateUser implements pb.UserServiceServer
func (h *userHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	// Simulasi: Kamu bisa panggil h.userService.Create jika sudah tersedia
	dtoReq := &dto.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	// Panggil service
	user, err := h.userService.Create(ctx, dtoReq)
	if err != nil {
		return nil, err
	}

	// Map ke gRPC response
	return &pb.UserResponse{
		Id:        uint32(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil

}

// GetUser implements pb.UserServiceServer
func (h *userHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	// Panggil service
	user, err := h.userService.GetByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	// Map hasilnya ke proto response
	return &pb.UserResponse{
		Id:        uint32(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// UpdateUser implements pb.UserServiceServer
func (h *userHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	// Konversi request dari proto ke DTO
	dtoReq := &dto.UpdateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		IsActive: req.IsActive,
	}

	// Panggil service
	user, err := h.userService.Update(ctx, uint(req.Id), dtoReq)
	if err != nil {
		// Di Fiber kamu pakai utils.SendError, di gRPC cukup return nil, err
		return nil, err
	}

	// Mapping hasil service ke proto response
	return &pb.UserResponse{
		Id:        uint32(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// DeleteUser implements pb.UserServiceServer
func (h *userHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := h.userService.Delete(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}

// ListUsers implements pb.UserServiceServer
func (h *userHandler) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := h.userService.GetAll(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	// Mapping ke []*pb.UserResponse
	var pbUsers []*pb.UserResponse
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.UserResponse{
			Id:        uint32(u.ID),
			Name:      u.Name,
			Email:     u.Email,
			IsActive:  u.IsActive,
			CreatedAt: u.CreatedAt.Format(time.RFC3339),
			UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &pb.ListUsersResponse{
		Users: pbUsers,
		Total: int32(len(pbUsers)), // bisa diganti kalau ada total dari service
	}, nil
}
