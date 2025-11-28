# Book Claims API

Sistema de gestión de reclamos con arquitectura multi-tenant desarrollado en Go con Gin Framework.

## 📋 Descripción

Book Claims API es una aplicación para la gestión de reclamos que permite a diferentes organizaciones (tenants) manejar sus quejas y reclamos de manera independiente y segura.

## 🚀 Características

- **Multi-tenant**: Soporte para múltiples organizaciones
- **Autenticación JWT**: Sistema de autenticación con tokens de acceso y refresh
- **API Keys**: Control de acceso mediante claves API
- **Gestión de Roles**: Sistema de roles por tenant
- **Ubicaciones**: Manejo de ubicaciones/sucursales por tenant
- **Reportes**: Generación de reportes resumen

## 🔧 Tecnologías

- **Go 1.21+**
- **Gin Framework**: Framework web
- **Zap Logger**: Sistema de logging
- **JWT**: Autenticación
- **UUID**: Identificadores únicos

## 📡 API Endpoints

### Públicos (sin autenticación)
- `POST /api/v1/tenant` - Crear tenant
- `GET /api/v1/tenant/:id` - Obtener tenant
- `PATCH /api/v1/tenant/:id` - Actualizar tenant
- `POST /api/v1/user/login` - Login de usuario (deprecated)
- `POST /api/v1/user` - Crear usuario
- `POST /api/v1/role` - Crear rol
- `POST /api/v1/auth/login` - Autenticación
- `POST /api/v1/auth/refresh` - Renovar token
- `POST /api/v1/auth/logout` - Cerrar sesión

### Públicos con API Key
- `POST /api/v1/complaint` - Crear reclamo
- `GET /api/v1/complaint/code/:code` - Obtener reclamo por código

### Privados (requieren autenticación JWT)
- `GET /api/v1/complaint` - Listar reclamos
- `GET /api/v1/complaint/:id` - Obtener reclamo
- `POST /api/v1/complaint/:id/action` - Actualizar reclamo
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

## 🏗️ Flujo de Configuración Inicial

### Paso 1: Crear Tenant (Organización)
```bash
POST /api/v1/tenant
Content-Type: application/json

{
  "name": "Mi Empresa S.A.",
  "ruc": "12345678901",
  "email_contact": "contacto@miempresa.com",
  "phone_contact": "+51999999999",
  "is_active": true
}
```

**Respuesta**: Se obtiene el `tenant_id` que será necesario para los siguientes pasos.

### Paso 2: Crear Rol Administrativo
```bash
POST /api/v1/role
Content-Type: application/json

{
  "tenant_id": "uuid-del-tenant",
  "name": "Administrador",
  "description": "Rol administrativo con acceso completo",
  "is_system": false
}
```

**Respuesta**: Se obtiene el `role_id` del rol administrativo.

### Paso 3: Crear Ubicación Principal
```bash
POST /api/v1/tenant/{tenant_id}/location
Content-Type: application/json

{
  "name": "Sede Principal",
  "address": "Av. Principal 123",
  "tenant_id": "uuid-del-tenant"
}
```

**Respuesta**: Se obtiene el `location_id` de la ubicación principal.

### Paso 4: Crear Usuario Administrativo
```bash
POST /api/v1/user
Content-Type: application/json

{
  "tenant_id": "uuid-del-tenant",
  "role_id": "uuid-del-rol",
  "location_id": "uuid-de-ubicacion",
  "email": "admin@miempresa.com",
  "password": "password123",
  "first_name": "Admin",
  "last_name": "Sistema",
  "full_name": "Admin Sistema",
  "user_name": "admin",
  "phone": "+51999999999",
  "is_active": true
}
```

### Paso 5: Crear API Key para Reclamos Públicos
```bash
POST /api/v1/tenant/{tenant_id}/api-keys
Content-Type: application/json

{
  "name": "API Key Principal",
  "tenant_id": "uuid-del-tenant"
}
```

**Respuesta**: Se obtiene la API key que permitirá recibir reclamos desde formularios públicos.

### Paso 6: Autenticación del Usuario
```bash
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "admin",
  "password": "password123"
}
```

**Respuesta**: Se obtienen los tokens `access_token` y `refresh_token` para usar en endpoints privados.

## 🔐 Autenticación

### JWT Tokens
- **Access Token**: Para autenticar peticiones a endpoints privados
- **Refresh Token**: Para renovar access tokens expirados

### API Keys
- Se usan en endpoints públicos para crear reclamos
- Se incluyen en el header: `X-API-Key: your-api-key`

## 📝 Uso Operativo

### Para recibir un reclamo (público):
```bash
POST /api/v1/complaint
X-API-Key: your-api-key
Content-Type: application/json

{
  "title": "Problema con el servicio",
  "description": "Descripción del reclamo",
  "customer_email": "cliente@email.com",
  "customer_phone": "+51999999999"
}
```

### Para gestionar reclamos (privado):
```bash
GET /api/v1/complaint
Authorization: Bearer your-access-token
```

### Para consulta pública de reclamo:
```bash
GET /api/v1/complaint/code/ABC123
X-API-Key: your-api-key
```

## 🏃‍♂️ Ejecución

```bash
# Instalar dependencias
go mod tidy

# Ejecutar la aplicación
go run cmd/main.go
```

## 📚 Documentación (Swagger)

Se utiliza swaggo para generar la documentación OpenAPI. Instrucciones:

1. Instala la herramienta `swag` si aún no la tienes:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Genera la documentación:

```bash
cd c:\PetProject\book-claims
swag init -g cmd/app/main.go -o internal/infrastructure/http/docs
```

3. Ejecuta la aplicación y accede a la UI en:

```
http://localhost:8080/swagger/index.html
```

Nota: Ya dejé las anotaciones en los handlers y DTOs principales. Ejecuta `swag init` para generar los archivos `docs`.

## 📊 Estructura del Proyecto

```
book-claims/
├── internal/
│   ├── core/
│   │   ├── domain/model/     # Modelos de dominio
│   │   └── port/            # Interfaces/puertos
│   └── infrastructure/
│       └── http/
│           ├── handler/     # Controladores HTTP
│           ├── dto/         # DTOs para HTTP
│           ├── ctxutil/     # Utilidades de contexto
│           └── router.go    # Configuración de rutas
└── cmd/
    └── main.go             # Punto de entrada
```

## 🔍 Notas Importantes

1. **Orden de creación**: Es crucial seguir el orden: Tenant → Rol → Ubicación → Usuario → API Key
2. **Multi-tenancy**: Cada tenant opera de forma independiente
3. **Seguridad**: Los endpoints privados requieren JWT, los públicos de reclamos requieren API Key
4. **Logging**: El sistema incluye logging detallado para auditoría y debugging

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request