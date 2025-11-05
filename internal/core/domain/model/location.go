package model

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	ID         uuid.UUID
	TenantID   uuid.UUID
	Name       string
	Address    string
	Type       string
	IsActive   bool
	PublicCode string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
