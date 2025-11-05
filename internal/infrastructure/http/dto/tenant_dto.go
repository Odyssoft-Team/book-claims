package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateTenantDTO struct {
	Name         string `json:"name" validate:"required"`
	Ruc          string `json:"ruc" validate:"required"`
	EmailContact string `json:"email_contact" validate:"required,email"`
	PhoneContact string `json:"phone_contact" validate:"required"`
	IsActive     bool   `json:"is_active"`
}

type UpdateTenantDTO struct {
	Name         string `json:"name" validate:"required"`
	Ruc          string `json:"ruc" validate:"required"`
	EmailContact string `json:"email_contact" validate:"required,email"`
	PhoneContact string `json:"phone_contact" validate:"required"`
	IsActive     bool   `json:"is_active"`
}

type TenantResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Ruc          string    `json:"ruc"`
	EmailContact string    `json:"email_contact"`
	PhoneContact string    `json:"phone_contact"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
