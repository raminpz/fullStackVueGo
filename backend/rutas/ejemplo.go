package rutas

// RUTAS DE TIPO HANDLER

import (
	"backend/dto"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Ejemplo_get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Estado":  "Ok",
		"Mensaje": "Método GET",
	})
}

func Ejemplo_get_parametros(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Estado":  "Ok",
		"Mensaje": "Método GET Parametros | id=" + c.Param("id"),
	})
}

func Ejemplo_post(c *gin.Context) {
	var body dto.EjemploDto
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Estado":  "Error",
			"Mensaje": "JSON inválido",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Estado":   "Ok",
		"Mensaje":  "Método POST",
		"correo":   body.Correo,
		"password": body.Password,
	})
}
func Ejemplo_put(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Estado":  "Ok",
		"Mensaje": "Método PUT | id=" + c.Param("id"),
	})
}
func Ejemplo_delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"Estado":  "Ok",
		"Mensaje": "Método DELETE | id=" + c.Param("id"),
	})
}

func Ejemplo_querystring(c *gin.Context) {
	c.JSON(200, gin.H{
		"Estado":  "Ok",
		"Mensaje": "Método GET querystring | id=" + c.Query("id") + " | slug=" + c.Query("slug"),
	})
}

func Ejemplo_upload(c *gin.Context) {
	file, err := c.FormFile("foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Estado":  "Error",
			"Mensaje": "No se recibió ningún archivo o el parámetro es incorrecto",
		})
		return
	}

	// Validar nombre y extensión
	nombre := file.Filename
	if nombre == "" || strings.HasPrefix(nombre, ".") {
		c.JSON(http.StatusBadRequest, gin.H{
			"Estado":  "Error",
			"Mensaje": "El archivo no tiene un nombre válido",
		})
		return
	}

	extension := ""
	if punto := strings.LastIndex(nombre, "."); punto != -1 && punto < len(nombre)-1 {
		extension = nombre[punto+1:]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Estado":  "Error",
			"Mensaje": "El archivo no tiene extensión",
		})
		return
	}

	// Generar nombre único seguro
	foto := fmt.Sprintf("upload_%d.%s", time.Now().UnixNano(), extension)

	// Guardar archivo en la ruta fija (no temporal)
	archivo := "public/uploads/fotos/" + foto
	err = c.SaveUploadedFile(file, archivo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Estado":  "Error",
			"Mensaje": "No se pudo guardar el archivo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Estado":  "Ok",
		"Mensaje": "Archivo subido correctamente",
		"archivo": foto,
	})
}
