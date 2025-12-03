# Book Claims API

Sistema de gestión de reclamos multi-tenant desarrollado en Go con Gin.

## 📋 Descripción

Book Claims API permite a organizaciones (tenants) gestionar quejas y reclamos de forma aislada.

## 🚀 Características

- Multi-tenant
- Autenticación JWT (access + refresh)
- API Keys para endpoints públicos
- Gestión de roles por tenant
- Ubicaciones por tenant
- Reportes resumen

## 🔧 Tecnologías

- Go 1.21+
- Gin Framework
- Gorm + PostgreSQL
- Zap Logger
- UUID

## 📡 Endpoints (resumen)

### Públicos
- `POST /api/v1/tenant` - Crear tenant
- `GET /api/v1/tenant/:id` - Obtener tenant
- `PATCH /api/v1/tenant/:id` - Actualizar tenant
- `POST /api/v1/auth/login` - Autenticación
- `POST /api/v1/auth/refresh` - Renovar token
- `POST /api/v1/auth/logout` - Cerrar sesión

### Públicos con API Key
- `POST /api/v1/complaint` - Crear reclamo (X-API-Key)
- `GET /api/v1/complaint/code/:code` - Consultar reclamo por código público (X-API-Key)

### Privados (JWT)
- `GET /api/v1/complaint` - Listar reclamos
- `GET /api/v1/complaint/:id` - Obtener reclamo
- `POST /api/v1/complaint/:id/action` - Actualizar reclamo (guardar borrador / enviar respuesta / cambiar estado)
- `GET /api/v1/report/summary` - Reporte resumen
- `GET /api/v1/user/:id` - Obtener usuario
- `GET /api/v1/role/:id` - Obtener rol
- `POST /api/v1/location` - Crear ubicación (se utiliza ruta por tenant; ver nota)
- `GET /api/v1/location/:id` - Obtener ubicación por location_id
- `POST /api/v1/api_key` - Crear API key
- `GET /api/v1/api_key/:id` - Obtener API key

### Específicos por Tenant
- `POST /api/v1/tenant/:id/location` - Crear ubicación para tenant (tenant_id en path)
- `GET /api/v1/tenant/:id/locations` - Listar ubicaciones de un tenant
- `POST /api/v1/tenant/:id/api-keys` - Crear API key para tenant

## 🏗 Nuevos campos relevantes

- Tenant: country (Perú/España/Colombia/Chile), department, province, district, address, postal_code, logo_url.
- Location: department, province, district, postal_code, type (FISICO/ONLINE/AMBOS), url.
- Complaint: response_text, response_status (DRAFT|SENT), responder_id, response_sent_at.

## 🔁 Flujo de respuestas en Complaints

- Guardar borrador: POST `/api/v1/complaint/{id}/action` con `{ "response_text": "...", "response_status": "DRAFT" }` → guarda texto sin cambiar estado.
- Enviar respuesta: `{ "response_text": "...", "response_status": "SENT" }` → fija `response_sent_at`, asigna `responder_id` (si no viene, se toma del token) y cambia status a `ATENDIDO` si aplica.
- Cambiar sólo estado: `{ "new_status": "EN PROCESO" }`.

## 🔐 Notas de seguridad / tenant scoping

- El endpoint `POST /api/v1/tenant/:id/location` toma el tenant_id desde el path y lo usa como fuente de verdad.
- `GET /api/v1/location/:id` usa location_id (no tenant_id).
- `GET /api/v1/tenant/:id/locations` lista ubicaciones del tenant.
- Recomendado: usar JWT o API Key que incluya tenant_id y verificar coincidencia entre token y path para evitar accesos entre tenants.

## 🏃‍♂️ Ejecución y migraciones

1. Instalar dependencias

```bash
go mod tidy
```

2. Ejecutar migraciones automáticas (AutoMigrate) al arrancar:

```bash
RUN_MIGRATIONS=true go run cmd/app/main.go
```

3. Migración manual SQL (opcional):

```
internal/infrastructure/persistence/database/migrations/20251201_add_tenant_and_complaint_fields.sql
```

Aplica ese script si prefieres control explícito.

4. Ejecutar la aplicación

```bash
go run cmd/app/main.go
```

## 📚 Documentación (Swagger)

Generar docs (desde la raíz del repo):

```bash
# con swag instalado
swag init -g ./cmd/app/main.go -o internal/infrastructure/http/docs

# o sin instalar
go run github.com/swaggo/swag/cmd/swag@v1.16.6 init -g ./cmd/app/main.go -o internal/infrastructure/http/docs
```

También incluí scripts para facilitarlo:
- scripts/generate_swagger.sh
- scripts/generate_swagger.ps1

Después de generar, abre la UI en:

```
http://localhost:8080/swagger/index.html
```

## 📝 Uso rápido: crear Location (ejemplo)

POST http://localhost:8080/api/v1/tenant/{tenant_id}/location
Headers: Content-Type: application/json
Body mínimo:

{
  "name":"Sede Principal",
  "address":"Av. Principal 123",
  "department":"Lima",
  "province":"Lima",
  "district":"Miraflores",
  "type":"FISICO",
  "public_code":"SEDE-001"
}

## 📚 Estructura del proyecto

```
book-claims/
├── cmd/app/main.go
├── internal/
│   ├── core/
│   │   └── domain/model/
│   └── infrastructure/
│       └── http/
│           ├── handler/
│           ├── dto/
│           ├── docs/  # generado por swag
│           └── router.go
└── scripts/
    ├── generate_swagger.sh
    └── generate_swagger.ps1
```

---

Si quieres, añado ejemplos de payload para los endpoints de complaint (save/send) y sample responses.