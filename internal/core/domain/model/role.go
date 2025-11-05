package model

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID
	Name        string
	Description string
	TenantID    uuid.UUID
	IsSystem    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
