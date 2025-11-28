# 📡 Documentación de la API

Documentación detallada de todos los endpoints disponibles en la API.

**Base URL:** `http://localhost:8081/api/v1`

---

## 🔑 Autenticación

La API usa **JWT (JSON Web Tokens)** para autenticar usuarios en rutas protegidas.

### Obtener un Token

Primero debes registrarte y luego iniciar sesión para obtener un token:

```http
POST /api/v1/seguridad/login
Content-Type: application/json

{
  "correo": "tu_correo@example.com",
  "password": "tu_contraseña"
}
```

**Respuesta exitosa:**
```json
{
  "estado": "ok",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "mensaje": "Login exitoso"
}
```

### Usar el Token

Para acceder a rutas protegidas, incluye el token en el header `Authorization`:

```http
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

## 👤 Seguridad/Autenticación

### Registrar Usuario

Crea una nueva cuenta de usuario.

**Endpoint:** `POST /seguridad/registro`  
**Autenticación:** No requerida

**Request Body:**
```json
{
  "nombre": "Juan Pérez",
  "correo": "juan@example.com",
  "password": "miPassword123"
}
```

**Respuesta exitosa (201):**
```json
{
  "estado": "ok",
  "mensaje": "Usuario registrado correctamente. Por favor verifica tu correo."
}
```

**Notas:**
- Se envía un correo de verificación al email proporcionado
- El usuario queda con estado "Inactivo" hasta verificar el correo
- La contraseña se hashea con bcrypt antes de guardarse

---

### Verificar Cuenta

Verifica la cuenta de usuario mediante el token enviado por correo.

**Endpoint:** `GET /seguridad/verificacion/:token`  
**Autenticación:** No requerida

**Ejemplo:**
```http
GET /api/v1/seguridad/verificacion/550e8400-e29b-41d4-a716-446655440000
```

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "mensaje": "Cuenta verificada correctamente"
}
```

---

### Iniciar Sesión

Autentica un usuario y obtiene un token JWT.

**Endpoint:** `POST /seguridad/login`  
**Autenticación:** No requerida

**Request Body:**
```json
{
  "correo": "juan@example.com",
  "password": "miPassword123"
}
```

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb3JyZW8iOiJqdWFuQGV4YW1wbGUuY29tIiwiaWQiOjEsIm5vbWJyZSI6Ikp1YW4gUMOpcmV6IiwiaWF0IjoxNzAwMDAwMDAwLCJleHAiOjE3MDAwODY0MDB9.xyz",
  "mensaje": "Login exitoso"
}
```

**Respuesta de error (401):**
```json
{
  "estado": "error",
  "mensaje": "Credenciales incorrectas"
}
```

**Notas:**
- El token expira en 24 horas
- La cuenta debe estar verificada (estado "Activo")

---

## 🏷️ Categorías

### Listar Categorías

Obtiene todas las categorías (no eliminadas).

**Endpoint:** `GET /categorias`  
**Autenticación:** No requerida

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "datos": [
    {
      "id": 1,
      "nombre": "Bebidas",
      "slug": "bebidas",
      "created_at": "2025-11-01T10:00:00Z",
      "updated_at": "2025-11-01T10:00:00Z",
      "deleted_at": null
    },
    {
      "id": 2,
      "nombre": "Postres",
      "slug": "postres",
      "created_at": "2025-11-01T10:05:00Z",
      "updated_at": "2025-11-01T10:05:00Z",
      "deleted_at": null
    }
  ]
}
```

---

### Obtener Categoría por ID

Obtiene una categoría específica.

**Endpoint:** `GET /categorias/:id`  
**Autenticación:** No requerida

**Ejemplo:**
```http
GET /api/v1/categorias/1
```

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "datos": {
    "id": 1,
    "nombre": "Bebidas",
    "slug": "bebidas",
    "created_at": "2025-11-01T10:00:00Z",
    "updated_at": "2025-11-01T10:00:00Z",
    "deleted_at": null
  }
}
```

**Respuesta de error (404):**
```json
{
  "estado": "error",
  "mensaje": "Recurso no disponible"
}
```

---

### Crear Categoría

Crea una nueva categoría.

**Endpoint:** `POST /categorias`  
**Autenticación:** ✅ JWT requerido

**Request Body:**
```json
{
  "nombre": "Platos principales"
}
```

**Respuesta exitosa (201):**
```json
{
  "estado": "ok",
  "datos": {
    "id": 3,
    "nombre": "Platos principales",
    "slug": "platos-principales",
    "created_at": "2025-11-27T15:30:00Z",
    "updated_at": "2025-11-27T15:30:00Z",
    "deleted_at": null
  }
}
```

**Notas:**
- El slug se genera automáticamente del nombre
- Valida que no exista otra categoría con el mismo nombre

---

### Actualizar Categoría

Actualiza una categoría existente.

**Endpoint:** `PUT /categorias/:id`  
**Autenticación:** ✅ JWT requerido

**Request Body:**
```json
{
  "nombre": "Bebidas refrescantes"
}
```

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "datos": {
    "id": 1,
    "nombre": "Bebidas refrescantes",
    "slug": "bebidas-refrescantes",
    "created_at": "2025-11-01T10:00:00Z",
    "updated_at": "2025-11-27T15:35:00Z",
    "deleted_at": null
  }
}
```

---

### Eliminar Categoría

Elimina una categoría (soft delete).

**Endpoint:** `DELETE /categorias/:id`  
**Autenticación:** ✅ JWT requerido

**Ejemplo:**
```http
DELETE /api/v1/categorias/1
Authorization: Bearer <tu_token>
```

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "mensaje": "Categoría eliminada correctamente"
}
```

**Notas:**
- La categoría no se elimina físicamente, solo se marca como eliminada (soft delete)
- Puedes restaurarla con un UPDATE en la BD

---

## 🍽️ Recetas

### Listar Recetas

Obtiene todas las recetas con sus categorías y usuarios.

**Endpoint:** `GET /recetas`  
**Autenticación:** No requerida

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "datos": [
    {
      "id": 1,
      "nombre": "Pastel de chocolate",
      "slug": "pastel-de-chocolate",
      "categoria_id": 2,
      "categoria": "Postres",
      "usuario_id": 1,
      "usuario": "Juan Pérez",
      "tiempo": "45 min",
      "foto": "pastel.jpg",
      "descripcion": "Delicioso pastel de chocolate con cobertura",
      "fecha": "27/11/2025"
    }
  ]
}
```

---

### Obtener Receta por ID

Obtiene una receta específica con todos sus datos.

**Endpoint:** `GET /recetas/:id`  
**Autenticación:** No requerida

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "datos": {
    "id": 1,
    "nombre": "Pastel de chocolate",
    "slug": "pastel-de-chocolate",
    "categoria_id": 2,
    "categoria": "Postres",
    "usuario_id": 1,
    "usuario": "Juan Pérez",
    "tiempo": "45 min",
    "foto": "pastel.jpg",
    "descripcion": "Delicioso pastel de chocolate con cobertura de ganache",
    "fecha": "27/11/2025"
  }
}
```

---

### Crear Receta

Crea una nueva receta.

**Endpoint:** `POST /recetas`  
**Autenticación:** ✅ JWT requerido

**Request Body:**
```json
{
  "nombre": "Tarta de manzana",
  "tiempo": "60 min",
  "descripcion": "Tarta casera de manzana con canela",
  "categoria_id": 2
}
```

**Respuesta exitosa (201):**
```json
{
  "estado": "ok",
  "datos": {
    "id": 5,
    "nombre": "Tarta de manzana",
    "slug": "tarta-de-manzana",
    "categoria_id": 2,
    "usuario_id": 1,
    "tiempo": "60 min",
    "foto": "img.png",
    "descripcion": "Tarta casera de manzana con canela",
    "fecha": "2025-11-27T15:40:00Z"
  }
}
```

**Notas:**
- El usuario_id se obtiene automáticamente del JWT
- La foto por defecto es "img.png" (se puede cambiar luego)
- El slug se genera automáticamente del nombre

---

### Actualizar Receta

Actualiza una receta existente.

**Endpoint:** `PUT /recetas/:id`  
**Autenticación:** ✅ JWT requerido

**Request Body:**
```json
{
  "nombre": "Tarta de manzana con helado",
  "tiempo": "65 min",
  "descripcion": "Tarta casera de manzana con canela, servida con helado de vainilla",
  "categoria_id": 2
}
```

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "datos": {
    "id": 5,
    "nombre": "Tarta de manzana con helado",
    "slug": "tarta-de-manzana-con-helado",
    "categoria_id": 2,
    "usuario_id": 1,
    "tiempo": "65 min",
    "foto": "img.png",
    "descripcion": "Tarta casera de manzana con canela, servida con helado de vainilla",
    "fecha": "2025-11-27T15:40:00Z",
    "updated_at": "2025-11-27T16:00:00Z"
  }
}
```

---

### Eliminar Receta

Elimina una receta (soft delete) y su foto asociada.

**Endpoint:** `DELETE /recetas/:id`  
**Autenticación:** ✅ JWT requerido

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "mensaje": "Receta eliminada correctamente"
}
```

**Notas:**
- Elimina físicamente la foto del servidor (`public/recetas/`)
- La receta se marca como eliminada en la BD (soft delete)

---

## 🔍 Recetas - Helpers (Búsqueda y Filtros)

### Recetas para Home

Obtiene recetas para mostrar en la página principal.

**Endpoint:** `GET /recetas-helpers/home`  
**Autenticación:** No requerida

**Respuesta:** Similar a `GET /recetas`

---

### Buscar Receta por Slug

Obtiene una receta por su slug (URL amigable).

**Endpoint:** `GET /recetas-helpers/slug/:slug`  
**Autenticación:** No requerida

**Ejemplo:**
```http
GET /api/v1/recetas-helpers/slug/pastel-de-chocolate
```

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "datos": {
    "id": 1,
    "nombre": "Pastel de chocolate",
    "slug": "pastel-de-chocolate",
    ...
  }
}
```

---

### Buscador de Recetas

Busca recetas por categoría y/o texto.

**Endpoint:** `GET /recetas-helpers/buscador`  
**Autenticación:** No requerida

**Query Parameters:**
- `categoria_id` (opcional): ID de la categoría
- `search` (opcional): Texto a buscar en nombre/descripción

**Ejemplos:**
```http
GET /api/v1/recetas-helpers/buscador?categoria_id=2
GET /api/v1/recetas-helpers/buscador?search=chocolate
GET /api/v1/recetas-helpers/buscador?categoria_id=2&search=chocolate
```

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "datos": [
    {
      "id": 1,
      "nombre": "Pastel de chocolate",
      "categoria": "Postres",
      ...
    }
  ]
}
```

---

### Recetas de un Usuario

Obtiene todas las recetas de un usuario específico.

**Endpoint:** `GET /recetas-helpers/usuarios/:id`  
**Autenticación:** ✅ JWT requerido

**Ejemplo:**
```http
GET /api/v1/recetas-helpers/usuarios/1
Authorization: Bearer <tu_token>
```

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "datos": [
    {
      "id": 1,
      "nombre": "Pastel de chocolate",
      "usuario": "Juan Pérez",
      ...
    }
  ]
}
```

---

### Subir Foto de Receta

Sube una foto para una receta existente.

**Endpoint:** `POST /recetas-helpers/foto`  
**Autenticación:** No requerida

**Request (multipart/form-data):**
- `foto`: Archivo de imagen (JPG, PNG, etc.)
- `id`: ID de la receta

**Ejemplo con curl:**
```bash
curl -X POST http://localhost:8081/api/v1/recetas-helpers/foto \
  -F "foto=@/ruta/a/imagen.jpg" \
  -F "id=1"
```

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "mensaje": "Foto actualizada correctamente",
  "foto": "upload_1234567890.jpg"
}
```

---

## 📧 Contacto

### Enviar Mensaje de Contacto

Envía un mensaje de contacto (se guarda en BD y se envía por email).

**Endpoint:** `POST /contactanos`  
**Autenticación:** No requerida

**Request Body:**
```json
{
  "nombre": "María López",
  "correo": "maria@example.com",
  "telefono": "+51 999 888 777",
  "mensaje": "Me gustaría más información sobre sus recetas vegetarianas."
}
```

**Respuesta exitosa (200):**
```json
{
  "estado": "ok",
  "mensaje": "Mensaje enviado correctamente"
}
```

**Notas:**
- El mensaje se guarda en la tabla `contacto`
- Se envía un email al administrador con los datos del contacto

---

## 📋 Códigos de Estado HTTP

| Código | Significado | Cuándo se usa |
|--------|-------------|---------------|
| 200 | OK | Operación exitosa (GET, PUT, DELETE) |
| 201 | Created | Recurso creado exitosamente (POST) |
| 204 | No Content | Petición OPTIONS (CORS preflight) |
| 400 | Bad Request | Datos inválidos o faltantes |
| 401 | Unauthorized | Token inválido o expirado |
| 404 | Not Found | Recurso no encontrado |
| 500 | Internal Server Error | Error del servidor |

---

## 🧪 Ejemplos con curl (PowerShell)

### Registrar usuario
```powershell
curl -X POST http://localhost:8081/api/v1/seguridad/registro `
  -H "Content-Type: application/json" `
  -d '{"nombre":"Juan","correo":"juan@example.com","password":"pass123"}'
```

### Login
```powershell
curl -X POST http://localhost:8081/api/v1/seguridad/login `
  -H "Content-Type: application/json" `
  -d '{"correo":"juan@example.com","password":"pass123"}'
```

### Crear categoría (con JWT)
```powershell
curl -X POST http://localhost:8081/api/v1/categorias `
  -H "Content-Type: application/json" `
  -H "Authorization: Bearer TU_TOKEN_AQUI" `
  -d '{"nombre":"Sopas"}'
```

### Buscar recetas
```powershell
curl "http://localhost:8081/api/v1/recetas-helpers/buscador?categoria_id=1&search=chocolate"
```

---

## 📝 Notas Importantes

### Archivos Estáticos

Las imágenes subidas están disponibles públicamente en:

```
http://localhost:8081/public/recetas/nombre_imagen.jpg
http://localhost:8081/public/uploads/fotos/nombre_imagen.jpg
```

### Validaciones

- **Email:** Debe ser un formato válido
- **Contraseña:** Sin restricciones específicas (recomendado: mínimo 8 caracteres)
- **Nombres:** No pueden estar vacíos
- **IDs:** Deben existir en la base de datos

### Rate Limiting

Actualmente no hay rate limiting implementado. Considera agregarlo en producción.

---

## 🔄 Versionado de la API

Versión actual: **v1**

Todas las rutas están bajo el prefijo `/api/v1/`

---

## 🆘 Manejo de Errores

Todas las respuestas de error siguen este formato:

```json
{
  "estado": "error",
  "mensaje": "Descripción general del error",
  "error": "Detalles específicos (opcional)"
}
```

---

¿Preguntas? Revisa el [README](README.md) o abre un issue en GitHub.

