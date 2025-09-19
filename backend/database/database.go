package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Conectar() {
	// Cargamos las variables de entorno
	err := godotenv.Load()
	if err != nil {
		panic("Error cargando variables de entorno: " + err.Error())
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_SERVER") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	// Solo mostrar el nombre de la base de datos para debug seguro
	fmt.Println("Conectando a la base de datos:", os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error al conectar a la base de datos")
		panic(err)
	}
	fmt.Println("Conexión exitosa a la base de datos")
	Database = db
}
