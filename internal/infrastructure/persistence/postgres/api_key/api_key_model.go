package apikey

import (
	"claimbook-api/internal/core/domain/model"
	"time"

	"github.com/google/uuid"
)

type ApiKeyModel struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TenantID   uuid.UUID `gorm:"type:uuid;not null;index"`
	LocationID uuid.UUID `gorm:"type:uuid;index"`
	ApiKey     string    `gorm:"type:text"`
	Scope      string    `gorm:"type:text"`
	IsActive   bool      `gorm:"type:boolean;default:true;not null"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

func (u *ApiKeyModel) ToDomain() *model.ApiKey {
	return &model.ApiKey{
		ID:         u.ID,
		TenantID:   u.TenantID,
		LocationID: u.LocationID,
		ApiKey:     u.ApiKey,
		Scope:      u.Scope,
		IsActive:   u.IsActive,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
func ApiKeyModelFromDomain(u *model.ApiKey) *ApiKeyModel {
	return &ApiKeyModel{
		ID:         u.ID,
		TenantID:   u.TenantID,
		LocationID: u.LocationID,
		ApiKey:     u.ApiKey,
		Scope:      u.Scope,
		IsActive:   u.IsActive,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
func (ApiKeyModel) TableName() string {
	return "api_key"
}
