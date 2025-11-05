package location

import (
	"claimbook-api/internal/core/domain/model"
	"time"

	"github.com/google/uuid"
)

type LocationModel struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TenantID   uuid.UUID `gorm:"type:uuid;not null;index"`
	Name       string    `gorm:"type:varchar(100);not null"`
	Address    string    `gorm:"type:varchar(255);not null"`
	Type       string    `gorm:"type:varchar(100);not null"`
	IsActive   bool      `gorm:"type:boolean;default:true;not null"`
	PublicCode string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

func (u *LocationModel) ToDomain() *model.Location {
	return &model.Location{
		ID:         u.ID,
		TenantID:   u.TenantID,
		Name:       u.Name,
		Address:    u.Address,
		Type:       u.Type,
		IsActive:   u.IsActive,
		PublicCode: u.PublicCode,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
func LocationModelFromDomain(u *model.Location) *LocationModel {
	return &LocationModel{
		ID:         u.ID,
		TenantID:   u.TenantID,
		Name:       u.Name,
		Address:    u.Address,
		Type:       u.Type,
		IsActive:   u.IsActive,
		PublicCode: u.PublicCode,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
func (LocationModel) TableName() string {
	return "location"
}
