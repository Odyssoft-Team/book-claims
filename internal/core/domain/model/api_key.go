package model

import (
	"time"

	"github.com/google/uuid"
)

type ApiKey struct {
	ID         uuid.UUID
	TenantID   uuid.UUID
	LocationID uuid.UUID
	ApiKey     string
	Scope      string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
