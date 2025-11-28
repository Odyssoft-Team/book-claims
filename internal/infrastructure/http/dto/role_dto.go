package dto

import (
	"github.com/google/uuid"
)

// CreateRoleDTO representa los datos para crear un rol
// swagger:model CreateRoleDTO
type CreateRoleDTO struct {
	TenantID    uuid.UUID `json:"tenant_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	IsSystem    bool      `json:"is_system"`
}

// UpdateRoleDTO representa campos opcionales para actualizar rol
// swagger:model UpdateRoleDTO
type UpdateRoleDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	IsSystem    *bool   `json:"is_system"`
}

// RoleResponseDTO representa la respuesta de un rol
// swagger:model RoleResponseDTO
type RoleResponseDTO struct {
	ID          uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TenantID    uuid.UUID `json:"tenant_id"`
	IsSystem    bool      `json:"is_system"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
