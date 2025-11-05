package postgres

import (
	"context"

	"gorm.io/gorm"

	"gorm.io/gorm/clause"
)

func RegisterTenantScope(db *gorm.DB) {
	db.Callback().Query().Before("gorm:query").Register("tenant:filter", func(tx *gorm.DB) {
		tenantID, ok := GetTenantID(tx.Statement.Context)
		if !ok {
			return
		}

		// Verifica si el modelo tiene un campo TenantID
		if tx.Statement.Schema != nil {
			if _, exists := tx.Statement.Schema.FieldsByName["TenantID"]; exists {
				tx.Statement.AddClause(clause.Where{
					Exprs: []clause.Expression{
						clause.Expr{SQL: "tenant_id = ?", Vars: []interface{}{tenantID}},
					},
				})
			}
		}
	})
}

type ctxTenantKey string

const tenantKey ctxTenantKey = "tenant_id"

// Guardar el tenant_id en el contexto
func WithTenantID(ctx context.Context, tenantID string) context.Context {
	return context.WithValue(ctx, tenantKey, tenantID)
}

// Obtener el tenant_id del contexto
func GetTenantID(ctx context.Context) (string, bool) {
	tenantID, ok := ctx.Value(tenantKey).(string)
	return tenantID, ok
}
