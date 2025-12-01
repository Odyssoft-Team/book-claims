package mapper

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/infrastructure/http/dto"
	"time"
)

func CreateLocationDTOToDomain(c dto.CreateLocationDTO) *model.Location {
	return &model.Location{
		TenantID:   c.TenantID,
		Name:       c.Name,
		Address:    c.Address,
		Department: c.Department,
		Province:   c.Province,
		District:   c.District,
		PostalCode: c.PostalCode,
		Type:       model.EstablishmentType(c.Type),
		URL:        c.URL,
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
		Department: location.Department,
		Province:   location.Province,
		District:   location.District,
		PostalCode: location.PostalCode,
		Type:       string(location.Type),
		URL:        location.URL,
		IsActive:   location.IsActive,
		PublicCode: location.PublicCode,
		CreatedAt:  location.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:  location.UpdatedAt.UTC().Format(time.RFC3339),
	}
}

func UpdateLocationFromDTO(existing *model.Location, dto dto.UpdateLocationDTO) {
	if dto.Name != nil {
		existing.Name = *dto.Name
	}
	if dto.Address != nil {
		existing.Address = *dto.Address
	}
	if dto.Department != nil {
		existing.Department = *dto.Department
	}
	if dto.Province != nil {
		existing.Province = *dto.Province
	}
	if dto.District != nil {
		existing.District = *dto.District
	}
	if dto.PostalCode != nil {
		existing.PostalCode = *dto.PostalCode
	}
	if dto.Type != nil {
		existing.Type = model.EstablishmentType(*dto.Type)
	}
	if dto.URL != nil {
		existing.URL = *dto.URL
	}
	if dto.PublicCode != nil {
		existing.PublicCode = *dto.PublicCode
	}
	if dto.IsActive != nil {
		existing.IsActive = *dto.IsActive
	}
}
