package dto

import (
	"github.com/google/uuid"
)

// CreateUserDTO representa los datos para crear un usuario
// swagger:model CreateUserDTO
type CreateUserDTO struct {
	TenantID   uuid.UUID `json:"tenant_id" validate:"required"`
	RoleID     uuid.UUID `json:"role_id" validate:"required"`
	LocationID uuid.UUID `json:"location_id" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required,min=8,max=50"`
	FirstName  string    `json:"first_name" validate:"required,min=2,max=100"`
	LastName   string    `json:"last_name" validate:"required,min=2,max=100"`
	FullName   string    `json:"full_name" validate:"required,min=2,max=200"`
	UserName   string    `json:"user_name" validate:"required,min=5,max=50"`
	Phone      string    `json:"phone" validate:"omitempty,min=7,max=20"`
	IsActive   bool      `json:"is_active"`
}

// UpdateUserDTO representa campos opcionales para actualizar usuario
// swagger:model UpdateUserDTO
type UpdateUserDTO struct {
	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	UserName   *string `json:"username"`
	Password   *string `json:"password"`
	Email      *string `json:"email"`
	Phone      *string `json:"phone"`
	IsActive   *bool   `json:"is_active"`
	LocationID *uuid.UUID
	RoleID     *uuid.UUID
}

// UserResponseDTO representa la respuesta del usuario
// swagger:model UserResponseDTO
type UserResponseDTO struct {
	ID         uuid.UUID `json:"user_id"`
	TenantID   uuid.UUID `json:"tenant_id"`
	RoleID     uuid.UUID `json:"role_id"`
	RoleName   string    `json:"role_name"`
	LocationID uuid.UUID `json:"location_id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	FullName   string    `json:"full_name"`
	UserName   string    `json:"username"`
	Phone      string    `json:"phone"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  string    `json:"created_at"`
	UpdatedAt  string    `json:"updated_at"`
}
