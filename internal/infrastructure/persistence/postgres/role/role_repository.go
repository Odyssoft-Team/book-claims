package role

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/core/port"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type rolePGRepository struct {
	db *gorm.DB
}

func NewRolePGRepository(db *gorm.DB) port.RoleRepository {
	return &rolePGRepository{db: db}
}

func (r *rolePGRepository) CreateRole(ctx context.Context, user *model.Role) (*model.Role, error) {
	dbModel := RoleModelFromDomain(user)
	if err := r.db.WithContext(ctx).Create(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *rolePGRepository) GetRoleById(ctx context.Context, id uuid.UUID) (*model.Role, error) {
	var dbModel RoleModel
	if err := r.db.WithContext(ctx).First(&dbModel, "id", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *rolePGRepository) UpdateRole(ctx context.Context, user *model.Role) (*model.Role, error) {
	dbModel := RoleModelFromDomain(user)
	if err := r.db.WithContext(ctx).Save(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *rolePGRepository) GetRoles(ctx context.Context) ([]*model.Role, error) {
	var dbModels []RoleModel
	if err := r.db.WithContext(ctx).Find(&dbModels).Error; err != nil {
		return nil, err
	}

	var result []*model.Role
	for _, m := range dbModels {
		result = append(result, m.ToDomain())
	}
	return result, nil
}
