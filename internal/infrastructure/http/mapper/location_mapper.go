package mapper

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/infrastructure/http/dto"
)

func CreateLocationDTOToDomain(c dto.CreateLocationDTO) *model.Location {
	return &model.Location{
		TenantID:   c.TenantID,
		Name:       c.Name,
		Address:    c.Address,
		Type:       c.Type,
		IsActive:   c.IsActive,
		PublicCode: c.PublicCode,
	}
}

func LocationToResponseDTO(location *model.Location) dto.LocationResponseDTO {
	return dto.LocationResponseDTO{
		ID:         location.ID,
		TenantID:   location.TenantID,
		Name:       location.Name,
		Address:    location.Address,
		Type:       location.Type,
		IsActive:   location.IsActive,
		PublicCode: location.PublicCode,
		CreatedAt:  location.CreatedAt,
		UpdatedAt:  location.UpdatedAt,
	}
}

func UpdateLocationFromDTO(existing *model.Location, dto dto.UpdateLocationDTO) {

	existing.Name = *dto.Name
	existing.Address = *dto.Address
	existing.Type = *dto.Type
	existing.PublicCode = *dto.PublicCode
	existing.IsActive = *dto.IsActive
}
