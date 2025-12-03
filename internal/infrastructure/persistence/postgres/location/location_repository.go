package location

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/core/port"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type locationPGRepository struct {
	db *gorm.DB
}

func NewLocationPGRepository(db *gorm.DB) port.LocationRepository {
	return &locationPGRepository{db: db}
}

func (r *locationPGRepository) CreateLocation(ctx context.Context, user *model.Location) (*model.Location, error) {
	dbModel := LocationModelFromDomain(user)
	if err := r.db.WithContext(ctx).Create(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *locationPGRepository) GetLocationById(ctx context.Context, id uuid.UUID) (*model.Location, error) {
	var dbModel LocationModel
	if err := r.db.WithContext(ctx).First(&dbModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *locationPGRepository) UpdateLocation(ctx context.Context, location *model.Location) (*model.Location, error) {
	dbModel := LocationModelFromDomain(location)
	if err := r.db.WithContext(ctx).Save(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *locationPGRepository) GetLocations(ctx context.Context) ([]*model.Location, error) {
	var dbModels []LocationModel
	if err := r.db.WithContext(ctx).Find(&dbModels).Error; err != nil {
		return nil, err
	}

	var result []*model.Location
	for _, m := range dbModels {
		result = append(result, m.ToDomain())
	}
	return result, nil
}

func (r *locationPGRepository) GetLocationsByTenant(ctx context.Context, tenantID uuid.UUID) ([]*model.Location, error) {
	var dbModels []LocationModel
	if err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Find(&dbModels).Error; err != nil {
		return nil, err
	}
	var result []*model.Location
	for _, m := range dbModels {
		result = append(result, m.ToDomain())
	}
	return result, nil
}
