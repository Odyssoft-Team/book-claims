package role

import (
	"claimbook-api/internal/core/domain/model"
	"time"

	"github.com/google/uuid"
)

type RoleModel struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name        string    `gorm:"type:varchar(100);not null;uniqueIndex:idx_role_name_tenant"`
	Description string    `gorm:"type:varchar(255)"`
	TenantID    uuid.UUID `gorm:"type:uuid;index;uniqueIndex:idx_role_name_tenant"`
	IsSystem    bool      `gorm:"type:boolean;default:false;not null"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (r *RoleModel) ToDomain() *model.Role {
	return &model.Role{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		TenantID:    r.TenantID,
		IsSystem:    r.IsSystem,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

func RoleModelFromDomain(r *model.Role) *RoleModel {
	return &RoleModel{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		TenantID:    r.TenantID,
		IsSystem:    r.IsSystem,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}
func (RoleModel) TableName() string {
	return "role"
}
