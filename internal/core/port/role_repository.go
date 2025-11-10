package port

import (
	"claimbook-api/internal/core/domain/model"
	"context"

	"github.com/google/uuid"
)

type RoleRepository interface {
	CreateRole(ctx context.Context, role *model.Role) (*model.Role, error)
	GetRoleById(ctx context.Context, id uuid.UUID) (*model.Role, error)
	UpdateRole(ctx context.Context, role *model.Role) (*model.Role, error)
	GetRoles(ctx context.Context) ([]*model.Role, error)
	CreateRoleBatchByTenant(ctx context.Context, roles []*model.Role) ([]*model.Role, error)
}
