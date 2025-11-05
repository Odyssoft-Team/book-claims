package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateApiKeyDTO struct {
	TenantID   uuid.UUID `json:"tenant_id"  validate:"required"`
	LocationID uuid.UUID `json:"location_id"  validate:"required"`
	ApiKey     string    `json:"api_key"  validate:"required"`
	Scope      string    `json:"scope"  validate:"required"`
	IsActive   bool      `json:"is_active"`
}

// DTO de entrada para actualizar un usuario
type UpdateApiKeyDTO struct {
	IsActive *bool `json:"is_active"`
}

// DTO de salida para obtener un usuario
type ApiKeyResponseDTO struct {
	ID         uuid.UUID `json:"api_key_id"`
	TenantID   uuid.UUID `json:"tenant_id"`
	LocationID uuid.UUID `json:"location_id"`
	ApiKey     string    `json:"api_key"`
	Scope      string    `json:"scope"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
