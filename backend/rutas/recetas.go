package rutas

import (
	"backend/database"
	"backend/dto"
	"backend/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func Receta_get(c *gin.Context) {
	var recetas []models.Receta
	result := database.Database.Preload("Categoria").Preload("Usuario").Find(&recetas)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"estado":  "error",
			"mensaje": result.Error.Error(),
		})
		return
	}

	respuestas := make(dto.RecetasResponses, 0, len(recetas))
	// Construye la URL base una sola vez
	schema := "http"
	if c.Request.TLS != nil {
		schema = "https"
	}
	baseURL := schema + "://" + c.Request.Host

	for _, r := range recetas {
		fecha := ""
		if !r.Fecha.IsZero() {
			fecha = r.Fecha.Format("02/01/2006")
		}

		// Validación segura de relaciones que pueden ser nil
		categoriaNombre := ""
		if r.Categoria != nil {
			categoriaNombre = r.Categoria.Nombre
		}

		usuarioNombre := ""
		if r.Usuario != nil {
			usuarioNombre = r.Usuario.Nombre
		}

		respuestas = append(respuestas, dto.RecetaResponse{
			Id:          r.ID,
			Nombre:      r.Nombre,
			Slug:        r.Slug,
			CategoriaId: r.CategoriaID,
			Categoria:   categoriaNombre,
			UsuarioId:   r.UsuarioID,
			Usuario:     usuarioNombre,
			Tiempo:      r.Tiempo,
			Foto:        baseURL + "/public/recetas/" + r.Foto,
			Descripcion: r.Descripcion,
			Fecha:       fecha,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"estado": "ok",
		"datos":  respuestas,
	})
}

func Receta_getId(c *gin.Context) {
	id := c.Param("id")
	var receta models.Receta
	result := database.Database.Preload("Categoria").Preload("Usuario").First(&receta, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
			"error":   result.Error.Error(),
		})
		return
	}

	schema := "http"
	if c.Request.TLS != nil {
		schema = "https"
	}
	baseURL := schema + "://" + c.Request.Host

	fecha := ""
	if !receta.Fecha.IsZero() {
		fecha = receta.Fecha.Format("02/01/2006")
	}

	// Validación segura de relaciones que pueden ser nil
	categoriaNombre := ""
	if receta.Categoria != nil {
		categoriaNombre = receta.Categoria.Nombre
	}

	usuarioNombre := ""
	if receta.Usuario != nil {
		usuarioNombre = receta.Usuario.Nombre
	}

	respuesta := dto.RecetaResponse{
		Id:          receta.ID,
		Nombre:      receta.Nombre,
		Slug:        receta.Slug,
		CategoriaId: receta.CategoriaID,
		Categoria:   categoriaNombre,
		UsuarioId:   receta.UsuarioID,
		Usuario:     usuarioNombre,
		Tiempo:      receta.Tiempo,
		Foto:        baseURL + "/public/recetas/" + receta.Foto,
		Descripcion: receta.Descripcion,
		Fecha:       fecha,
	}

	c.JSON(http.StatusOK, gin.H{
		"estado": "ok",
		"datos":  respuesta,
	})
}

// validateRecetaForm valida los campos del formulario y devuelve el mapa de errores
// y un objeto models.Receta con los campos ya parseados (sin Foto, Slug ni Fecha).
func validateRecetaForm(c *gin.Context) (map[string][]string, models.Receta) {
	errorValidacion := map[string][]string{}
	const (
		maxNombre      = 50
		maxTiempo      = 50
		maxDescripcion = 1000
	)

	nombre := strings.TrimSpace(c.PostForm("nombre"))
	categoriaStr := strings.TrimSpace(c.PostForm("categoria_id"))
	usuarioStr := strings.TrimSpace(c.PostForm("usuario_id"))
	tiempo := strings.TrimSpace(c.PostForm("tiempo"))
	descripcion := strings.TrimSpace(c.PostForm("descripcion"))

	// nombre: obligatorio y máximo de caracteres
	if nombre == "" {
		errorValidacion["nombre"] = append(errorValidacion["nombre"], "El campo nombre es obligatorio")
	} else if utf8.RuneCountInString(nombre) > maxNombre {
		errorValidacion["nombre"] = append(errorValidacion["nombre"], fmt.Sprintf("El nombre no debe tener más de %d caracteres", maxNombre))
	}

	// categoria_id: obligatorio y numérico (>0)
	var categoriaID uint64
	if categoriaStr == "" {
		errorValidacion["categoria_id"] = append(errorValidacion["categoria_id"], "El campo categoria_id es obligatorio")
	} else {
		if val, err := strconv.ParseUint(categoriaStr, 10, 64); err != nil || val == 0 {
			errorValidacion["categoria_id"] = append(errorValidacion["categoria_id"], "categoria_id debe ser un número entero positivo válido")
		} else {
			categoriaID = val
		}
	}

	// usuario_id: obligatorio y numérico (>0)
	var usuarioID uint64
	if usuarioStr == "" {
		errorValidacion["usuario_id"] = append(errorValidacion["usuario_id"], "El campo usuario_id es obligatorio")
	} else {
		if val, err := strconv.ParseUint(usuarioStr, 10, 64); err != nil || val == 0 {
			errorValidacion["usuario_id"] = append(errorValidacion["usuario_id"], "usuario_id debe ser un número entero positivo válido")
		} else {
			usuarioID = val
		}
	}

	// tiempo: obligatorio y límite de caracteres
	if tiempo == "" {
		errorValidacion["tiempo"] = append(errorValidacion["tiempo"], "El campo tiempo es obligatorio")
	} else if utf8.RuneCountInString(tiempo) > maxTiempo {
		errorValidacion["tiempo"] = append(errorValidacion["tiempo"], fmt.Sprintf("El campo tiempo no debe exceder %d caracteres", maxTiempo))
	}

	// descripcion: obligatorio y límite de caracteres
	if descripcion == "" {
		errorValidacion["descripcion"] = append(errorValidacion["descripcion"], "El campo descripcion es obligatorio")
	} else if utf8.RuneCountInString(descripcion) > maxDescripcion {
		errorValidacion["descripcion"] = append(errorValidacion["descripcion"], fmt.Sprintf("El campo descripcion no debe exceder %d caracteres", maxDescripcion))
	}

	receta := models.Receta{
		CategoriaID: uint(categoriaID),
		UsuarioID:   uint(usuarioID),
		Nombre:      nombre,
		Tiempo:      tiempo,
		Descripcion: descripcion,
	}

	return errorValidacion, receta
}

func Receta_post(c *gin.Context) {
	file, err := c.FormFile("foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":         "error",
			"mensaje":        "Ocurrío un error inesperado",
			"estadoOpcional": "No viene la foto",
		})
		return
	}
	// validar mimetype de foto
	if file.Header["Content-Type"][0] == "image/jpeg" || file.Header["Content-Type"][0] == "image/png" {

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":         "error",
			"mensaje":        "Ocurrío un error inesperado",
			"estadoOpcional": "El archivo de la foto no es compatible, debe ser JPG o PNG",
		})
		return
	}

	// uso de la función de validación extraída
	errorValidacion, recetaVal := validateRecetaForm(c)
	if len(errorValidacion) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": errorValidacion,
		})
		return
	}
	// Validamos que exista la categoria por id
	catExiste := models.Categoria{}
	if err := database.Database.First(&catExiste, recetaVal.CategoriaID); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
			"error":   "La categoría especificada no existe",
		})
		return
	}

	// Validamos que exista el usuario por id
	usuarioExiste := models.Usuario{}
	if err := database.Database.First(&usuarioExiste, recetaVal.UsuarioID); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
			"error":   "El usuario especificado no existe",
		})
		return
	}

	// Validamos que no existe el registro (usando el nombre validado)
	var existe models.Recetas
	database.Database.Where(&models.Receta{Nombre: recetaVal.Nombre}).Find(&existe)
	if len(existe) >= 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "Ya existe un registro con ese nombre: " + recetaVal.Nombre,
		})
		return
	}
	// nombramos el archivo (generamos un nombre único y seguro)
	parts := strings.Split(file.Filename, ".")
	extension := "jpg"
	if len(parts) > 1 {
		extension = strings.ToLower(parts[len(parts)-1])
	}
	foto := fmt.Sprintf("%d.%s", time.Now().UnixNano(), extension)
	archivo := "public/recetas/" + foto
	if err := c.SaveUploadedFile(file, archivo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"estado":  "error",
			"mensaje": "No se pudo guardar la foto",
			"error":   err.Error(),
		})
		return
	}

	// Creamos el registro con los valores ya validados y parseados
	receta := models.Receta{
		CategoriaID: recetaVal.CategoriaID,
		UsuarioID:   recetaVal.UsuarioID,
		Nombre:      recetaVal.Nombre,
		Slug:        slug.Make(recetaVal.Nombre),
		Tiempo:      recetaVal.Tiempo,
		Foto:        foto,
		Descripcion: recetaVal.Descripcion,
		Fecha:       time.Now(),
	}
	database.Database.Save(&receta)

	// retornamos
	c.JSON(http.StatusCreated, gin.H{
		"estado":  "Ok",
		"mensaje": "Registro creado correctamente",
	})
}

func Receta_put(c *gin.Context) {
	// Obtenemos datos del json request
	var body dto.RecetaDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "Ocurrio un error inesperado",
			"error":   err.Error(),
		})
		return
	}

	// Obtenemos el id de la receta a actualizar
	id := c.Param("id")

	// Validamos que exista la receta por id
	var receta models.Receta
	result := database.Database.First(&receta, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
			"error":   result.Error.Error(),
		})
		return
	}

	// Actualizamos solo los campos necesarios con Updates (no toca created_at)
	updates := map[string]interface{}{
		"nombre":       body.Nombre,
		"slug":         slug.Make(body.Nombre),
		"tiempo":       body.Tiempo,
		"descripcion":  body.Descripcion,
		"categoria_id": body.CategoriaId,
	}

	if err := database.Database.Model(&receta).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"estado":  "error",
			"mensaje": "No se pudo actualizar el registro",
			"error":   err.Error(),
		})
		return
	}

	// retornamos
	c.JSON(http.StatusOK, gin.H{
		"estado":  "Ok",
		"mensaje": "Registro actualizado correctamente",
	})

}

func Receta_delete(c *gin.Context) {
	// Obtenemos el id de la receta a eliminar
	id := c.Param("id")
	// Validamos que exista la receta por id
	dato := models.Receta{}
	result := database.Database.First(&dato, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
			"error":   result.Error.Error(),
		})
		return
	}
	// Borramos la foto
	borrar := "public/recetas/" + dato.Foto
	e := os.Remove(borrar)
	if e != nil {
		log.Fatal(e)
	}
	// Eliminamos registro
	database.Database.Delete(&dato)
	// retornamos
	c.JSON(http.StatusOK, gin.H{
		"estado":  "Ok",
		"mensaje": "Registro eliminado correctamente",
	})
}
