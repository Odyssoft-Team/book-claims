package tenant

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/core/port"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tenantPGRepository struct {
	db *gorm.DB
}

func NewTenantPGRepository(db *gorm.DB) port.TenantRepository {
	return &tenantPGRepository{db: db}
}

func (r *tenantPGRepository) CreateTenant(ctx context.Context, tenant *model.Tenant) (*model.Tenant, error) {
	dbModel := TenantModelFromDomain(tenant)
	if err := r.db.WithContext(ctx).Create(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *tenantPGRepository) GetTenantById(ctx context.Context, id uuid.UUID) (*model.Tenant, error) {
	var dbModel TenantModel
	if err := r.db.WithContext(ctx).First(&dbModel, "id", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *tenantPGRepository) UpdateTenant(ctx context.Context, user *model.Tenant) (*model.Tenant, error) {
	dbModel := TenantModelFromDomain(user)
	if err := r.db.WithContext(ctx).Save(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *tenantPGRepository) GetTenants(ctx context.Context) ([]*model.Tenant, error) {
	var dbModels []TenantModel
	if err := r.db.WithContext(ctx).Find(&dbModels).Error; err != nil {
		return nil, err
	}

	var result []*model.Tenant
	for _, m := range dbModels {
		result = append(result, m.ToDomain())
	}
	return result, nil
}
