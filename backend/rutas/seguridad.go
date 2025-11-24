package rutas

import (
	"backend/database"
	"backend/dto"
	"backend/jwt"
	"backend/models"
	"backend/utilidades"
	"backend/validaciones"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Seguridad_registro(c *gin.Context) {
	var body dto.UsuarioDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Datos de entrada inválidos. Verifique que todos los campos requeridos estén completos y sean correctos.",
			"errorOpcional": err.Error(),
		})
		return
	}
	if len(body.Nombre) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error al registrar el usuario.",
			"errorOpcional": "El campo nombre es obligatorio",
		})
		return
	}
	if len(body.Correo) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error al registrar el usuario.",
			"errorOpcional": "El campo correo es obligatorio",
		})
		return
	}
	if validaciones.Regex_correo.FindStringSubmatch(body.Correo) == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error al registrar el usuario.",
			"errorOpcional": "El correo ingresado no es válido.",
		})
		return
	}
	if len(body.Password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error al registrar el usuario.",
			"errorOpcional": "El campo password es obligatorio",
		})
		return
	}

	// Validamos que el correo no exista en la tabla usuario
	existe := models.Usuarios{}
	result := database.Database.Where(&models.Usuario{Correo: body.Correo}).First(&existe)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error al registrar el usuario.",
			"errorOpcional": "El correo " + body.Correo + " ya está registrado.",
		})
		return
	}
	// Generamos el hash con bcrypt
	costo := 10
	bytes, _ := bcrypt.GenerateFromPassword([]byte(body.Password), costo)
	// Generamos un token
	token := uuid.New()

	// TODO: Implementar registro de usuario
	save := models.Usuario{
		Nombre:   body.Nombre,
		Correo:   body.Correo,
		Token:    token.String(),
		Password: string(bytes),
		EstadoID: 2,
		Fecha:    time.Now(),
	}
	database.Database.Save(&save)
	// Enviar mail de verificacion
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	url := scheme + "://" + c.Request.Host + "/api/v1/seguridad/verificacion/" + token.String()
	var mensaje = "<h1>Verificación de cuenta</h1>" +
		"Hola " + body.Nombre + ",<br><br>" +
		"Para verificar su cuenta haga click en el siguiente enlace:<br>" +
		"<a href='" + url + "'>" + url + "</a><br><br>" +
		"O copia y pega el siguiente enlace en tu navegador:<br>" +
		url + "<br><br>" +
		"Gracias por registrarse."
	utilidades.EnviarCorreo(body.Correo, "Verificación de cuenta - "+body.Nombre, mensaje)

	// retornamos
	c.JSON(http.StatusCreated, gin.H{
		"estado":  "ok",
		"mensaje": "Registro creado correctamente",
	})

}

func Seguridad_verificacion(c *gin.Context) {
	token := c.Param("token")

	// Validar token no vacío antes de consultar
	if len(token) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Token de verificación inválido.",
			"errorOpcional": "El token no puede estar vacío.",
		})
		return
	}

	// Buscar el usuario con el token (estado 2 = pendiente)
	user := models.Usuario{}
	res := database.Database.Where(&models.Usuario{Token: token, EstadoID: 2}).First(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		// No encontramos usuario con ese token
		c.JSON(http.StatusNotFound, gin.H{
			"estado":        "error",
			"mensaje":       "Token de verificación inválido o expirado.",
			"errorOpcional": "No se encontró un usuario pendiente con ese token.",
		})
		return
	}

	// Modificamos el registro
	user.Token = ""
	user.EstadoID = 1
	if err := database.Database.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error al verificar el usuario.",
			"errorOpcional": err.Error(),
		})
		return
	}

	// retornamos
	c.Redirect(http.StatusMovedPermanently, os.Getenv("RUTA_FRONTEND"))

}

func Seguridad_login(c *gin.Context) {
	var body dto.LoginDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error inesperado.",
			"errorOpcional": err.Error(),
		})
		return
	}
	if len(body.Correo) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error al registrar el usuario.",
			"errorOpcional": "El campo correo es obligatorio",
		})
		return
	}
	if validaciones.Regex_correo.FindStringSubmatch(body.Correo) == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error al registrar el usuario.",
			"errorOpcional": "El correo ingresado no es válido.",
		})
		return
	}
	if len(body.Password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error al registrar el usuario.",
			"errorOpcional": "El campo password es obligatorio",
		})
		return
	}
	// Validamos que el correo no exista en la tabla usuario
	usuario := models.Usuarios{}
	database.Database.Where(&models.Usuario{Correo: body.Correo, EstadoID: 1}).Find(&usuario)
	if len(usuario) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error inesperado.",
			"errorOpcional": "Las credenciales no son válidas | El usuario se encontró pero no está activo.",
		})
		return
	}
	// Usamos bcrypt para comparar los password
	passwordByte := []byte(body.Password)
	passwordBD := []byte(usuario[0].Password)
	errorPassword := bcrypt.CompareHashAndPassword(passwordBD, passwordByte)
	if errorPassword != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":        "error",
			"mensaje":       "Ocurrió un error al registrar el usuario.",
			"errorOpcional": "Las credenciales no son válidas | El password ingresado ingresado no es válido.",
		})
		return

	} else {
		jwtKey, errJWT := jwt.GenerarJWT(usuario[0].Correo, usuario[0].Nombre, usuario[0].ID)
		if errJWT != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"estado":        "error",
				"mensaje":       "Ocurrió un error inesperado.",
				"errorOpcional": "Ocurrió un error al intentar generar el token" + errJWT.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":     usuario[0].ID,
			"nombre": usuario[0].Nombre,
			"token":  jwtKey,
		})
	}
}
