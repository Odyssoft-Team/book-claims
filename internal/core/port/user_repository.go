package port

import (
	"claimbook-api/internal/core/domain/model"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUsers(ctx context.Context) ([]*model.User, error)
	FindByUserAuth(ctx context.Context, validateNameOREmail string) (*model.User, error)
}
