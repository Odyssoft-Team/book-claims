package model

import (
	"time"

	"github.com/google/uuid"
)

type Complaint struct {
	ID              uuid.UUID
	TenantID        uuid.UUID
	LocationID      uuid.UUID
	TypeID          uuid.UUID
	StatusID        uuid.UUID
	CategoryID      uuid.UUID
	SourceID        uuid.UUID
	ApiKeyID        uuid.UUID
	CodePublic      string
	Description     string
	RequestedAction string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	ResolvedAt      *time.Time
	IsClosed        bool
}
