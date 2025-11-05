package user

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/core/port"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userPGRepository struct {
	db *gorm.DB
}

func NewUserPGRepository(db *gorm.DB) port.UserRepository {
	return &userPGRepository{db: db}
}

func (r *userPGRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	dbModel := UserModelFromDomain(user)
	if err := r.db.WithContext(ctx).Create(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *userPGRepository) GetUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var dbModel UserModel
	if err := r.db.WithContext(ctx).First(&dbModel, "id", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *userPGRepository) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	dbModel := UserModelFromDomain(user)
	if err := r.db.WithContext(ctx).Save(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *userPGRepository) GetUsers(ctx context.Context) ([]*model.User, error) {
	var dbModels []UserModel
	if err := r.db.WithContext(ctx).Find(&dbModels).Error; err != nil {
		return nil, err
	}

	var result []*model.User
	for _, m := range dbModels {
		result = append(result, m.ToDomain())
	}
	return result, nil
}

func (r *userPGRepository) FindByUserAuth(ctx context.Context, identifier string) (*model.User, error) {
	var user UserModel

	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("user_name = ? OR email = ?", identifier, identifier).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	domainUser := user.ToDomain()
	domainUser.RoleName = user.Role.Name

	return domainUser, nil
}
