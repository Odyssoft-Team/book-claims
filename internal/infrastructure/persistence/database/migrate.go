package database

import (
	apikey "claimbook-api/internal/infrastructure/persistence/postgres/api_key"
	"claimbook-api/internal/infrastructure/persistence/postgres/complaint"
	"claimbook-api/internal/infrastructure/persistence/postgres/location"
	"claimbook-api/internal/infrastructure/persistence/postgres/role"
	"claimbook-api/internal/infrastructure/persistence/postgres/session"
	"claimbook-api/internal/infrastructure/persistence/postgres/tenant"
	"claimbook-api/internal/infrastructure/persistence/postgres/user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&complaint.ComplaintModel{},
		&user.UserModel{},
		&role.RoleModel{},
		&location.LocationModel{},
		&session.SessionModel{},
		&tenant.TenantModel{},
		&apikey.ApiKeyModel{},
	)
	if err != nil {
		return err
	}
	return nil
}
