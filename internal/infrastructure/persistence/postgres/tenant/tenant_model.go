package tenant

import (
	"claimbook-api/internal/core/domain/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TenantModel struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name           string    `gorm:"column:tenant_name;type:varchar(100);not null;unique"`
	Ruc            string    `gorm:"column:tenant_ruc;type:varchar(20);not null;unique"`
	EmailContact   string    `gorm:"column:tenant_email_contact;not null;unique;index"`
	PhoneContact   string    `gorm:"column:tenant_phone_contact;not null"`
	IsConfirm      bool      `gorm:"column:is_confirm; notnull;default:false"`
	IsActive       bool      `gorm:"column:is_active; not null;default:false"`
	Country        string    `gorm:"column:country;type:varchar(50);default:''"`
	Department     string    `gorm:"column:department;type:varchar(100);default:''"`
	Province       string    `gorm:"column:province;type:varchar(100);default:''"`
	District       string    `gorm:"column:district;type:varchar(100);default:''"`
	Address        string    `gorm:"column:address;type:varchar(255);default:''"`
	PostalCode     string    `gorm:"column:postal_code;type:varchar(20);default:''"`
	LogoURL        string    `gorm:"column:logo_url;type:varchar(255);default:''"`
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamptz;not null"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamptz;not null"`
	gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (t *TenantModel) ToDomain() *model.Tenant {
	return &model.Tenant{
		ID:           t.ID,
		Name:         t.Name,
		Ruc:          t.Ruc,
		EmailContact: t.EmailContact,
		PhoneContact: t.PhoneContact,
		IsConfirm:    t.IsConfirm,
		IsActive:     t.IsActive,
		Country:      model.Country(t.Country),
		Department:   t.Department,
		Province:     t.Province,
		District:     t.District,
		Address:      t.Address,
		PostalCode:   t.PostalCode,
		LogoURL:      t.LogoURL,
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
	}
}

func TenantModelFromDomain(t *model.Tenant) *TenantModel {
	return &TenantModel{
		ID:           t.ID,
		Name:         t.Name,
		Ruc:          t.Ruc,
		EmailContact: t.EmailContact,
		PhoneContact: t.PhoneContact,
		IsConfirm:    t.IsConfirm,
		IsActive:     t.IsActive,
		Country:      string(t.Country),
		Department:   t.Department,
		Province:     t.Province,
		District:     t.District,
		Address:      t.Address,
		PostalCode:   t.PostalCode,
		LogoURL:      t.LogoURL,
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
	}
}

func (TenantModel) TableName() string {
	return "tenant"
}
