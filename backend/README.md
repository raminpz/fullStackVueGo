# 🍳 API REST - Backend Full Stack (Go + Gin + GORM)

API RESTful desarrollada con **Go**, utilizando el framework **Gin** para el manejo de rutas HTTP y **GORM** como ORM para la gestión de la base de datos MySQL. Este proyecto está diseñado para gestionar un sistema de recetas de cocina con autenticación JWT, envío de correos y administración de usuarios.

---

## 📋 Tabla de Contenidos

- [Características](#-características)
- [Tecnologías Utilizadas](#-tecnologías-utilizadas)
- [Estructura del Proyecto](#-estructura-del-proyecto)
- [Requisitos Previos](#-requisitos-previos)
- [Instalación](#-instalación)
- [Configuración](#-configuración)
- [Migraciones de Base de Datos](#-migraciones-de-base-de-datos)
- [Endpoints de la API](#-endpoints-de-la-api)
- [Modelos de Datos](#-modelos-de-datos)
- [Autenticación y Autorización](#-autenticación-y-autorización)
- [Middleware](#-middleware)
- [Utilidades](#-utilidades)
- [Docker](#-docker)
- [Scripts Útiles](#-scripts-útiles)

---

## ✨ Características

- ✅ **API RESTful** completa con operaciones CRUD
- ✅ **Autenticación JWT** (JSON Web Tokens) con expiración de 24 horas
- ✅ **Registro de usuarios** con verificación por correo electrónico
- ✅ **Gestión de recetas** con categorías, usuarios y fotos
- ✅ **Subida de archivos** (imágenes de recetas)
- ✅ **Soft Delete** en todos los modelos (DeletedAt)
- ✅ **Envío de correos** para verificación de cuenta y contacto
- ✅ **Validaciones** de datos con DTOs
- ✅ **CORS** configurado para integraciones frontend
- ✅ **Búsqueda y filtrado** de recetas
- ✅ **Pool de conexiones** MySQL optimizado
- ✅ **Manejo de errores** centralizado
- ✅ **Variables de entorno** para configuración segura
- ✅ **Docker Compose** para MySQL

---

## 🛠 Tecnologías Utilizadas

| Tecnología | Versión | Descripción |
|-----------|---------|-------------|
| **Go** | 1.24+ | Lenguaje de programación principal |
| **Gin** | v1.10.1 | Framework web minimalista y rápido |
| **GORM** | v1.31.0 | ORM para Go con soporte MySQL |
| **MySQL** | 8.0+ | Base de datos relacional |
| **JWT** | v5.3.0 | Autenticación con tokens |
| **Gomail** | v2 | Envío de correos electrónicos |
| **UUID** | v1.6.0 | Generación de tokens únicos |
| **Godotenv** | v1.5.1 | Carga de variables de entorno |
| **bcrypt** | - | Hashing de contraseñas |
| **Slug** | v1.15.0 | Generación de URLs amigables |

---

## 📁 Estructura del Proyecto

```
backend/
├── database/
│   └── database.go          # Configuración y conexión a MySQL
├── dto/
│   └── dto.go               # Data Transfer Objects (validación)
├── jwt/
│   └── jwt.go               # Generación y validación de tokens JWT
├── middleware/
│   └── middlware.go         # Middleware de autenticación JWT
├── models/
│   └── modelos.go           # Modelos de datos (GORM)
├── public/
│   ├── recetas/             # Imágenes de recetas subidas
│   └── uploads/
│       └── fotos/           # Imágenes de usuarios
├── rutas/
│   ├── categorias.go        # Endpoints de categorías
│   ├── contactanos.go       # Endpoint de contacto
│   ├── ejemplo.go           # Endpoints de ejemplo/prueba
│   ├── recetas.go           # Endpoints de recetas (CRUD)
│   ├── rutas_helper.go      # Endpoints auxiliares (búsqueda, filtros)
│   └── seguridad.go         # Endpoints de autenticación
├── utilidades/
│   └── utilidades.go        # Funciones auxiliares (envío de correos)
├── validaciones/
│   └── validaciones.go      # Validaciones personalizadas
├── .env                     # Variables de entorno (NO subir a Git)
├── .env.example             # Ejemplo de variables de entorno
├── docker-compose.yml       # Configuración de MySQL en Docker
├── go.mod                   # Dependencias del proyecto
├── go.sum                   # Checksums de dependencias
├── main.go                  # Punto de entrada de la aplicación
└── README.md                # Este archivo
```

---

## 📦 Requisitos Previos

Antes de comenzar, asegúrate de tener instalado:

- **Go** 1.24 o superior - [Descargar](https://golang.org/dl/)
- **MySQL** 8.0+ o **Docker** - [Descargar MySQL](https://dev.mysql.com/downloads/) | [Descargar Docker](https://www.docker.com/get-started)
- **Git** - [Descargar](https://git-scm.com/downloads)
- **Postman** o **Thunder Client** (opcional, para probar la API)

---

## 🚀 Instalación

### 1. Clonar el repositorio

```bash
git clone <URL_DEL_REPOSITORIO>
cd fullStackVueGo/backend
```

### 2. Instalar dependencias de Go

```bash
go mod download
```

### 3. Configurar MySQL con Docker (Opcional)

Si no tienes MySQL instalado localmente, puedes usar Docker:

```bash
docker-compose up -d
```

Esto levantará un contenedor MySQL 8 en el puerto `3306` con:
- Base de datos: `go_fullstack`
- Usuario: `root`
- Contraseña: `xxxx`

### 4. Crear archivo de variables de entorno

Copia el archivo de ejemplo y configúralo con tus credenciales:

```bash
cp .env.example .env
```

Edita `.env` con tus valores (ver sección [Configuración](#-configuración))

---

## ⚙️ Configuración

Edita el archivo `.env` con las siguientes variables:

```env
# Puerto de la aplicación
PORT=8081

# Configuración de la base de datos MySQL
DB_NAME=go_fullstack
DB_USER=root
DB_PASSWORD=rami123
DB_SERVER=localhost
DB_PORT=3306

# Clave secreta para JWT (genera una clave segura)
SECRET_JWT=tu_clave_secreta_muy_segura_aqui

# Configuración SMTP para envío de correos
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=tu_correo@gmail.com
SMTP_PASSWORD=tu_contraseña_de_aplicacion
SMTP_FROM_EMAIL=noreply@example.com

# URL del frontend para verificación de cuenta
RUTA_FRONTEND=http://localhost:3000
```

### 📧 Configuración de Gmail para SMTP

Si usas Gmail, necesitas generar una **contraseña de aplicación**:

1. Ve a tu [cuenta de Google](https://myaccount.google.com/)
2. Seguridad → Verificación en dos pasos (actívala si no lo está)
3. Contraseñas de aplicaciones → Generar nueva
4. Copia la contraseña generada en `SMTP_PASSWORD`

---

## 🗄️ Migraciones de Base de Datos

El proyecto utiliza **AutoMigrate** de GORM para crear automáticamente las tablas al iniciar la aplicación.

### Tablas creadas automáticamente:

- `categoria` - Categorías de recetas (Bebidas, Sopas, Postres, etc.)
- `receta` - Recetas con relaciones a categoría y usuario
- `contacto` - Mensajes de contacto
- `estado` - Estados de usuarios (Activo/Inactivo)
- `usuario` - Usuarios registrados

### Insertar datos iniciales (Estados)

Después de iniciar la aplicación por primera vez, ejecuta este SQL para crear los estados:

```sql
INSERT INTO go_fullstack.estados(nombre) VALUES ('Activo'), ('Inactivo');
```

---

## 🌐 Endpoints de la API

### Base URL
```
http://localhost:8081/api/v1
```

---

### 📌 **Autenticación y Seguridad**

| Método | Endpoint | Descripción | Auth |
|--------|----------|-------------|------|
| POST | `/seguridad/registro` | Registrar nuevo usuario | ❌ |
| GET | `/seguridad/verificacion/:token` | Verificar cuenta por email | ❌ |
| POST | `/seguridad/login` | Iniciar sesión (devuelve JWT) | ❌ |

#### Ejemplo: Registro de usuario

```json
POST /api/v1/seguridad/registro
Content-Type: application/json

{
  "nombre": "Juan Pérez",
  "correo": "juan@example.com",
  "password": "tu_password_seguro"
}
```

#### Ejemplo: Login

```json
POST /api/v1/seguridad/login
Content-Type: application/json

{
  "correo": "juan@example.com",
  "password": "tu_password_seguro"
}

// Respuesta:
{
  "estado": "ok",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "mensaje": "Login exitoso"
}
```

---

### 🏷️ **Categorías**

| Método | Endpoint | Descripción | Auth |
|--------|----------|-------------|------|
| GET | `/categorias` | Obtener todas las categorías | ❌ |
| GET | `/categorias/:id` | Obtener categoría por ID | ❌ |
| POST | `/categorias` | Crear nueva categoría | ✅ JWT |
| PUT | `/categorias/:id` | Actualizar categoría | ✅ JWT |
| DELETE | `/categorias/:id` | Eliminar categoría (soft delete) | ✅ JWT |

#### Ejemplo: Crear categoría

```json
POST /api/v1/categorias
Authorization: Bearer <TOKEN_JWT>
Content-Type: application/json

{
  "nombre": "Postres"
}
```

---

### 🍽️ **Recetas**

| Método | Endpoint | Descripción | Auth |
|--------|----------|-------------|------|
| GET | `/recetas` | Obtener todas las recetas | ❌ |
| GET | `/recetas/:id` | Obtener receta por ID | ❌ |
| POST | `/recetas` | Crear nueva receta | ✅ JWT |
| PUT | `/recetas/:id` | Actualizar receta | ✅ JWT |
| DELETE | `/recetas/:id` | Eliminar receta | ✅ JWT |

#### Ejemplo: Crear receta

```json
POST /api/v1/recetas
Authorization: Bearer <TOKEN_JWT>
Content-Type: application/json

{
  "nombre": "Pastel de chocolate",
  "tiempo": "45 min",
  "descripcion": "Delicioso pastel de chocolate con cobertura",
  "categoria_id": 3
}
```

---

### 🔍 **Recetas - Helpers (Búsqueda y Filtros)**

| Método | Endpoint | Descripción | Auth |
|--------|----------|-------------|------|
| GET | `/recetas-helpers/home` | Recetas para página principal | ❌ |
| GET | `/recetas-helpers/slug/:slug` | Obtener receta por slug | ❌ |
| GET | `/recetas-helpers/buscador` | Buscar recetas (query params) | ❌ |
| GET | `/recetas-helpers/usuarios/:id` | Recetas de un usuario | ✅ JWT |
| POST | `/recetas-helpers/foto` | Subir foto de receta | ❌ |

#### Ejemplo: Buscar recetas

```
GET /api/v1/recetas-helpers/buscador?categoria_id=1&search=chocolate
```

---

### 📧 **Contacto**

| Método | Endpoint | Descripción | Auth |
|--------|----------|-------------|------|
| POST | `/contactanos` | Enviar mensaje de contacto | ❌ |

#### Ejemplo: Contacto

```json
POST /api/v1/contactanos
Content-Type: application/json

{
  "nombre": "María López",
  "correo": "maria@example.com",
  "telefono": "+51 999 888 777",
  "mensaje": "Me gustaría más información sobre sus recetas."
}
```

---

### 🧪 **Endpoints de Ejemplo (Testing)**

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/ejemplo` | Ejemplo GET |
| POST | `/ejemplo` | Ejemplo POST |
| PUT | `/ejemplo/:id` | Ejemplo PUT |
| DELETE | `/ejemplo/:id` | Ejemplo DELETE |
| GET | `/ejemplo/:id` | Ejemplo con parámetros |
| GET | `/ejemplo-querystring` | Ejemplo con query strings |
| POST | `/upload` | Subir archivo |

---

## 📊 Modelos de Datos

### 👤 Usuario

```go
type Usuario struct {
    ID       uint      `json:"id"`
    EstadoID uint      `json:"estado_id"`
    Estado   *Estado   `json:"estado"`
    Nombre   string    `json:"nombre"`
    Correo   string    `json:"correo"`
    Password string    `json:"password"` // Hasheado con bcrypt
    Token    string    `json:"token"`    // Token de verificación UUID
    Fecha    time.Time `json:"fecha"`
}
```

### 🏷️ Categoría

```go
type Categoria struct {
    ID        uint           `json:"id"`
    Nombre    string         `json:"nombre"`
    Slug      string         `json:"slug"` // URL amigable
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at"` // Soft delete
}
```

### 🍽️ Receta

```go
type Receta struct {
    ID          uint           `json:"id"`
    CategoriaID uint           `json:"categoria_id"`
    UsuarioID   uint           `json:"usuario_id"`
    Usuario     *Usuario       `json:"usuario"`
    Categoria   *Categoria     `json:"categoria"`
    Nombre      string         `json:"nombre"`
    Slug        string         `json:"slug"`
    Tiempo      string         `json:"tiempo"`
    Foto        string         `json:"foto"`
    Descripcion string         `json:"descripcion"`
    Fecha       time.Time      `json:"fecha"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
```

### 📧 Contacto

```go
type Contacto struct {
    Id       uint      `json:"id"`
    Nombre   string    `json:"nombre"`
    Correo   string    `json:"correo"`
    Telefono string    `json:"telefono"`
    Mensaje  string    `json:"mensaje"`
    Fecha    time.Time `json:"fecha"`
}
```

### 🚦 Estado

```go
type Estado struct {
    ID     uint   `json:"id"`
    Nombre string `json:"nombre"` // "Activo" o "Inactivo"
}
```

---

## 🔐 Autenticación y Autorización

### 🎫 JWT (JSON Web Tokens)

El sistema utiliza JWT para la autenticación. Cuando un usuario inicia sesión, recibe un token que debe incluir en las peticiones protegidas.

#### Estructura del token:

```json
{
  "correo": "usuario@example.com",
  "nombre": "Juan Pérez",
  "id": 1,
  "iat": 1700000000,
  "exp": 1700086400
}
```

- **iat** (Issued At): Fecha de emisión del token
- **exp** (Expiration): Fecha de expiración (24 horas después)

#### Uso del token:

```http
GET /api/v1/recetas
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### 🔒 Rutas Protegidas

Las siguientes rutas requieren autenticación JWT:

- ✅ POST/PUT/DELETE de **Categorías**
- ✅ POST/PUT/DELETE de **Recetas**
- ✅ GET `/recetas-helpers/usuarios/:id`

---

## 🛡️ Middleware

### ValidarJWTMiddleware

Middleware que valida el token JWT en las peticiones protegidas.

**Validaciones:**
1. Verifica que el header `Authorization` exista
2. Valida el formato `Bearer <token>`
3. Verifica la firma del token con la clave secreta
4. Valida que el token no haya expirado
5. Verifica que el usuario del token exista en la base de datos

**Ejemplo de uso en rutas:**

```go
router.POST("/categorias", middleware.ValidarJWTMiddleware, rutas.Categoria_post)
```

---

## 🔧 Utilidades

### 📧 EnviarCorreo

Función para enviar correos electrónicos usando SMTP (configurado para Gmail).

**Uso:**

```go
import "backend/utilidades"

err := utilidades.EnviarCorreo(
    "destino@example.com",
    "Asunto del correo",
    "<h1>Mensaje HTML</h1><p>Contenido del correo</p>"
)
```

**Casos de uso:**
- ✉️ Verificación de cuenta al registrarse
- ✉️ Mensajes de contacto
- ✉️ Recuperación de contraseña (futuro)

---

## 🐳 Docker

### Levantar MySQL con Docker Compose

```bash
docker-compose up -d
```

### Detener el contenedor

```bash
docker-compose down
```

### Ver logs del contenedor

```bash
docker-compose logs -f mysql
```

### Acceder al contenedor MySQL

```bash
docker exec -it mysql8 mysql -uroot -p
# Contraseña: rami123
```

---

## 📜 Scripts Útiles

### Ejecutar la aplicación

```bash
go run main.go
```

### Ejecutar con hot-reload (usando fresh)

Instala fresh:
```bash
go install github.com/gravityblast/fresh@latest
```

Ejecuta:
```bash
fresh
```

### Compilar el proyecto

```bash
go build -o backend.exe
```

### Ejecutar tests

```bash
go test ./...
```

### Formatear código

```bash
go fmt ./...
```

### Instalar dependencias faltantes

```bash
go mod tidy
```

---

## 🧪 Probar la API

### Con curl (PowerShell):

```powershell
# GET - Obtener todas las recetas
curl http://localhost:8081/api/v1/recetas

# POST - Crear categoría (requiere JWT)
curl -X POST http://localhost:8081/api/v1/categorias `
  -H "Authorization: Bearer TU_TOKEN_JWT" `
  -H "Content-Type: application/json" `
  -d '{"nombre":"Postres"}'

# GET - Buscar recetas
curl "http://localhost:8081/api/v1/recetas-helpers/buscador?categoria_id=1&search=chocolate"
```

### Con Postman:

1. Importa la colección de Postman (si existe en el proyecto)
2. Configura la variable `{{baseUrl}}` = `http://localhost:8081/api/v1`
3. Para rutas protegidas, agrega el header:
   - Key: `Authorization`
   - Value: `Bearer <tu_token_jwt>`

---

## 🎯 Flujo de Registro y Login

### 1. Registro de Usuario

```
POST /api/v1/seguridad/registro
→ Se crea usuario con estado "Inactivo"
→ Se genera token UUID único
→ Se envía correo de verificación
```

### 2. Verificación de Cuenta

```
Usuario hace clic en el enlace del correo
GET /api/v1/seguridad/verificacion/{token}
→ Se activa la cuenta (estado = "Activo")
→ Se puede iniciar sesión
```

### 3. Inicio de Sesión

```
POST /api/v1/seguridad/login
→ Valida correo y contraseña
→ Verifica que la cuenta esté activa
→ Genera y devuelve JWT con expiración de 24h
```

### 4. Uso del JWT

```
El usuario incluye el token en las peticiones protegidas
Authorization: Bearer <token>
→ El middleware valida el token
→ Permite acceso a los recursos
```

---

## 📈 Características Avanzadas

### 🔄 Soft Delete

Todos los modelos tienen `DeletedAt` para eliminación lógica:

```go
DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
```

Cuando se elimina un registro, no se borra de la BD, solo se marca como eliminado.

### 🔍 Búsqueda y Filtros

El endpoint `/recetas-helpers/buscador` permite:

- Filtrar por `categoria_id`
- Buscar por texto en nombre/descripción (`search`)
- Devuelve resultados formateados con datos de categoría y usuario

### 📸 Subida de Archivos

- Las imágenes se guardan en `public/recetas/` y `public/uploads/fotos/`
- Se generan nombres únicos con timestamp: `upload_1234567890.jpg`
- Se validan extensiones de archivo
- Se pueden servir como archivos estáticos: `http://localhost:8081/public/recetas/imagen.jpg`

### 🔐 Hash de Contraseñas

Las contraseñas se hashean con **bcrypt** antes de guardarse:

```go
import "golang.org/x/crypto/bcrypt"

hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
```

### 🌐 CORS

Configurado para permitir peticiones desde cualquier origen (ajustar en producción):

```go
Access-Control-Allow-Origin: *
Access-Control-Allow-Headers: Content-Type, Authorization
Access-Control-Allow-Methods: POST, OPTIONS, GET, PUT
```

---

## 🚨 Manejo de Errores

El proyecto tiene manejo centralizado de errores:

### Códigos de estado HTTP:

- `200 OK` - Operación exitosa
- `201 Created` - Recurso creado
- `400 Bad Request` - Error en los datos enviados
- `401 Unauthorized` - No autenticado o token inválido
- `404 Not Found` - Recurso no encontrado
- `500 Internal Server Error` - Error del servidor

### Formato de respuestas de error:

```json
{
  "estado": "error",
  "mensaje": "Ocurrió un error inesperado",
  "error": "Detalles específicos del error"
}
```

---

## 🔮 Mejoras Futuras

- [ ] Paginación en listados de recetas
- [ ] Recuperación de contraseña
- [ ] Subir múltiples fotos por receta
- [ ] Sistema de favoritos
- [ ] Comentarios y calificaciones
- [ ] Roles de usuario (admin, usuario normal)
- [ ] Rate limiting en endpoints públicos
- [ ] Caché con Redis
- [ ] Tests unitarios y de integración
- [ ] Documentación con Swagger/OpenAPI
- [ ] CI/CD con GitHub Actions
- [ ] Deploy en producción (AWS, Heroku, etc.)

---

## 📝 Notas Importantes

### Seguridad:

- ⚠️ **NO subas el archivo `.env` a Git** (ya está en `.gitignore`)
- ⚠️ Usa contraseñas seguras para `SECRET_JWT` en producción
- ⚠️ Configura CORS adecuadamente en producción
- ⚠️ Usa HTTPS en producción

### Base de Datos:

- Las migraciones se ejecutan automáticamente al iniciar
- Asegúrate de tener la base de datos `go_fullstack` creada
- Inserta los estados manualmente después de la primera ejecución

### Correos:

- Si usas Gmail, activa la verificación en 2 pasos
- Genera una contraseña de aplicación específica
- Verifica que no bloquee el acceso de aplicaciones menos seguras

---

## 🤝 Contribuir

Si deseas contribuir al proyecto:

1. Haz fork del repositorio
2. Crea una rama para tu feature (`git checkout -b feature/nueva-funcionalidad`)
3. Commitea tus cambios (`git commit -m 'Agrega nueva funcionalidad'`)
4. Push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Abre un Pull Request

---

## 📄 Licencia

Este proyecto está bajo la licencia MIT. Ver el archivo `LICENSE` para más detalles.

---

## 👨‍💻 Autor

**Ramiro Nuñez Perez**
**Telefono: +51961501468**

---

## 📞 Soporte

Si tienes preguntas o problemas:

1. Revisa este README
2. Verifica las variables de entorno en `.env`
3. Revisa los logs de la aplicación
4. Abre un issue en GitHub

---

## 🎉 ¡Gracias por usar esta API!

Si este proyecto te fue útil, considera darle una ⭐ en GitHub.

---

**Última actualización:** Noviembre 2025
**Versión:** 1.0.0

