package usecase

import (
	"claimbook-api/internal/core/port"
	"claimbook-api/internal/infrastructure/http/dto"
	"claimbook-api/internal/infrastructure/http/mapper"
	"claimbook-api/pkg/util/apperror"
	"context"

	"github.com/google/uuid"
)

type TenantUseCase struct {
	tenantRepo port.TenantRepository
}

func NewTenantUseCase(repo port.TenantRepository) *TenantUseCase {
	return &TenantUseCase{tenantRepo: repo}
}

func (uc *TenantUseCase) CreateTenant(ctx context.Context, tenantDTO *dto.CreateTenantDTO) (*dto.TenantResponseDTO, error) {

	domainModel := mapper.CreateTenantDTOToDomain(*tenantDTO)

	created, err := uc.tenantRepo.CreateTenant(ctx, domainModel)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to create tenant", err)
	}
	resp := mapper.TenantToResponseDTO(created)
	return &resp, nil
}

func (uc *TenantUseCase) GetTenantById(ctx context.Context, id uuid.UUID) (*dto.TenantResponseDTO, error) {
	tenant, err := uc.tenantRepo.GetTenantById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve tenant", err)
	}
	if tenant == nil {
		return nil, apperror.NewNotFoundError("Tenant not found")
	}
	resp := mapper.TenantToResponseDTO(tenant)
	return &resp, nil
}

func (uc *TenantUseCase) UpdateTenant(ctx context.Context, id uuid.UUID, updateDTO *dto.UpdateTenantDTO) (*dto.TenantResponseDTO, error) {
	tenant, err := uc.tenantRepo.GetTenantById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve tenant for update", err)
	}
	if tenant == nil {
		return nil, apperror.NewNotFoundError("Tenant not found")
	}

	mapper.UpdateTenantFromDTO(tenant, *updateDTO)

	updated, err := uc.tenantRepo.UpdateTenant(ctx, tenant)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to update tenant", err)
	}

	resp := mapper.TenantToResponseDTO(updated)
	return &resp, nil
}

func (uc *TenantUseCase) GetTenants(ctx context.Context) ([]dto.TenantResponseDTO, error) {
	tenants, err := uc.tenantRepo.GetTenants(ctx)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve tenants", err)
	}

	if len(tenants) == 0 {
		return nil, apperror.NewNotFoundError("No tenants found for the tenant")
	}

	var responses []dto.TenantResponseDTO
	for _, tenant := range tenants {
		responses = append(responses, mapper.TenantToResponseDTO(tenant))
	}

	return responses, nil
}
