package middleware

import (
	"backend/database"
	"backend/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func ValidarJWTMiddleware(c *gin.Context) {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado.",
		})
		return
	}
	// Obtenemos el secret jwt
	miClave := []byte(os.Getenv("SECRET_JWT"))
	// Obtenemos el encabezado
	var haader = c.GetHeader("Authorization")
	// Validamos si el header tiene valor
	if len(haader) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"estado":         "error",
			"mensaje":        "No autorizado",
			"estadoOpcional": "Error porque no viene el token",
		})
		return
	}
	// Separamos el valor del header
	splitBearer := strings.Split(haader, " ")
	if len(splitBearer) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"estado":         "error",
			"mensaje":        "No autorizado",
			"estadoOpcional": "Error porque el token trae un solo texto",
		})
		return
	}
	// Validamos el formato del token
	tk := strings.TrimSpace(splitBearer[1])
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"estado":         "error",
				"mensaje":        "No autorizado",
				"estadoOpcional": "Error con la firma del token",
			})
		}
		return miClave, nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"estado":         "error",
			"mensaje":        "No autorizado",
			"estadoOpcional": "Error con la firma del token",
		})
		return
	}
	// Obtenemos los datos del JWT
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// preguntamos si existe correo
		datos := models.Usuario{}

		result := database.Database.First(&datos, claims["id"])
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"estado":         "error",
				"mensaje":        "No autorizado",
				"estadoOpcional": "Error con el id del usuario informado en el token",
			})
			return
		}
		c.Next()

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"estado":         "error",
			"mensaje":        "No autorizado",
			"estadoOpcional": "Error tratando de leer los datos",
		})
		return
	}

}
