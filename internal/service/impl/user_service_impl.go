package serviceimpl

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	dto "github.com/faizalnurrozi/go-starter-kit/internal/dto/request"
	response "github.com/faizalnurrozi/go-starter-kit/internal/dto/response"
	"github.com/faizalnurrozi/go-starter-kit/internal/entity"
	"github.com/faizalnurrozi/go-starter-kit/internal/errors"
	"github.com/faizalnurrozi/go-starter-kit/internal/logger"
	"github.com/faizalnurrozi/go-starter-kit/internal/repository/interfaces"
	iUc "github.com/faizalnurrozi/go-starter-kit/internal/service/interfaces"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	userRepo interfaces.UserRepository
	redis    *redis.Client
}

func NewUserService(userRepo interfaces.UserRepository, redis *redis.Client) iUc.UserService {
	return &userService{
		userRepo: userRepo,
		redis:    redis,
	}
}

func (s *userService) Create(ctx context.Context, req *dto.CreateUserRequest) (*response.UserResponse, error) {
	logger.WithFields(logrus.Fields{
		"email":  req.Email,
		"action": "create_user",
	}).Info("Creating new user")

	// Check if user already exists
	_, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err == nil {
		return nil, errors.NewBusinessError("Email already exists")
	}
	if err != gorm.ErrRecordNotFound {
		logger.Error("Error checking existing user: ", err)
		return nil, errors.NewInternalError("Failed to check existing user")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Error hashing password: ", err)
		return nil, errors.NewInternalError("Failed to hash password")
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		IsActive: true,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		logger.Error("Error creating user: ", err)
		return nil, errors.NewInternalError("Failed to create user")
	}

	// Cache user
	s.cacheUser(ctx, user)

	logger.WithFields(logrus.Fields{
		"user_id": user.ID,
		"email":   user.Email,
	}).Info("User created successfully")

	return response.NewUserResponse(user), nil
}

func (s *userService) GetByID(ctx context.Context, id uint) (*response.UserResponse, error) {
	// Try to get from cache first
	if user := s.getCachedUser(ctx, id); user != nil {
		return response.NewUserResponse(user), nil
	}

	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("User")
		}
		logger.Error("Error getting user: ", err)
		return nil, errors.NewInternalError("Failed to get user")
	}

	// Cache user
	s.cacheUser(ctx, user)

	return response.NewUserResponse(user), nil
}

func (s *userService) GetAll(ctx context.Context, limit, offset int) ([]*response.UserResponse, error) {
	users, err := s.userRepo.GetAll(ctx, limit, offset)
	if err != nil {
		logger.Error("Error getting users: ", err)
		return nil, errors.NewInternalError("Failed to get users")
	}

	return response.NewUserListResponse(users), nil
}

func (s *userService) Update(ctx context.Context, id uint, req *dto.UpdateUserRequest) (*response.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("User")
		}
		logger.Error("Error getting user: ", err)
		return nil, errors.NewInternalError("Failed to get user")
	}

	// Update fields if provided
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		logger.Error("Error updating user: ", err)
		return nil, errors.NewInternalError("Failed to update user")
	}

	// Invalidate cache
	s.invalidateUserCache(ctx, id)

	return response.NewUserResponse(user), nil
}

func (s *userService) Delete(ctx context.Context, id uint) error {
	_, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.NewNotFoundError("User")
		}
		logger.Error("Error getting user: ", err)
		return errors.NewInternalError("Failed to get user")
	}

	if err := s.userRepo.Delete(ctx, id); err != nil {
		logger.Error("Error deleting user: ", err)
		return errors.NewInternalError("Failed to delete user")
	}

	// Invalidate cache
	s.invalidateUserCache(ctx, id)

	return nil
}

func (s *userService) cacheUser(ctx context.Context, user *entity.User) {
	if s.redis == nil {
		return
	}

	key := fmt.Sprintf("user:%d", user.ID)
	data, _ := json.Marshal(user)
	s.redis.Set(ctx, key, data, 15*time.Minute)
}

func (s *userService) getCachedUser(ctx context.Context, id uint) *entity.User {
	if s.redis == nil {
		return nil
	}

	key := fmt.Sprintf("user:%d", id)
	data, err := s.redis.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	var user entity.User
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil
	}

	return &user
}

func (s *userService) invalidateUserCache(ctx context.Context, id uint) {
	if s.redis == nil {
		return
	}

	key := fmt.Sprintf("user:%d", id)
	s.redis.Del(ctx, key)
}
