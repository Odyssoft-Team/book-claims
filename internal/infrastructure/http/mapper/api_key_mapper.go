package mapper

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/infrastructure/http/dto"
	"time"
)

func CreateApiKeyDTOToDomain(c dto.CreateApiKeyDTO) *model.ApiKey {
	return &model.ApiKey{
		TenantID:   c.TenantID,
		LocationID: c.LocationID,
		ApiKey:     c.ApiKey,
		Scope:      c.Scope,
		IsActive:   c.IsActive,
	}
}

func ApiKeyToResponseDTO(apiKey *model.ApiKey) dto.ApiKeyResponseDTO {
	return dto.ApiKeyResponseDTO{
		ID:         apiKey.ID,
		TenantID:   apiKey.TenantID,
		LocationID: apiKey.LocationID,
		ApiKey:     apiKey.ApiKey,
		Scope:      apiKey.Scope,
		IsActive:   apiKey.IsActive,
		CreatedAt:  apiKey.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:  apiKey.UpdatedAt.UTC().Format(time.RFC3339),
	}
}

func UpdateApiKeyFromDTO(existing *model.ApiKey, dto dto.UpdateApiKeyDTO) {
	existing.IsActive = *dto.IsActive
}
