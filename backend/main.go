package main

import (
	"backend/database"
	"backend/middleware"
	"backend/models"
	"backend/rutas"
	"os"

	"github.com/gin-gonic/gin"
)

var pathh = "/api/v1/"

func main() {
	// SetMode
	gin.SetMode(gin.ReleaseMode)
	// Inicializa el router de Gin con middleware predeterminado (Logger y Recovery)
	router := gin.Default()

	// Aplica middleware CORS para permitir peticiones desde el frontend
	router.Use(corsMiddleware())

	// Establece conexión con la base de datos MySQL
	database.Conectar()

	// Ejecuta las migraciones automáticas de GORM (crea tablas si no existen)
	models.Migraciones()

	// Configura la carpeta 'public' para servir archivos estáticos (imágenes, etc.)
	// Accesible en: http://localhost:PORT/public/...
	router.Static("/public", "./public")

	// Manejo personalizado para rutas no encontradas (404)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"Estado":  "Error",
			"Mensaje": "Recurso no disponible",
		})
	})

	// Ruta raíz - Página de bienvenida
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Estado":  "Ok",
			"Mensaje": "Bienvenido a mi API Golang con Gin Framework con GORM",
		})
	})

	// ==================== RUTAS DE EJEMPLO/PRUEBA ====================
	// Endpoints para probar funcionalidades básicas de la API	// ==================== RUTAS DE EJEMPLO/PRUEBA ====================
	// Endpoints para probar funcionalidades básicas de la API

	router.GET(pathh+"ejemplo", rutas.Ejemplo_get)
	router.POST(pathh+"ejemplo", rutas.Ejemplo_post)
	router.PUT(pathh+"ejemplo/:id", rutas.Ejemplo_put)
	router.DELETE(pathh+"ejemplo/:id", rutas.Ejemplo_delete)
	router.GET(pathh+"ejemplo/:id", rutas.Ejemplo_get_parametros)
	router.GET(pathh+"ejemplo-querystring", rutas.Ejemplo_querystring)
	router.POST(pathh+"upload", rutas.Ejemplo_upload) // Subida de archivos de prueba

	// ==================== RUTAS DE CATEGORÍAS ====================
	// CRUD completo de categorías de recetas
	// Rutas protegidas: POST, PUT, DELETE requieren JWT

	router.GET(pathh+"categorias", rutas.Categoria_get)                                            // Obtener todas
	router.GET(pathh+"categorias/:id", rutas.Categoria_getId)                                      // Obtener por ID
	router.POST(pathh+"categorias", middleware.ValidarJWTMiddleware, rutas.Categoria_post)         // Crear (requiere JWT)
	router.PUT(pathh+"categorias/:id", middleware.ValidarJWTMiddleware, rutas.Categoria_put)       // Actualizar (requiere JWT)
	router.DELETE(pathh+"categorias/:id", middleware.ValidarJWTMiddleware, rutas.Categoria_delete) // Eliminar (requiere JWT)

	// ==================== RUTAS DE RECETAS ====================
	// CRUD completo de recetas de cocina
	// Rutas protegidas: POST, PUT, DELETE requieren JWT

	router.GET(pathh+"recetas", rutas.Receta_get)                                            // Obtener todas las recetas
	router.GET(pathh+"recetas/:id", rutas.Receta_getId)                                      // Obtener receta por ID
	router.POST(pathh+"recetas", middleware.ValidarJWTMiddleware, rutas.Receta_post)         // Crear receta (requiere JWT)
	router.PUT(pathh+"recetas/:id", middleware.ValidarJWTMiddleware, rutas.Receta_put)       // Actualizar receta (requiere JWT)
	router.DELETE(pathh+"recetas/:id", middleware.ValidarJWTMiddleware, rutas.Receta_delete) // Eliminar receta (requiere JWT)

	// ==================== RUTA DE CONTACTO ====================
	// Endpoint público para enviar mensajes de contacto

	router.POST(pathh+"contactanos", rutas.Contactanos) // Envía correo y guarda en BD

	// ==================== RUTAS DE SEGURIDAD/AUTENTICACIÓN ====================
	// Endpoints para registro, verificación y login de usuarios

	router.POST(pathh+"seguridad/registro", rutas.Seguridad_registro)               // Registrar nuevo usuario
	router.GET(pathh+"seguridad/verificacion/:token", rutas.Seguridad_verificacion) // Verificar cuenta por email
	router.POST(pathh+"seguridad/login", rutas.Seguridad_login)                     // Iniciar sesión (devuelve JWT)

	// ==================== RUTAS AUXILIARES (HELPERS) ====================
	// Endpoints especializados para búsqueda, filtros y operaciones específicas

	router.GET(pathh+"recetas-helpers/usuarios/:id", middleware.ValidarJWTMiddleware, rutas.Receta_Helper_Usuario) // Recetas de un usuario (requiere JWT)
	router.GET(pathh+"recetas-helpers/home", rutas.Receta_Helper_Home)                                             // Recetas para página principal
	router.GET(pathh+"recetas-helpers/slug/:slug", rutas.Receta_Helper_Slug)                                       // Obtener receta por slug (URL amigable)
	router.GET(pathh+"recetas-helpers/buscador", rutas.Receta_Helper_Buscador)                                     // Buscar recetas con filtros
	router.POST(pathh+"recetas-helpers/foto", rutas.Receta_Helper_Editar_Foto)                                     // Subir foto de receta

	// ==================== INICIAR SERVIDOR ====================
	// ==================== INICIAR SERVIDOR ====================
	// Obtiene el puerto desde variable de entorno y levanta el servidor

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		panic("Error al iniciar el servidor: " + err.Error())
	}
}

// corsMiddleware configura los headers CORS para permitir peticiones desde el frontend
// CORS (Cross-Origin Resource Sharing) permite que el frontend en otro dominio/puerto
// pueda hacer peticiones a esta API
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Permite peticiones desde cualquier origen (⚠️ en producción, especificar dominios concretos)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Headers permitidos en las peticiones
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Métodos HTTP permitidos
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		// Manejo de peticiones preflight (OPTIONS)
		// Los navegadores envían OPTIONS antes de peticiones con headers personalizados
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // 204 No Content
			return
		}

		// Continúa con el siguiente middleware/handler
		c.Next()
	}
}
