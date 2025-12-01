# Book Claims API

Sistema de gestión de reclamos con arquitectura multi-tenant desarrollado en Go con Gin Framework.

## 📋 Descripción

Book Claims API permite a organizaciones (tenants) gestionar reclamos y quejas de forma aislada.

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

## 📡 API Endpoints (resumen)

### Públicos (sin autenticación)
- `POST /api/v1/tenant` - Crear tenant
- `GET /api/v1/tenant/:id` - Obtener tenant
- `PATCH /api/v1/tenant/:id` - Actualizar tenant (parcial)
- `POST /api/v1/user` - Crear usuario
- `POST /api/v1/role` - Crear rol
- `POST /api/v1/auth/login` - Autenticación
- `POST /api/v1/auth/refresh` - Renovar token
- `POST /api/v1/auth/logout` - Cerrar sesión

### Públicos con API Key
- `POST /api/v1/complaint` - Crear reclamo (X-API-Key)
- `GET /api/v1/complaint/code/:code` - Consultar reclamo por código público (X-API-Key)

### Privados (requieren JWT)
- `GET /api/v1/complaint` - Listar reclamos
- `GET /api/v1/complaint/:id` - Obtener reclamo
- `POST /api/v1/complaint/:id/action` - Actualizar reclamo (guardar borrador / enviar respuesta / cambiar estado)
- `GET /api/v1/report/summary` - Reporte resumen
- `GET /api/v1/user/:id` - Obtener usuario
- `GET /api/v1/role/:id` - Obtener rol
- `POST /api/v1/location` - Crear ubicación
- `GET /api/v1/location/:id` - Obtener ubicación
- `POST /api/v1/api_key` - Crear API key
- `GET /api/v1/api_key/:id` - Obtener API key

### Específicos por Tenant
- `POST /api/v1/tenant/:id/location` - Crear ubicación para tenant
- `POST /api/v1/tenant/:id/api-keys` - Crear API key para tenant

## 🏗️ Nuevos campos relevantes

- Tenant: country (Perú/España/Colombia/Chile), department, province, district, address, postal_code, logo_url.
- Location: department, province, district, postal_code, type (FISICO/ONLINE/AMBOS), url.
- Complaint: response_text, response_status (DRAFT|SENT), responder_id, response_sent_at. Estos permiten guardar borradores de respuesta y enviar respuestas oficiales.

## 🔁 Flujo de respuestas en Complaints

- Guardar borrador: PATCH/POST `/api/v1/complaint/{id}/action` con body { "response_text": "...", "response_status": "DRAFT" }
- Enviar respuesta: `{ "response_text": "...", "response_status": "SENT" }` → la aplicación fijará `response_sent_at` y cambiará el estado del reclamo a `ATENDIDO` cuando aplique. Si no se envía `responder_id`, se usa el user_id del token.
- Cambiar solo estado: `{ "new_status": "EN PROCESO" }`

Ejemplo: enviar respuesta

```json
POST /api/v1/complaint/{id}/action
Authorization: Bearer <token>
Content-Type: application/json
{
  "response_text": "Respuesta oficial enviada al cliente.",
  "response_status": "SENT"
}
```

## 🏃‍♂️ Ejecución y migraciones

1. Instalar dependencias

```bash
go mod tidy
```

2. Ejecutar migraciones automáticas (AutoMigrate) durante arranque:

```bash
RUN_MIGRATIONS=true go run cmd/app/main.go
```

AutoMigrate actualizará las tablas del proyecto. En producción se recomienda revisar y aplicar migraciones SQL controladas.

3. Alternativamente aplicar manualmente el script SQL creado en:

```
internal/infrastructure/persistence/database/migrations/20251201_add_tenant_and_complaint_fields.sql
```

Aplica ese script a tu base de datos si necesitas control fino.

4. Ejecutar la aplicación

```bash
go run cmd/app/main.go
```

## 📚 Documentación (Swagger)

Instala la herramienta `swag` y genera docs:

```bash
go install github.com/swaggo/swag/cmd/swag@v1.16.6
cd C:\PetProject\book-claims
swag init -g cmd/app/main.go -o internal/infrastructure/http/docs
```

Luego levanta la app y accede a:

```
http://localhost:8080/swagger/index.html
```

## 🔐 Notas de seguridad

- Endpoints privados requieren JWT.
- Endpoints públicos para reclamos requieren `X-API-Key`.
- Se recomienda restringir acciones de envío de respuestas a roles administrativos (puedo añadir RoleAuthorizationMiddleware si lo deseas).

## 🤝 Contribución

1. Fork
2. Crear rama
3. Commit y PR

---

Si quieres que actualice el README con ejemplos adicionales (migraciones SQL para producción, diagramas ER o política de roles), dime cuál y lo añado.