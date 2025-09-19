package main

import (
	"backend/database"
	"backend/models"
	"backend/rutas"

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

	// Ejecucion
	router.Run() // listen and serve on 0.0.0.0:8080
}
