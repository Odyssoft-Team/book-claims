package dto

import (
	"github.com/google/uuid"
)

type CreateComplaintDTO struct {
	TenantID        uuid.UUID `json:"tenant_id" binding:"required,uuid"`
	LocationID      uuid.UUID `json:"location_id" binding:"required,uuid"`
	Type            string    `json:"type" binding:"required,oneof=QUEJA RECLAMO"`
	Status          string    `json:"status" binding:"required,oneof=RECIBIDO EVALUACION PROCESO ATENDIDO CERRADO"`
	CategoryID      uuid.UUID `json:"category_id" binding:"required,uuid"`
	Source          string    `json:"source" binding:"required,oneof=WEB OFICINA API CALL"`
	ApiKeyID        uuid.UUID `json:"api_key_id"`
	CodePublic      string    `json:"code_public"`
	Description     string    `json:"description" binding:"required"`
	RequestedAction string    `json:"requested_action"`
	IsClosed        bool      `json:"is_closed"`
}

type ComplaintResponse struct {
	ID         uuid.UUID `json:"id"`
	TenantID   uuid.UUID `json:"tenant_id"`
	LocationID uuid.UUID `json:"location_id"`
	Type       string    `json:"type"`
	Status     string    `json:"status"`
	CategoryID uuid.UUID `json:"category_id"`
	Source     string    `json:"source"`
	ApiKeyID   uuid.UUID `json:"api_key_id"`

	CodePublic      string `json:"code_public"`
	Description     string `json:"description"`
	RequestedAction string `json:"requested_action"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	ResolvedAt      string `json:"resolved_at"`
	IsClosed        bool   `json:"is_closed"`

	// Respuesta
	ResponseText   string `json:"response_text"`
	ResponseStatus string `json:"response_status"`
	ResponderID    string `json:"responder_id"`
	ResponseSentAt string `json:"response_sent_at"`
}

type UpdateComplaintDTO struct {
	Status string `json:"status" binding:"required,oneof='RECIBIDO' 'EN EVALUACION' 'PROCESO' 'ATENDIDO' 'CERRADO'"`

	// Respuesta del sistema
	ResponseText   *string    `json:"response_text"`
	ResponseStatus *string    `json:"response_status"` // DRAFT o SENT
	ResponderID    *uuid.UUID `json:"responder_id"`
	// new_status permite cambiar status sin respuesta
	NewStatus *string `json:"new_status"`
}
