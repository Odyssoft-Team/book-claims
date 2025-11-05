package usecase

import (
	"claimbook-api/internal/core/port"
	"claimbook-api/internal/infrastructure/http/dto"
	"claimbook-api/internal/infrastructure/http/mapper"
	"claimbook-api/pkg/util/apperror"
	"context"

	"github.com/google/uuid"
)

type RoleUseCase struct {
	roleRepo port.RoleRepository
}

func NewRoleUseCase(repo port.RoleRepository) *RoleUseCase {
	return &RoleUseCase{roleRepo: repo}
}

func (uc *RoleUseCase) CreateRole(ctx context.Context, roleDTO *dto.CreateRoleDTO) (*dto.RoleResponseDTO, error) {

	domainModel := mapper.CreateRoleDTOToDomain(*roleDTO)

	created, err := uc.roleRepo.CreateRole(ctx, domainModel)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to create role", err)
	}
	resp := mapper.RoleToResponseDTO(created)
	return &resp, nil
}

func (uc *RoleUseCase) GetRoleById(ctx context.Context, id uuid.UUID) (*dto.RoleResponseDTO, error) {
	role, err := uc.roleRepo.GetRoleById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve role", err)
	}
	if role == nil {
		return nil, apperror.NewNotFoundError("Role not found")
	}
	resp := mapper.RoleToResponseDTO(role)
	return &resp, nil
}

func (uc *RoleUseCase) UpdateRole(ctx context.Context, id uuid.UUID, updateDTO *dto.UpdateRoleDTO) (*dto.RoleResponseDTO, error) {
	role, err := uc.roleRepo.GetRoleById(ctx, id)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve role for update", err)
	}
	if role == nil {
		return nil, apperror.NewNotFoundError("Role not found")
	}

	mapper.UpdateRoleFromDTO(role, *updateDTO)

	updated, err := uc.roleRepo.UpdateRole(ctx, role)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to update role", err)
	}

	resp := mapper.RoleToResponseDTO(updated)
	return &resp, nil
}

func (uc *RoleUseCase) GetRoles(ctx context.Context) ([]dto.RoleResponseDTO, error) {
	roles, err := uc.roleRepo.GetRoles(ctx)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to retrieve roless by tenant ID", err)
	}

	if len(roles) == 0 {
		return nil, apperror.NewNotFoundError("No roless found for the tenant")
	}

	var responses []dto.RoleResponseDTO
	for _, role := range roles {
		responses = append(responses, mapper.RoleToResponseDTO(role))
	}

	return responses, nil
}
