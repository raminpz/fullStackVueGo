package rutas

import (
	"backend/database"
	"backend/dto"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

func Categoria_get(c *gin.Context) {
	var datos []models.Categoria
	result := database.Database.Find(&datos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"estado":  "error",
			"mensaje": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"estado": "ok",
		"datos":  datos,
	})
}

func Categoria_getId(c *gin.Context) {
	id := c.Param("id")
	datos := models.Categoria{}
	result := database.Database.First(&datos, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
			"error":   result.Error.Error(), //Opcional
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"estado": "ok",
		"datos":  datos, // Devuelve toda la información del registro
	})
}

func Categoria_post(c *gin.Context) {
	var body dto.CategoriaDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "Oocurrio un error inesperado",
			"error":   err.Error(),
		})
		return
	}
	// Validamos que no existe el registro
	var existe models.Categoria
	result := database.Database.Where("nombre = ?", body.Nombre).First(&existe)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "Ya existe un registro con ese nombre: " + body.Nombre,
		})
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"estado":  "error",
			"mensaje": "Error al consultar la base de datos",
			"error":   result.Error.Error(),
		})
		return
	}
	// Creamos el registro
	datos := models.Categoria{Nombre: body.Nombre, Slug: slug.Make(body.Nombre)}
	database.Database.Save(&datos)
	c.JSON(http.StatusCreated, gin.H{
		"estado":  "ok",
		"mensaje": "Registro creado correctamente",
	})
}

func Categoria_put(c *gin.Context) {
	var body dto.CategoriaDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "Oocurrio un error inesperado",
			"error":   err.Error(),
		})
		return
	}
	// Obtener el id path
	id := c.Param("id")
	// Validamos que exista la categoria por id
	datos := models.Categoria{}
	if err := database.Database.First(&datos, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
			"error":   err.Error(),
		})
		return
	}
	// Actualizamos el registro
	datos.Nombre = body.Nombre
	datos.Slug = slug.Make(body.Nombre)
	database.Database.Save(&datos)
	// Retornamos el registro actualizado
	c.JSON(http.StatusOK, gin.H{
		"estado":  "ok",
		"mensaje": "Registro actualizado correctamente",
	})

}

func Categoria_delete(c *gin.Context) {
	id := c.Param("id")
	datos := models.Categoria{}
	result := database.Database.First(&datos, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
			"error":   result.Error.Error(),
		})
		return
	}
	// Validamos que exista la receta por categoria_id
	categoria_id, _ := strconv.ParseUint(id, 10, 64)
	existe := models.Recetas{}
	database.Database.Where(&models.Receta{CategoriaID: uint(categoria_id)}).Find(&existe)
	if len(existe) >= 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "No se puede eliminar la categoria porque tiene recetas asociadas",
		})
		return
	}

	// Eliminamos el registro
	database.Database.Delete(&datos)
	// Retornamos
	c.JSON(http.StatusOK, gin.H{
		"estado":  "ok",
		"mensaje": "Registro eliminado correctamente",
	})
}
