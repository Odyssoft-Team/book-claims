package dto

import (
	"time"

	"github.com/google/uuid"
)

// CreateSessionDTO representa la sesión a crear
// swagger:model CreateSessionDTO
type CreateSessionDTO struct {
	UserID       uuid.UUID
	TenantID     uuid.UUID
	RefreshToken string
	IP           string
	UserAgent    string
	ExpiresAt    time.Time
}

// ResponseSessionDTO representa la respuesta de una sesión
// swagger:model ResponseSessionDTO
type ResponseSessionDTO struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	TenantID     uuid.UUID `json:"tenant_id"`
	RefreshToken string    `json:"refresh_token"`
	IP           string    `json:"ip"`
	UserAgent    string    `json:"user_agent"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
	ExpiresAt    string    `json:"expires_at"`
	Revoked      bool      `json:"revoked"`
}

// UpdateSessionDTO representa campos para actualizar sesión
// swagger:model UpdateSessionDTO
type UpdateSessionDTO struct {
	Revoked *bool `json:"revoked"`
}
