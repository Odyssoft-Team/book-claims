package dto

import (
	"claimbook-api/internal/core/domain/model"
	"time"

	"github.com/google/uuid"
)

type CreateComplaintDTO struct {
	TenantID        uuid.UUID             `json:"tenant_id" binding:"required,uuid"`
	LocationID      uuid.UUID             `json:"location_id" binding:"required,uuid"`
	Type            model.ComplaintType   `json:"type" binding:"required,oneof=QUEJA RECLAMO"`
	Status          model.ComplaintStatus `json:"status" binding:"required,oneof=RECIBIDO EVALUACION PROCESO ATENDIDO CERRADO"`
	CategoryID      uuid.UUID             `json:"category_id" binding:"required,uuid"`
	SourceID        uuid.UUID             `json:"source_id" binding:"required,uuid"`
	ApiKeyID        uuid.UUID             `json:"api_key_id" binding:"required,uuid"`
	CodePublic      string                `json:"code_public"`
	Description     string                `json:"description" binding:"required"`
	RequestedAction string                `json:"requested_action"`
	IsClosed        bool                  `json:"is_closed"`
}

type ComplaintResponse struct {
	ID         uuid.UUID             `json:"id"`
	TenantID   uuid.UUID             `json:"tenant_id"`
	LocationID uuid.UUID             `json:"location_id"`
	Type       model.ComplaintType   `json:"type"`
	Status     model.ComplaintStatus `json:"status"`
	CategoryID uuid.UUID             `json:"category_id"`
	SourceID   uuid.UUID             `json:"source_id"`
	ApiKeyID   uuid.UUID             `json:"api_key_id"`

	CodePublic      string     `json:"code_public"`
	Description     string     `json:"description"`
	RequestedAction string     `json:"requested_action"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	ResolvedAt      *time.Time `json:"resolved_at"`
	IsClosed        bool       `json:"is_closed"`
}

type UpdateComplaintDTO struct {
	Status model.ComplaintStatus `json:"status" binding:"required,uuid"`
}
