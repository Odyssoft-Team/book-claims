package port

import (
	"claimbook-api/internal/core/domain/model"
	"context"

	"github.com/google/uuid"
)

type ApiKeyRepository interface {
	CreateApiKey(ctx context.Context, apiKey *model.ApiKey) (*model.ApiKey, error)
	GetApiKeyById(ctx context.Context, id uuid.UUID) (*model.ApiKey, error)
	UpdateApiKey(ctx context.Context, user *model.ApiKey) (*model.ApiKey, error)
	GetApiKeys(ctx context.Context) ([]*model.ApiKey, error)
	IsValidApiKey(ctx context.Context, apiKey string) (bool, error)
}
