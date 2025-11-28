# 📝 Changelog

Todos los cambios notables en este proyecto serán documentados en este archivo.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/es-ES/1.0.0/),
y este proyecto adhiere a [Versionado Semántico](https://semver.org/lang/es/).

---

## [1.0.0] - 2025-11-27

### ✨ Agregado

#### Documentación Completa
- **README.md**: Documentación principal con instalación, configuración y uso
- **CONTRIBUTING.md**: Guía para contribuidores y desarrolladores
- **API_DOCS.md**: Documentación detallada de todos los endpoints
- **scripts.sql**: Scripts SQL útiles para gestión de BD y datos de prueba
- **.env.example**: Archivo de ejemplo con todas las variables de entorno
- **.gitignore**: Exclusiones de archivos para Git
- **CHANGELOG.md**: Este archivo

#### Comentarios en el Código
- Documentación completa en `main.go` explicando cada sección
- Comentarios descriptivos en configuración CORS
- Comentarios en todas las rutas de la API
- Documentación del flujo de autenticación

#### Configuración
- Docker Compose documentado para MySQL 8
- Variables de entorno organizadas y comentadas
- Configuración de CORS mejorada

#### Funcionalidades del Proyecto

**Autenticación y Seguridad**
- Sistema de registro de usuarios con verificación por email
- Login con JWT (expiración de 24 horas)
- Middleware de validación de tokens
- Hash de contraseñas con bcrypt
- Tokens de verificación con UUID

**CRUD de Categorías**
- Listar todas las categorías
- Obtener categoría por ID
- Crear categoría (requiere JWT)
- Actualizar categoría (requiere JWT)
- Eliminar categoría - soft delete (requiere JWT)
- Generación automática de slugs

**CRUD de Recetas**
- Listar todas las recetas con datos relacionados
- Obtener receta por ID con formato personalizado
- Crear receta con validaciones (requiere JWT)
- Actualizar receta (requiere JWT)
- Eliminar receta con borrado de foto (requiere JWT)
- Relaciones con categorías y usuarios
- Formateo de fechas personalizado (dd/mm/yyyy)

**Búsqueda y Filtros**
- Buscar recetas por slug
- Filtrar recetas por categoría
- Búsqueda por texto en nombre/descripción
- Obtener recetas de un usuario específico
- Endpoint para página principal

**Subida de Archivos**
- Subir fotos de recetas
- Validación de extensiones
- Nombres únicos con timestamp
- Servir archivos estáticos desde `/public`

**Contacto**
- Formulario de contacto con validaciones
- Guardado en base de datos
- Envío de correos SMTP

**Base de Datos**
- Migraciones automáticas con GORM
- Soft delete en todos los modelos
- Relaciones entre tablas (Foreign Keys)
- Pool de conexiones optimizado
- Logger personalizado de consultas

### 🔧 Configuración Técnica

**Dependencias principales**
- Gin v1.10.1 - Framework web
- GORM v1.31.0 - ORM
- MySQL Driver v1.6.0
- JWT v5.3.0
- Gomail v2 - Envío de correos
- Bcrypt - Hash de contraseñas
- UUID v1.6.0 - Tokens únicos
- Godotenv v1.5.1 - Variables de entorno
- Slug v1.15.0 - URLs amigables

**Modelos de Datos**
- Usuario (con estado activo/inactivo)
- Categoría (con soft delete)
- Receta (con relaciones a categoría y usuario)
- Contacto (mensajes de formulario)
- Estado (activo/inactivo para usuarios)

**Características Técnicas**
- CORS configurado para desarrollo
- Manejo centralizado de errores 404
- Logger de Gin con timestamps
- Recovery middleware para panics
- Validaciones con tags de GORM
- DTOs para validación de entrada

### 📦 Estructura del Proyecto
```
backend/
├── database/      # Conexión a BD
├── dto/          # Data Transfer Objects
├── jwt/          # Generación de tokens
├── middleware/   # Middleware de autenticación
├── models/       # Modelos GORM
├── public/       # Archivos estáticos
├── rutas/        # Handlers de endpoints
├── utilidades/   # Funciones auxiliares
├── validaciones/ # Validaciones personalizadas
└── main.go       # Punto de entrada
```

---

## [0.1.0] - Versión Inicial

### Agregado
- Estructura base del proyecto
- Configuración inicial de Gin y GORM
- Endpoints de ejemplo para testing

---

## Notas de Versiones

### Sobre el Versionado

- **MAJOR** (X.0.0): Cambios incompatibles con versiones anteriores
- **MINOR** (0.X.0): Nuevas funcionalidades compatibles con versiones anteriores
- **PATCH** (0.0.X): Correcciones de bugs

### Categorías de Cambios

- **Agregado**: Para funcionalidades nuevas
- **Cambiado**: Para cambios en funcionalidades existentes
- **Obsoleto**: Para funcionalidades que serán eliminadas
- **Eliminado**: Para funcionalidades eliminadas
- **Corregido**: Para correcciones de bugs
- **Seguridad**: Para actualizaciones de seguridad

---

## [Próximas Versiones]

### [1.1.0] - Planificado

**Mejoras**
- [ ] Paginación en listados de recetas
- [ ] Rate limiting en endpoints públicos
- [ ] Recuperación de contraseña
- [ ] Sistema de favoritos
- [ ] Roles de usuario (admin, usuario normal)

**Optimizaciones**
- [ ] Caché con Redis para consultas frecuentes
- [ ] Compresión de respuestas JSON
- [ ] Optimización de queries con índices

**Testing**
- [ ] Tests unitarios para handlers
- [ ] Tests de integración para endpoints
- [ ] Cobertura mínima del 80%

**Documentación**
- [ ] Documentación con Swagger/OpenAPI
- [ ] Colección de Postman exportada
- [ ] Video tutorial de uso

**DevOps**
- [ ] CI/CD con GitHub Actions
- [ ] Dockerfile para la aplicación
- [ ] Scripts de deploy automatizado

---

## Mantener este Archivo

Cuando hagas cambios significativos:

1. Agrega una nueva entrada en la sección correspondiente
2. Usa la fecha actual en formato ISO (YYYY-MM-DD)
3. Categoriza los cambios apropiadamente
4. Sé descriptivo pero conciso
5. Actualiza la versión en `go.mod` si aplica

---

**Última actualización:** 27 de noviembre de 2025

