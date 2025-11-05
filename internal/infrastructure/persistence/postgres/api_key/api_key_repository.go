package apikey

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/core/port"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type apiKeyPGRepository struct {
	db *gorm.DB
}

func NewApiKeyPGRepository(db *gorm.DB) port.ApiKeyRepository {
	return &apiKeyPGRepository{db: db}
}

func (r *apiKeyPGRepository) CreateApiKey(ctx context.Context, apiKey *model.ApiKey) (*model.ApiKey, error) {
	dbModel := ApiKeyModelFromDomain(apiKey)
	if err := r.db.WithContext(ctx).Create(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *apiKeyPGRepository) GetApiKeyById(ctx context.Context, id uuid.UUID) (*model.ApiKey, error) {
	var dbModel ApiKeyModel
	if err := r.db.WithContext(ctx).First(&dbModel, "id", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *apiKeyPGRepository) UpdateApiKey(ctx context.Context, apiKey *model.ApiKey) (*model.ApiKey, error) {
	dbModel := ApiKeyModelFromDomain(apiKey)
	if err := r.db.WithContext(ctx).Save(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *apiKeyPGRepository) GetApiKeys(ctx context.Context) ([]*model.ApiKey, error) {
	var dbModels []ApiKeyModel
	if err := r.db.WithContext(ctx).Find(&dbModels).Error; err != nil {
		return nil, err
	}

	var result []*model.ApiKey
	for _, m := range dbModels {
		result = append(result, m.ToDomain())
	}
	return result, nil
}
