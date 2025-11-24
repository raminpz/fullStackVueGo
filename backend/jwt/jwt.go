package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerarJWT(correo string, nombre string, id uint) (string, error) {
	miClave := []byte(os.Getenv("SECRET_JWT"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"correo": correo,
		"nombre": nombre,
		"id":     id,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(miClave)
	return tokenString, err
}
