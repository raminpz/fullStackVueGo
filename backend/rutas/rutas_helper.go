package rutas

import (
	"backend/database"
	"backend/dto"
	"backend/models"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Receta_Helper_Home(c *gin.Context) {
	// Hacemos la consulta a la base de datos
	var recetas []models.Receta
	result := database.Database.Preload("Categoria").Preload("Usuario").Limit(3).Find(&recetas)
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

func Receta_Helper_Usuario(c *gin.Context) {
	// Obtenemos y formateamos el id
	usuario_id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || usuario_id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "ID de usuario inválido",
		})
		return
	}

	// Validamos que el usuario exista
	usuario := models.Usuario{}
	result := database.Database.First(&usuario, usuario_id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
			"error":   "El usuario especificado no existe",
		})
		return
	}

	// Hacemos la consulta a la base de datos
	var recetas []models.Receta
	result = database.Database.Where(&models.Receta{UsuarioID: uint(usuario_id)}).Preload("Categoria").Preload("Usuario").Find(&recetas)
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

func Receta_Helper_Slug(c *gin.Context) {
	// Obtenemos el slug del parámetro
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "El slug es obligatorio",
		})
		return
	}

	// Buscamos la receta por slug
	var receta models.Receta
	result := database.Database.Where("slug = ?", slug).Preload("Categoria").Preload("Usuario").First(&receta)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
			"error":   "El slug ingresado no existe",
		})
		return
	}

	// Obtenemos la URL del proyecto
	schema := "http"
	if c.Request.TLS != nil {
		schema = "https"
	}
	baseURL := schema + "://" + c.Request.Host

	// Formateamos la fecha
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

	// Construimos la respuesta
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

func Receta_Helper_Buscador(c *gin.Context) {
	categoria_id, _ := strconv.ParseUint(c.Query("categoria_id"), 10, 64)
	searchTerm := c.Query("search")

	// Validamos que la categoría exista solo si se especificó una
	if categoria_id > 0 {
		cat := models.Categoria{}
		if errorCategoria := database.Database.First(&cat, categoria_id); errorCategoria.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"estado":  "error",
				"mensaje": "La categoría especificada no existe",
			})
			return
		}
	}

	// Construimos la consulta base
	query := database.Database.Model(&models.Receta{})

	// Aplicamos filtro de búsqueda por nombre si existe
	if searchTerm != "" {
		query = query.Where("nombre LIKE ?", "%"+searchTerm+"%")
	}

	// Aplicamos filtro de categoría solo si se especificó
	if categoria_id > 0 {
		query = query.Where("categoria_id = ?", categoria_id)
	}

	// Ejecutamos la consulta con los Preloads
	var recetas []models.Receta
	result := query.Preload("Categoria").Preload("Usuario").Find(&recetas)
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

func Receta_Helper_Editar_Foto(c *gin.Context) {
	// Validamos que venga el archivo
	file, err := c.FormFile("foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
			"error":   "No se recibió la foto",
		})
		return
	}

	// Validar mimetype de foto
	contentType := file.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
			"error":   "El archivo debe ser JPG o PNG",
		})
		return
	}

	// Validamos los campos
	recetaIdStr := strings.TrimSpace(c.PostForm("receta_id"))
	if recetaIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "El campo receta_id es obligatorio",
		})
		return
	}

	// Convertimos el receta_id a uint
	recetaId, err := strconv.ParseUint(recetaIdStr, 10, 64)
	if err != nil || recetaId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "El receta_id debe ser un número válido",
		})
		return
	}

	// Buscamos la receta en la base de datos
	var receta models.Receta
	result := database.Database.First(&receta, recetaId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"estado":  "error",
			"mensaje": "La receta especificada no existe",
		})
		return
	}

	// Obtenemos la extensión del archivo
	nombrePartes := strings.Split(file.Filename, ".")
	extension := ""
	if len(nombrePartes) > 1 {
		extension = nombrePartes[len(nombrePartes)-1]
	}

	// Generamos un nombre único para la foto usando timestamp
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	nuevoNombreFoto := "upload_" + timestamp + "." + extension

	// Guardamos el archivo en la carpeta de recetas
	rutaArchivo := "public/recetas/" + nuevoNombreFoto
	if err := c.SaveUploadedFile(file, rutaArchivo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"estado":  "error",
			"mensaje": "No se pudo guardar el archivo",
		})
		return
	}

	// Eliminamos la foto anterior si existe
	if receta.Foto != "" {
		_ = os.Remove("public/recetas/" + receta.Foto)
	}

	// Actualizamos la foto en la base de datos
	receta.Foto = nuevoNombreFoto
	if err := database.Database.Save(&receta).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"estado":  "error",
			"mensaje": "No se pudo actualizar la receta",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"estado":  "ok",
		"mensaje": "Foto actualizada correctamente",
		"foto":    nuevoNombreFoto,
	})
}
