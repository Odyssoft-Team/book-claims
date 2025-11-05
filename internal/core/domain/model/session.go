package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID
	TenantID     uuid.UUID
	UserID       uuid.UUID
	RefreshToken string
	IP           string
	UserAgent    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ExpiresAt    time.Time
	Revoked      bool
}
