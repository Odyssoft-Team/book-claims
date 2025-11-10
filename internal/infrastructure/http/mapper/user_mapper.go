package mapper

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/infrastructure/http/dto"
	"time"
)

func CreateUserDTOToDomain(c dto.CreateUserDTO) *model.User {
	return &model.User{
		TenantID:   c.TenantID,
		RoleID:     c.RoleID,
		LocationID: c.LocationID,
		Email:      c.Email,
		Password:   c.Password,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		FullName:   c.FullName,
		UserName:   c.UserName,
		Phone:      c.Phone,
	}
}

func UserToResponseDTO(user *model.User) dto.UserResponseDTO {
	return dto.UserResponseDTO{
		ID:         user.ID,
		TenantID:   user.TenantID,
		RoleID:     user.RoleID,
		RoleName:   user.RoleName,
		LocationID: user.LocationID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		FullName:   user.FullName,
		UserName:   user.UserName,
		Phone:      user.Phone,
		IsActive:   user.IsActive,
		CreatedAt:  user.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:  user.UpdatedAt.UTC().Format(time.RFC3339),
	}
}

func UpdateUserFromDTO(existing *model.User, dto dto.UpdateUserDTO) {
	existing.RoleID = *dto.RoleID
	existing.LocationID = *dto.LocationID
	existing.FirstName = *dto.FirstName
	existing.LastName = *dto.LastName
	existing.UserName = *dto.UserName
	existing.Password = *dto.Password
	existing.Email = *dto.Email
	existing.Phone = *dto.Phone
	existing.IsActive = *dto.IsActive
}
