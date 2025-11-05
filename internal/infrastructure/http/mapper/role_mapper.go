package mapper

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/infrastructure/http/dto"
)

func CreateRoleDTOToDomain(c dto.CreateRoleDTO) *model.Role {
	return &model.Role{
		TenantID:    c.TenantID,
		Name:        c.Name,
		Description: c.Description,
		IsSystem:    c.IsSystem,
	}
}

func RoleToResponseDTO(role *model.Role) dto.RoleResponseDTO {
	return dto.RoleResponseDTO{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		TenantID:    role.TenantID,
		IsSystem:    role.IsSystem,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}

func UpdateRoleFromDTO(existing *model.Role, dto dto.UpdateRoleDTO) {
	existing.Name = *dto.Name
	existing.Description = *dto.Description
	existing.IsSystem = *dto.IsSystem
}
