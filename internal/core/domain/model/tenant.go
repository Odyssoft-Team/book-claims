package model

import (
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID           uuid.UUID
	Name         string
	Ruc          string
	EmailContact string
	PhoneContact string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
