package usecase

import (
	"claimbook-api/internal/core/port"
	"claimbook-api/internal/infrastructure/http/dto"
	"claimbook-api/internal/infrastructure/http/mapper"
	"claimbook-api/pkg/util/apperror"
	"context"

	"github.com/google/uuid"
)

type ApiKeyUseCase struct {
	apiKeyRepo port.ApiKeyRepository
}

func NewApiKeyUseCase(repo port.ApiKeyRepository) *ApiKeyUseCase {
	return &ApiKeyUseCase{apiKeyRepo: repo}
}

func (uc *ApiKeyUseCase) CreateApiKey(ctx context.Context, apiKeyDTO *dto.CreateApiKeyDTO) (*dto.ApiKeyResponseDTO, error) {
	domainModel := mapper.CreateApiKeyDTOToDomain(*apiKeyDTO)

	created, err := uc.apiKeyRepo.CreateApiKey(ctx, domainModel)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to create apiKey", err)
	}
	resp := mapper.ApiKeyToResponseDTO(created)
	return &resp, nil
}

func (uc *ApiKeyUseCase) GetApiKeyById(ctx context.Context, id uuid.UUID) (*dto.ApiKeyResponseDTO, error) {
	apiKey, err := uc.apiKeyRepo.GetApiKeyById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve apiKey", err)
	}
	if apiKey == nil {
		return nil, apperror.NewNotFoundError("ApiKey not found")
	}
	resp := mapper.ApiKeyToResponseDTO(apiKey)
	return &resp, nil
}

func (uc *ApiKeyUseCase) UpdateApiKey(ctx context.Context, id uuid.UUID, updateDTO *dto.UpdateApiKeyDTO) (*dto.ApiKeyResponseDTO, error) {
	apiKey, err := uc.apiKeyRepo.GetApiKeyById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve apiKey for update", err)
	}
	if apiKey == nil {
		return nil, apperror.NewNotFoundError("ApiKey not found")
	}

	mapper.UpdateApiKeyFromDTO(apiKey, *updateDTO)

	updated, err := uc.apiKeyRepo.UpdateApiKey(ctx, apiKey)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to update apiKey", err)
	}

	resp := mapper.ApiKeyToResponseDTO(updated)
	return &resp, nil
}

func (uc *ApiKeyUseCase) GetApiKeys(ctx context.Context) ([]dto.ApiKeyResponseDTO, error) {
	apiKeys, err := uc.apiKeyRepo.GetApiKeys(ctx)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve apiKeys by tenant ID", err)
	}

	if len(apiKeys) == 0 {
		return nil, apperror.NewNotFoundError("No apiKeys found for the tenant")
	}

	var responses []dto.ApiKeyResponseDTO
	for _, apiKey := range apiKeys {
		responses = append(responses, mapper.ApiKeyToResponseDTO(apiKey))
	}

	return responses, nil
}
