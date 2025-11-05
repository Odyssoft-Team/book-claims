package port

import (
	"claimbook-api/internal/core/domain/model"
	"context"

	"github.com/google/uuid"
)

type TenantRepository interface {
	CreateTenant(ctx context.Context, role *model.Tenant) (*model.Tenant, error)
	GetTenantById(ctx context.Context, id uuid.UUID) (*model.Tenant, error)
	UpdateTenant(ctx context.Context, role *model.Tenant) (*model.Tenant, error)
	GetTenants(ctx context.Context) ([]*model.Tenant, error)
}
