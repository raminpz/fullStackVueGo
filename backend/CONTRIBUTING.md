# 🤝 Guía de Contribución

¡Gracias por tu interés en contribuir a este proyecto! Esta guía te ayudará a configurar el entorno de desarrollo y entender el flujo de trabajo.

---

## 📋 Requisitos del Desarrollador

- **Go** 1.24 o superior
- **MySQL** 8.0+ o Docker
- **Git**
- Editor de código (GoLand, VS Code, etc.)
- **Postman** o **Thunder Client** para probar la API

---

## 🔧 Configuración del Entorno de Desarrollo

### 1. Clonar el Repositorio

```bash
git clone <URL_DEL_REPOSITORIO>
cd fullStackVueGo/backend
```

### 2. Instalar Dependencias

```bash
go mod download
```

### 3. Configurar Variables de Entorno

```bash
cp .env.example .env
# Edita .env con tus configuraciones
```

### 4. Levantar Base de Datos

Con Docker:
```bash
docker-compose up -d
```

O instala MySQL 8 localmente y crea la base de datos:
```sql
CREATE DATABASE go_fullstack;
```

### 5. Ejecutar la Aplicación

Opción 1 - Sin hot-reload:
```bash
go run main.go
```

Opción 2 - Con hot-reload (recomendado):
```bash
# Instalar fresh
go install github.com/gravityblast/fresh@latest

# Ejecutar
fresh
```

---

## 📝 Convenciones de Código

### Nomenclatura

- **Paquetes**: lowercase sin guiones bajos (`database`, `models`)
- **Archivos**: snake_case (`rutas_helper.go`)
- **Funciones públicas**: PascalCase (`Categoria_get`, `EnviarCorreo`)
- **Funciones privadas**: camelCase (`corsMiddleware`)
- **Variables**: camelCase (`miClave`, `smtpHost`)
- **Constantes**: UPPER_CASE o PascalCase (`SECRET_JWT`, `pathh`)
- **Structs**: PascalCase (`Usuario`, `Receta`)
- **Campos de struct**: PascalCase (`Nombre`, `CategoriaID`)

### Formato de Código

Usa `gofmt` para formatear el código:

```bash
go fmt ./...
```

### Comentarios

- Documenta todas las funciones públicas
- Usa comentarios en español (como en el resto del proyecto)
- Los comentarios de documentación comienzan con el nombre de la función/tipo

Ejemplo:
```go
// EnviarCorreo envía un email a través de SMTP
// Parámetros:
//   - correo: dirección de email del destinatario
//   - asunto: asunto del mensaje
//   - mensaje: contenido HTML del email
// Retorna error si el envío falla
func EnviarCorreo(correo, asunto, mensaje string) error {
    // ...
}
```

---

## 🏗️ Estructura de una Ruta Nueva

### 1. Definir el Modelo (si aplica)

En `models/modelos.go`:

```go
type MiModelo struct {
    ID        uint           `json:"id"`
    Nombre    string         `gorm:"type:varchar(100);not null" json:"nombre"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Agregar a la función Migraciones()
err := database.Database.AutoMigrate(&MiModelo{})
```

### 2. Crear el DTO (si aplica)

En `dto/dto.go`:

```go
type MiModeloDto struct {
    Nombre string `json:"nombre" binding:"required"`
}
```

### 3. Implementar el Handler

Crea un archivo en `rutas/` (ejemplo: `rutas/mi_modelo.go`):

```go
package rutas

import (
    "backend/database"
    "backend/dto"
    "backend/models"
    "net/http"
    
    "github.com/gin-gonic/gin"
)

// MiModelo_get obtiene todos los registros
func MiModelo_get(c *gin.Context) {
    var datos []models.MiModelo
    result := database.Database.Find(&datos)
    
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "estado":  "error",
            "mensaje": "Error al obtener los datos",
        })
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "estado": "ok",
        "datos":  datos,
    })
}

// MiModelo_post crea un nuevo registro
func MiModelo_post(c *gin.Context) {
    var body dto.MiModeloDto
    
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "estado":  "error",
            "mensaje": "Datos inválidos",
            "error":   err.Error(),
        })
        return
    }
    
    datos := models.MiModelo{
        Nombre: body.Nombre,
    }
    
    result := database.Database.Create(&datos)
    
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "estado":  "error",
            "mensaje": "Error al crear el registro",
        })
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{
        "estado": "ok",
        "datos":  datos,
    })
}
```

### 4. Registrar la Ruta

En `main.go`:

```go
// Rutas públicas
router.GET(pathh+"mi-modelo", rutas.MiModelo_get)
router.GET(pathh+"mi-modelo/:id", rutas.MiModelo_getId)

// Rutas protegidas (requieren JWT)
router.POST(pathh+"mi-modelo", middleware.ValidarJWTMiddleware, rutas.MiModelo_post)
router.PUT(pathh+"mi-modelo/:id", middleware.ValidarJWTMiddleware, rutas.MiModelo_put)
router.DELETE(pathh+"mi-modelo/:id", middleware.ValidarJWTMiddleware, rutas.MiModelo_delete)
```

---

## 🧪 Testing

### Ejecutar Tests

```bash
# Todos los tests
go test ./...

# Con cobertura
go test -cover ./...

# Generar reporte HTML de cobertura
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Estructura de un Test

Crea archivos `*_test.go` junto a tus archivos de código:

```go
package rutas

import (
    "testing"
    "net/http/httptest"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestMiModelo_get(t *testing.T) {
    // Configurar
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.GET("/test", MiModelo_get)
    
    // Ejecutar
    req := httptest.NewRequest("GET", "/test", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    // Verificar
    assert.Equal(t, 200, w.Code)
}
```

---

## 📊 Base de Datos

### Migraciones

El proyecto usa GORM AutoMigrate. Cuando agregas un nuevo modelo:

1. Define el struct en `models/modelos.go`
2. Agrégalo a la función `Migraciones()`
3. Al iniciar la app, la tabla se creará automáticamente

### Seeds (Datos de Prueba)

Para insertar datos de prueba, puedes:

1. Ejecutar SQL directamente en MySQL
2. Crear una función de seeding en `models/`
3. Usar fixtures en los tests

Ejemplo de función de seeding:

```go
// En models/modelos.go
func SeedEstados() {
    estados := []Estado{
        {Nombre: "Activo"},
        {Nombre: "Inactivo"},
    }
    
    for _, estado := range estados {
        database.Database.FirstOrCreate(&estado, Estado{Nombre: estado.Nombre})
    }
}
```

---

## 🔍 Debugging

### Logs

Usa `fmt.Println()` o `log.Println()` para debug:

```go
import "log"

log.Printf("Debug: valor = %v", miVariable)
```

### GORM Logger

Para ver las queries SQL, ajusta el nivel de log en `database/database.go`:

```go
// Muestra todas las queries (modo debug)
LogLevel: logger.Info,

// Solo muestra queries lentas y errores (modo producción)
LogLevel: logger.Warn,
```

### Postman

Importa la colección de Postman (si existe) o crea tus propias requests:

1. Crear request
2. Agregar header `Authorization: Bearer <token>` para rutas protegidas
3. Guardar en una colección
4. Exportar y versionar la colección

---

## 🚀 Flujo de Trabajo con Git

### Branches

- `main` - Rama principal (producción)
- `develop` - Rama de desarrollo
- `feature/nombre-feature` - Para nuevas funcionalidades
- `fix/nombre-bug` - Para correcciones de bugs

### Workflow

1. Crear una rama desde `develop`:

```bash
git checkout develop
git pull origin develop
git checkout -b feature/mi-nueva-funcionalidad
```

2. Hacer cambios y commits:

```bash
git add .
git commit -m "feat: agregar endpoint de búsqueda de recetas"
```

3. Push y crear Pull Request:

```bash
git push origin feature/mi-nueva-funcionalidad
# Ir a GitHub/GitLab y crear un PR hacia develop
```

### Mensajes de Commit

Usa convenciones de [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` - Nueva funcionalidad
- `fix:` - Corrección de bug
- `docs:` - Cambios en documentación
- `refactor:` - Refactorización de código
- `test:` - Agregar o modificar tests
- `chore:` - Tareas de mantenimiento

Ejemplos:
```
feat: agregar endpoint de categorías
fix: corregir validación de email en registro
docs: actualizar README con ejemplos de uso
refactor: simplificar lógica de autenticación JWT
```

---

## 🐛 Reporte de Bugs

Cuando reportes un bug, incluye:

1. **Descripción clara** del problema
2. **Pasos para reproducir** el error
3. **Comportamiento esperado** vs **comportamiento actual**
4. **Logs de error** (si aplica)
5. **Versión de Go** y **dependencias** (`go version`, `cat go.mod`)
6. **Sistema operativo**

---

## ✅ Checklist antes de un Pull Request

- [ ] El código compila sin errores (`go build`)
- [ ] Todos los tests pasan (`go test ./...`)
- [ ] El código está formateado (`go fmt ./...`)
- [ ] Agregaste comentarios donde sea necesario
- [ ] Actualizaste el README si agregaste nuevas funcionalidades
- [ ] Probaste los endpoints en Postman
- [ ] No subiste archivos sensibles (`.env`, credenciales, etc.)

---

## 📚 Recursos Útiles

- [Documentación oficial de Go](https://golang.org/doc/)
- [Documentación de Gin](https://gin-gonic.com/docs/)
- [Documentación de GORM](https://gorm.io/docs/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)

---

## 💬 Comunicación

Si tienes dudas:

1. Revisa la documentación del README
2. Busca en issues cerrados si alguien ya preguntó lo mismo
3. Abre un issue con tus preguntas
4. Pregunta en el canal de Slack/Discord del equipo (si aplica)

---

¡Gracias por contribuir! 🎉

