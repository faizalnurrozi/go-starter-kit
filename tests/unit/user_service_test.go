package unit

import (
	"context"
	"testing"

	dto "github.com/faizalnurrozi/go-starter-kit/internal/dto/request"
	"github.com/faizalnurrozi/go-starter-kit/internal/entity"
	serviceimpl "github.com/faizalnurrozi/go-starter-kit/internal/service/impl"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id uint) (*entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) GetAll(ctx context.Context, limit, offset int) ([]entity.User, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) Count(ctx context.Context) (int64, error) {
	args := m.Called(ctx)
	return args.Get(0).(int64), args.Error(1)
}

func TestUserService_Create_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := serviceimpl.NewUserService(mockRepo, nil)

	ctx := context.Background()
	req := &dto.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	// Mock that email doesn't exist
	mockRepo.On("GetByEmail", ctx, req.Email).Return(nil, gorm.ErrRecordNotFound)

	// Mock successful creation
	mockRepo.On("Create", ctx, mock.AnythingOfType("*entity.User")).Return(nil)

	result, err := userService.Create(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Name, result.Name)
	assert.Equal(t, req.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Create_EmailExists(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := serviceimpl.NewUserService(mockRepo, nil)

	ctx := context.Background()
	req := &dto.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	existingUser := &entity.User{
		ID:    1,
		Email: req.Email,
	}

	// Mock that email exists
	mockRepo.On("GetByEmail", ctx, req.Email).Return(existingUser, nil)

	result, err := userService.Create(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Email already exists")
	mockRepo.AssertExpectations(t)
}
