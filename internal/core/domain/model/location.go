package model

import (
	"time"

	"github.com/google/uuid"
)

// EstablishmentType representa el tipo de establecimiento de la ubicación
type EstablishmentType string

const (
	EstablishmentPhysical EstablishmentType = "FISICO"
	EstablishmentOnline   EstablishmentType = "ONLINE"
	EstablishmentBoth     EstablishmentType = "AMBOS"
)

type Location struct {
	ID         uuid.UUID
	TenantID   uuid.UUID
	Name       string
	Address    string
	Department string
	Province   string
	District   string
	PostalCode string
	Type       EstablishmentType
	URL        string
	IsActive   bool
	PublicCode string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
