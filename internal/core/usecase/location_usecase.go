package usecase

import (
	"claimbook-api/internal/core/port"
	"claimbook-api/internal/infrastructure/http/dto"
	"claimbook-api/internal/infrastructure/http/mapper"
	"claimbook-api/pkg/util/apperror"
	"context"

	"github.com/google/uuid"
)

type LocationUseCase struct {
	locationRepo port.LocationRepository
}

func NewLocationUseCase(repo port.LocationRepository) *LocationUseCase {
	return &LocationUseCase{locationRepo: repo}
}

func (uc *LocationUseCase) CreateLocation(ctx context.Context, locationDTO *dto.CreateLocationDTO) (*dto.LocationResponseDTO, error) {

	domainModel := mapper.CreateLocationDTOToDomain(*locationDTO)

	created, err := uc.locationRepo.CreateLocation(ctx, domainModel)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to create location", err)
	}
	resp := mapper.LocationToResponseDTO(created)
	return &resp, nil
}

func (uc *LocationUseCase) GetLocationById(ctx context.Context, id uuid.UUID) (*dto.LocationResponseDTO, error) {
	location, err := uc.locationRepo.GetLocationById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve location", err)
	}
	if location == nil {
		return nil, apperror.NewNotFoundError("Location not found")
	}
	resp := mapper.LocationToResponseDTO(location)
	return &resp, nil
}

func (uc *LocationUseCase) UpdateLocation(ctx context.Context, id uuid.UUID, updateDTO *dto.UpdateLocationDTO) (*dto.LocationResponseDTO, error) {
	location, err := uc.locationRepo.GetLocationById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve location for update", err)
	}
	if location == nil {
		return nil, apperror.NewNotFoundError("Location not found")
	}

	mapper.UpdateLocationFromDTO(location, *updateDTO)

	updated, err := uc.locationRepo.UpdateLocation(ctx, location)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to update location", err)
	}

	resp := mapper.LocationToResponseDTO(updated)
	return &resp, nil
}

func (uc *LocationUseCase) GetLocations(ctx context.Context) ([]dto.LocationResponseDTO, error) {
	locations, err := uc.locationRepo.GetLocations(ctx)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve locations by tenant ID", err)
	}

	if len(locations) == 0 {
		return nil, apperror.NewNotFoundError("No locations found for the tenant")
	}

	var responses []dto.LocationResponseDTO
	for _, location := range locations {
		responses = append(responses, mapper.LocationToResponseDTO(location))
	}

	return responses, nil
}

func (uc *LocationUseCase) GetLocationsByTenant(ctx context.Context, tenantID uuid.UUID) ([]dto.LocationResponseDTO, error) {
	locations, err := uc.locationRepo.GetLocationsByTenant(ctx, tenantID)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve locations", err)
	}
	if len(locations) == 0 {
		return nil, apperror.NewNotFoundError("No locations found for the tenant")
	}
	var resp []dto.LocationResponseDTO
	for _, l := range locations {
		resp = append(resp, mapper.LocationToResponseDTO(l))
	}
	return resp, nil
}
