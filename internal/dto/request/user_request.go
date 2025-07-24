package dto

type CreateUserRequest struct {
    Name     string `json:"name" validate:"required,min=2,max=100"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
    Name     *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
    Email    *string `json:"email,omitempty" validate:"omitempty,email"`
    IsActive *bool   `json:"is_active,omitempty"`
}

type GetUserParams struct {
    ID uint `params:"id" validate:"required,min=1"`
}