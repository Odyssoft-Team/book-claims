package session

import (
	"claimbook-api/internal/core/domain/model"
	"time"

	"github.com/google/uuid"
)

type SessionModel struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;"`
	TenantID     uuid.UUID `gorm:"type:uuid;not null;"`
	RefreshToken string    `gorm:"type:text;not null;unique"`
	IP           string    `gorm:"type:varchar(45)"`
	UserAgent    string    `gorm:"type:text"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	ExpiresAt    time.Time `gorm:"not null"`
	Revoked      bool      `gorm:"not null;default:false"`
}

func (s *SessionModel) ToDomain() *model.Session {
	return &model.Session{
		ID:           s.ID,
		UserID:       s.UserID,
		TenantID:     s.TenantID,
		RefreshToken: s.RefreshToken,
		IP:           s.IP,
		UserAgent:    s.UserAgent,
		CreatedAt:    s.CreatedAt,
		ExpiresAt:    s.ExpiresAt,
		Revoked:      s.Revoked,
	}
}

func SessionModelFromDomain(s *model.Session) *SessionModel {
	return &SessionModel{
		ID:           s.ID,
		UserID:       s.UserID,
		TenantID:     s.TenantID,
		RefreshToken: s.RefreshToken,
		IP:           s.IP,
		UserAgent:    s.UserAgent,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
		ExpiresAt:    s.ExpiresAt,
		Revoked:      s.Revoked,
	}
}

func (SessionModel) TableName() string {
	return "session"
}
