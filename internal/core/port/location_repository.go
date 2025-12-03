package port

import (
	"claimbook-api/internal/core/domain/model"
	"context"

	"github.com/google/uuid"
)

type LocationRepository interface {
	CreateLocation(ctx context.Context, location *model.Location) (*model.Location, error)
	GetLocationById(ctx context.Context, id uuid.UUID) (*model.Location, error)
	UpdateLocation(ctx context.Context, location *model.Location) (*model.Location, error)
	GetLocations(ctx context.Context) ([]*model.Location, error)
	GetLocationsByTenant(ctx context.Context, tenantID uuid.UUID) ([]*model.Location, error)
}
