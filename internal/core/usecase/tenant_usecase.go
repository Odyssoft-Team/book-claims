package usecase

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/core/port"
	"claimbook-api/internal/infrastructure/http/dto"
	"claimbook-api/internal/infrastructure/http/mapper"
	"claimbook-api/pkg/util/apperror"
	"context"
	"crypto/rand"
	"encoding/base64"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type TenantUseCase struct {
	tenantRepo port.TenantRepository
	roleRepo   port.RoleRepository
	userRepo   port.UserRepository
	apiKeyRepo port.ApiKeyRepository
}

func NewTenantUseCase(
	repo port.TenantRepository,
	role port.RoleRepository,
	user port.UserRepository,
	apiKey port.ApiKeyRepository,
) *TenantUseCase {
	return &TenantUseCase{
		tenantRepo: repo,
		roleRepo:   role,
		userRepo:   user,
		apiKeyRepo: apiKey,
	}
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

	oldIsActive := tenant.IsActive

	mapper.UpdateTenantFromDTO(tenant, *updateDTO)

	updated, err := uc.tenantRepo.UpdateTenant(ctx, tenant)
	if err != nil {
		return nil, apperror.NewInternalError("Failed to update tenant", err)
	}

	if updateDTO.IsActive != nil && *updateDTO.IsActive != oldIsActive {
		roles := []*model.Role{
			{Name: "SUPERADMIN", Description: "Super administrator", TenantID: id, IsSystem: true},
			{Name: "ADMIN", Description: "Administrator", TenantID: id, IsSystem: true},
			{Name: "SELLER", Description: "Seller user", TenantID: id, IsSystem: true},
			{Name: "PUBLIC", Description: "Public user", TenantID: id, IsSystem: false},
		}

		createdRoles, err := uc.roleRepo.CreateRoleBatchByTenant(ctx, roles)
		if err != nil {
			return nil, apperror.NewInternalError("Failed to create default roles", err)
		}

		var adminRoleID uuid.UUID
		for _, r := range createdRoles {
			if r.Name == "ADMIN" {
				adminRoleID = r.ID
				break
			}
		}
		if adminRoleID == uuid.Nil {
			return nil, apperror.NewInternalError("ADMIN role not found after creation", nil)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("defaultPassword123!"), bcrypt.DefaultCost)
		if err != nil {
			return nil, apperror.NewInternalError("failed to hash password", err)
		}

		adminUser := &model.User{
			TenantID:  id,
			RoleID:    adminRoleID,
			Email:     "admin@" + strings.ToLower(tenant.Name) + ".com",
			Password:  string(hashedPassword),
			FirstName: "Admin",
			LastName:  "User",
			FullName:  "Administrator" + " " + tenant.Name,
			UserName:  "admin_" + strings.ToLower(tenant.Name),
			Phone:     tenant.PhoneContact,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		_, err = uc.userRepo.CreateUser(ctx, adminUser)
		if err != nil {
			return nil, apperror.NewInternalError("Failed to create admin user", err)
		}

		key, err := generateApiKey()
		if err != nil {
			return nil, apperror.NewInternalError("Failed to generate API key", err)
		}

		apiKey := &model.ApiKey{
			TenantID:  id,
			ApiKey:    key,
			Scope:     "GLOBAL",
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		_, err = uc.apiKeyRepo.CreateApiKey(ctx, apiKey)
		if err != nil {
			return nil, apperror.NewInternalError("Failed to create apiKey", err)
		}
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

func generateApiKey() (string, error) {
	bytes := make([]byte, 32) // Tamaño estándar razonable
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	// Codificar en base64 URL-safe para que sea fácil de guardar y usar en URLs/headers
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
