package complaint

import (
	"claimbook-api/internal/core/domain/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ComplaintModel struct {
	ID              uuid.UUID             `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TenantID        uuid.UUID             `gorm:"type:uuid;not null;index"`
	LocationID      uuid.UUID             `gorm:"type:uuid;not null;index"`
	Type            model.ComplaintType   `gorm:"type:varchar(50);not null;index"`
	Status          model.ComplaintStatus `gorm:"type:varchar(50);not null;index"`
	CategoryID      uuid.UUID             `gorm:"type:uuid;not null;index"`
	Source          model.ComplaintSource `gorm:"type:varchar(50);not null;index"`
	ApiKeyID        uuid.UUID             `gorm:"type:uuid;not null;index"`
	CodePublic      string                `gorm:"type:varchar(100);uniqueIndex;not null"`
	Description     string                `gorm:"type:text;not null"`
	RequestedAction string                `gorm:"type:text"`
	IsClosed        bool
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at"`
	ResolvedAt      *time.Time `json:"resolved_at" db:"resolved_at"`
}

// Hooks opcionales
func (g *ComplaintModel) BeforeCreate(tx *gorm.DB) error {
	return nil
}

func (g *ComplaintModel) ToDomain() *model.Complaint {
	return &model.Complaint{
		ID:              g.ID,
		TenantID:        g.TenantID,
		LocationID:      g.LocationID,
		Type:            g.Type,
		Status:          g.Status,
		CategoryID:      g.CategoryID,
		Source:          g.Source,
		ApiKeyID:        g.ApiKeyID,
		CodePublic:      g.CodePublic,
		Description:     g.Description,
		RequestedAction: g.RequestedAction,
		CreatedAt:       g.CreatedAt,
		UpdatedAt:       g.UpdatedAt,
		ResolvedAt:      g.ResolvedAt,
		IsClosed:        g.IsClosed,
	}
}

func ComplaintModelFromDomain(g *model.Complaint) *ComplaintModel {
	return &ComplaintModel{
		ID:              g.ID,
		TenantID:        g.TenantID,
		LocationID:      g.LocationID,
		Type:            g.Type,
		Status:          g.Status,
		CategoryID:      g.CategoryID,
		Source:          g.Source,
		ApiKeyID:        g.ApiKeyID,
		CodePublic:      g.CodePublic,
		Description:     g.Description,
		RequestedAction: g.RequestedAction,
		CreatedAt:       g.CreatedAt,
		UpdatedAt:       g.UpdatedAt,
		ResolvedAt:      g.ResolvedAt,
		IsClosed:        g.IsClosed,
	}
}

func (ComplaintModel) TableName() string {
	return "complaint"
}
