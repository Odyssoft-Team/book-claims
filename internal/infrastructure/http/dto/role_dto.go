package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateRoleDTO struct {
	TenantID    uuid.UUID `json:"tenant_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	IsSystem    bool      `json:"is_system"`
}

// DTO de entrada para actualizar un usuario
type UpdateRoleDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	IsSystem    *bool   `json:"is_system"`
}

// DTO de salida para obtener un usuario
type RoleResponseDTO struct {
	ID          uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TenantID    uuid.UUID `json:"tenant_id"`
	IsSystem    bool      `json:"is_system"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
