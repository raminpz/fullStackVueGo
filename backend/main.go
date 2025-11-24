package main

import (
	"backend/database"
	"backend/models"
	"backend/rutas"
	"os"

	"github.com/gin-gonic/gin"
)

var pathh = "/api/v1/"

func main() {
	// SetMode
	gin.SetMode(gin.ReleaseMode)
	// Inicializacion Gin
	router := gin.Default()

	// Conectar a la base de datos
	database.Conectar()

	// Ejecutar migraciones
	models.Migraciones()

	// archivos estaticos
	router.Static("/public", "./public")

	// Custom error 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"Estado":  "Error",
			"Mensaje": "Recurso no disponible",
		})
	})

	// Página principal
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Estado":  "Ok",
			"Mensaje": "Bienvenido a mi API Golang con Gin Framework con GORM",
		})
	})

	// Rutas

	router.GET(pathh+"ejemplo", rutas.Ejemplo_get)
	router.POST(pathh+"ejemplo", rutas.Ejemplo_post)
	router.PUT(pathh+"ejemplo/:id", rutas.Ejemplo_put)
	router.DELETE(pathh+"ejemplo/:id", rutas.Ejemplo_delete)
	router.GET(pathh+"ejemplo/:id", rutas.Ejemplo_get_parametros)
	router.GET(pathh+"ejemplo-querystring", rutas.Ejemplo_querystring)
	router.POST(pathh+"upload", rutas.Ejemplo_upload)

	router.GET(pathh+"categorias", rutas.Categoria_get)
	router.GET(pathh+"categorias/:id", rutas.Categoria_getId)
	router.POST(pathh+"categorias", rutas.Categoria_post)
	router.PUT(pathh+"categorias/:id", rutas.Categoria_put)
	router.DELETE(pathh+"categorias/:id", rutas.Categoria_delete)

	router.GET(pathh+"recetas", rutas.Receta_get)
	router.GET(pathh+"recetas/:id", rutas.Receta_getId)
	router.POST(pathh+"recetas", rutas.Receta_post)
	router.PUT(pathh+"recetas/:id", rutas.Receta_put)
	router.DELETE(pathh+"recetas/:id", rutas.Receta_delete)

	router.POST(pathh+"contactanos", rutas.Contactanos)

	router.POST(pathh+"seguridad/registro", rutas.Seguridad_registro)
	router.GET(pathh+"seguridad/verificacion/:token", rutas.Seguridad_verificacion)

	router.POST(pathh+"seguridad/login", rutas.Seguridad_login)

	// Ejecucion
	router.Run(":" + os.Getenv("PORT")) // listen and serve on 0.0.0.0:8081
}
