package mapper

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/infrastructure/http/dto"
	"time"
)

func CreateTenantDTOToDomain(c dto.CreateTenantDTO) *model.Tenant {
	return &model.Tenant{
		Name:         c.Name,
		Ruc:          c.Ruc,
		EmailContact: c.EmailContact,
		PhoneContact: c.PhoneContact,
		IsActive:     c.IsActive,
	}
}

func TenantToResponseDTO(tenant *model.Tenant) dto.TenantResponseDTO {
	return dto.TenantResponseDTO{
		ID:           tenant.ID,
		Name:         tenant.Name,
		Ruc:          tenant.Ruc,
		EmailContact: tenant.EmailContact,
		PhoneContact: tenant.PhoneContact,
		IsActive:     tenant.IsActive,
		CreatedAt:    tenant.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:    tenant.UpdatedAt.UTC().Format(time.RFC3339),
	}
}

func UpdateTenantFromDTO(existing *model.Tenant, dto dto.UpdateTenantDTO) {
	if dto.Name != nil {
		existing.Name = *dto.Name
	}

	if dto.Ruc != nil {
		existing.Ruc = *dto.Ruc
	}

	if dto.EmailContact != nil {
		existing.EmailContact = *dto.EmailContact
	}

	if dto.PhoneContact != nil {
		existing.PhoneContact = *dto.PhoneContact
	}

	if dto.IsActive != nil {
		existing.IsActive = *dto.IsActive
	}

	if dto.IsConfirm != nil {
		existing.IsConfirm = *dto.IsConfirm
	}

}
