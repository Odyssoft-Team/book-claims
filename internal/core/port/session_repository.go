package port

import (
	"claimbook-api/internal/core/domain/model"
	"context"
)

type SessionRepository interface {
	Create(ctx context.Context, session *model.Session) (*model.Session, error)
	FindByRefreshToken(ctx context.Context, token string) (*model.Session, error)
	Update(ctx context.Context, session *model.Session) error
}
