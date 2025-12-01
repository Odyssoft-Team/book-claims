package dto

import (
	"github.com/google/uuid"
)

type CreateLocationDTO struct {
	TenantID   uuid.UUID `json:"tenant_id" validate:"required"`
	Name       string    `json:"name" validate:"required,min=2,max=100"`
	Address    string    `json:"address" validate:"required,min=5,max=200"`
	Department string    `json:"department" validate:"required"`
	Province   string    `json:"province" validate:"required"`
	District   string    `json:"district" validate:"required"`
	PostalCode string    `json:"postal_code" binding:"omitempty"`
	Type       string    `json:"type" validate:"required,oneof=FISICO ONLINE AMBOS"`
	URL        string    `json:"url" binding:"omitempty,url"`
	IsActive   bool      `json:"is_active"`
	PublicCode string    `json:"public_code" validate:"required,min=2,max=50"`
}

type UpdateLocationDTO struct {
	Name       *string `json:"name"`
	Address    *string `json:"address"`
	Department *string `json:"department"`
	Province   *string `json:"province"`
	District   *string `json:"district"`
	PostalCode *string `json:"postal_code"`
	Type       *string `json:"type" binding:"omitempty,oneof=FISICO ONLINE AMBOS"`
	URL        *string `json:"url" binding:"omitempty,url"`
	IsActive   *bool   `json:"is_active"`
	PublicCode *string `json:"public_code"`
}

type LocationResponseDTO struct {
	ID         uuid.UUID `json:"user_id"`
	TenantID   uuid.UUID `json:"tenant_id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Department string    `json:"department"`
	Province   string    `json:"province"`
	District   string    `json:"district"`
	PostalCode string    `json:"postal_code"`
	Type       string    `json:"type"`
	URL        string    `json:"url"`
	IsActive   bool      `json:"is_active"`
	PublicCode string    `json:"public_code"`
	CreatedAt  string    `json:"created_at"`
	UpdatedAt  string    `json:"updated_at"`
}
