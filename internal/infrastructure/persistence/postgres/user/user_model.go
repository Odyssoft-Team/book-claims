package user

import (
	"claimbook-api/internal/core/domain/model"
	"claimbook-api/internal/infrastructure/persistence/postgres/role"
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TenantID   uuid.UUID      `gorm:"type:uuid;not null;index"`
	RoleID     uuid.UUID      `gorm:"type:uuid;not null;index"`
	Role       role.RoleModel `gorm:"foreignKey:RoleID"`
	LocationID uuid.UUID      `gorm:"type:uuid;index"`
	Email      string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password   string         `gorm:"type:varchar(255);not null"`
	FirstName  string         `gorm:"type:varchar(100);not null"`
	LastName   string         `gorm:"type:varchar(100);not null"`
	FullName   string         `gorm:"type:varchar(200);not null"`
	UserName   string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone      string         `gorm:"type:varchar(20)"`
	IsActive   bool           `gorm:"type:boolean;default:true;not null"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at"`
}

func (u *UserModel) ToDomain() *model.User {
	return &model.User{
		ID:         u.ID,
		TenantID:   u.TenantID,
		RoleID:     u.RoleID,
		RoleName:   u.Role.Name,
		LocationID: u.LocationID,
		Email:      u.Email,
		Password:   u.Password,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		FullName:   u.FullName,
		UserName:   u.UserName,
		Phone:      u.Phone,
		IsActive:   u.IsActive,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
func UserModelFromDomain(u *model.User) *UserModel {
	return &UserModel{
		ID:         u.ID,
		TenantID:   u.TenantID,
		RoleID:     u.RoleID,
		LocationID: u.LocationID,
		Email:      u.Email,
		Password:   u.Password,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		FullName:   u.FullName,
		UserName:   u.UserName,
		Phone:      u.Phone,
		IsActive:   u.IsActive,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
func (UserModel) TableName() string {
	return "user"
}
