package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	TenantID   uuid.UUID
	RoleID     uuid.UUID
	RoleName   string
	LocationID uuid.UUID
	Email      string
	Password   string
	FirstName  string
	LastName   string
	FullName   string
	UserName   string
	Phone      string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
