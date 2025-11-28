package dto

import (
	"github.com/google/uuid"
)

// CreateTenantDTO representa los datos necesarios para crear un tenant
// swagger:model CreateTenantDTO
type CreateTenantDTO struct {
	Name         string `json:"name" validate:"required"`
	Ruc          string `json:"ruc" validate:"required"`
	EmailContact string `json:"email_contact" validate:"required,email"`
	PhoneContact string `json:"phone_contact" validate:"required"`
	IsActive     bool   `json:"is_active"`
}

// UpdateTenantDTO representa los campos opcionales para actualizar un tenant
// swagger:model UpdateTenantDTO
type UpdateTenantDTO struct {
	Name         *string `json:"name" binding:"omitempty"`
	Ruc          *string `json:"ruc" binding:"omitempty"`
	EmailContact *string `json:"email_contact" binding:"omitempty"`
	PhoneContact *string `json:"phone_contact" binding:"omitempty"`
	IsActive     *bool   `json:"is_active" binding:"omitempty"`
	IsConfirm    *bool   `json:"is_confirm" binding:"omitempty"`
}

// TenantResponseDTO representa la respuesta de un tenant
// swagger:model TenantResponseDTO
type TenantResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Ruc          string    `json:"ruc"`
	EmailContact string    `json:"email_contact"`
	PhoneContact string    `json:"phone_contact"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
}

// Ejemplo de CreateTenantDTO
// {
//   "name": "Mi Empresa S.A.",
//   "ruc": "12345678901",
//   "email_contact": "contacto@miempresa.com",
//   "phone_contact": "+51999999999",
//   "is_active": false
// }
