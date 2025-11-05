package ctxutil

import "context"

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

type ctxLocationKey string

const locationKey ctxLocationKey = "location_id"

// Guardar el location_id en el contexto
func WithLocationID(ctx context.Context, locationID string) context.Context {
	return context.WithValue(ctx, locationKey, locationID)
}

// Obtener el location_id del contexto
func GetLocationID(ctx context.Context) (string, bool) {
	locationID, ok := ctx.Value(locationKey).(string)
	return locationID, ok
}
