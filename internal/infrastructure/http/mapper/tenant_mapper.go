package mapper

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/infrastructure/http/dto"
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
		CreatedAt:    tenant.CreatedAt,
		UpdatedAt:    tenant.UpdatedAt,
	}
}

func UpdateTenantFromDTO(existing *model.Tenant, dto dto.UpdateTenantDTO) {
	existing.Name = dto.Name
	existing.Ruc = dto.Ruc
	existing.EmailContact = dto.EmailContact
	existing.PhoneContact = dto.PhoneContact
	existing.IsActive = dto.IsActive
}
