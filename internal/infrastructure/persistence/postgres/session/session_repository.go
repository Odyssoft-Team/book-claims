package session

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/core/port"
	"context"

	"gorm.io/gorm"
)

type sessionPGRepository struct {
	db *gorm.DB
}

func NewSessionPGRepository(db *gorm.DB) port.SessionRepository {
	return &sessionPGRepository{db: db}
}

func (r *sessionPGRepository) Create(ctx context.Context, session *model.Session) (*model.Session, error) {
	dbModel := SessionModelFromDomain(session)
	if err := r.db.WithContext(ctx).Create(dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *sessionPGRepository) FindByRefreshToken(ctx context.Context, token string) (*model.Session, error) {
	var dbModel SessionModel
	if err := r.db.WithContext(ctx).Where("refresh_token = ?", token).First(&dbModel).Error; err != nil {
		return nil, err
	}
	return dbModel.ToDomain(), nil
}

func (r *sessionPGRepository) Update(ctx context.Context, session *model.Session) error {
	dbModel := SessionModelFromDomain(session)

	if err := r.db.WithContext(ctx).Save(dbModel).Error; err != nil {
		return err
	}
	return nil
}
