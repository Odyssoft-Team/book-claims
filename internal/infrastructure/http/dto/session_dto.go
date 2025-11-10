package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateSessionDTO struct {
	UserID       uuid.UUID
	TenantID     uuid.UUID
	RefreshToken string
	IP           string
	UserAgent    string
	ExpiresAt    time.Time
}

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

type UpdateSessionDTO struct {
	Revoked *bool `json:"revoked"`
}
