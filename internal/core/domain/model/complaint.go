package model

import (
	"time"

	"github.com/google/uuid"
)

type ComplaintType string

const (
	RECLAMO ComplaintType = "RECLAMO"
	QUEJA   ComplaintType = "QUEJA"
)

type ComplaintStatus string

const (
	RECIBIDO   ComplaintStatus = "RECIBIDO"
	EVALUACION ComplaintStatus = "EN EVALUACION"
	PROCESO    ComplaintStatus = "EN PROCESO"
	ATENDIDO   ComplaintStatus = "ATENDIDO"
	CERRADO    ComplaintStatus = "CERRADO"
)

type ComplaintSource string

const (
	WEB     ComplaintSource = "WEB PUBLICA"
	OFICINA ComplaintSource = "OFICINA FISICA"
	API     ComplaintSource = "API"
	CALL    ComplaintSource = "CALL CENTER"
)

type Complaint struct {
	ID              uuid.UUID
	TenantID        uuid.UUID
	LocationID      uuid.UUID
	Type            ComplaintType
	Status          ComplaintStatus
	CategoryID      uuid.UUID
	Source          ComplaintSource
	ApiKeyID        uuid.UUID
	CodePublic      string
	Description     string
	RequestedAction string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	ResolvedAt      *time.Time
	IsClosed        bool
}
