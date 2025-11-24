package rutas

import (
	"backend/database"
	"backend/dto"
	"backend/models"
	"backend/utilidades"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Contactanos(c *gin.Context) {
	var body dto.ContactanosDto
	if err := c.ShouldBindJSON(&body); err != nil {
		// Error de validación o parsing JSON
		log.Println("Contactanos - bind error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "error",
			"mensaje": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Creamos el registro
	datos := models.Contacto{
		Nombre:   body.Nombre,
		Correo:   body.Correo,
		Telefono: body.Telefono,
		Mensaje:  body.Mensaje,
		Fecha:    time.Now(),
	}

	// Insertar (Create para nueva fila)
	result := database.Database.Create(&datos)
	if result.Error != nil {
		log.Println("Contactanos - db error:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"estado":  "error",
			"mensaje": "No se pudo guardar el contacto",
			"error":   result.Error.Error(),
		})
		return
	}
	// Enviamos el correo de notificación al administrador
	emailAdmin := os.Getenv("SMTP_TO_EMAIL")
	if emailAdmin == "" {
		emailAdmin = "admin@example.com" // Email por defecto
	}

	var mensaje = "<h1>Nuevo mensaje de contacto</h1>" +
		"<p>Se ha recibido un nuevo mensaje desde el formulario de contacto:</p>" +
		"<ul>" +
		"<li><strong>Nombre:</strong> " + body.Nombre + "</li>" +
		"<li><strong>Email:</strong> " + body.Correo + "</li>" +
		"<li><strong>Teléfono:</strong> " + body.Telefono + "</li>" +
		"<li><strong>Mensaje:</strong> " + body.Mensaje + "</li>" +
		"</ul>"

	if err := utilidades.EnviarCorreo(emailAdmin, "Nuevo mensaje de contacto - "+body.Nombre, mensaje); err != nil {
		log.Println("Error al enviar correo:", err)
		// El contacto se guardó correctamente, pero el correo falló
		// No devolvemos error para no confundir al usuario
		c.JSON(http.StatusOK, gin.H{
			"estado":  "ok",
			"mensaje": "Se creó el registro exitosamente (nota: no se pudo enviar notificación por correo)",
			"data":    datos,
		})
		return
	}

	// Retornamos
	c.JSON(http.StatusOK, gin.H{
		"estado":  "ok",
		"mensaje": "Se creó el registro exitosamente y se envió la notificación",
		"data":    datos,
	})
}
