package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func Conectar() {
	// Cargamos las variables de entorno
	err := godotenv.Load()
	if err != nil {
		panic("Error cargando variables de entorno: " + err.Error())
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_SERVER") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println("Conectando a la base de datos:", os.Getenv("DB_NAME"))

	// Configurar logger personalizado: solo muestra consultas lentas (>= 500ms)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             500 * time.Millisecond, // Aumentar umbral de 200ms a 500ms
			LogLevel:                  logger.Warn,            // Solo muestra warnings y errores
			IgnoreRecordNotFoundError: true,                   // Ignora errores de registro no encontrado
			Colorful:                  true,                   // Logs con colores
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		fmt.Println("Error al conectar a la base de datos")
		panic(err)
	}

	// Configurar pool de conexiones para mejor rendimiento
	sqlDB, err := db.DB()
	if err != nil {
		panic("Error configurando pool de conexiones: " + err.Error())
	}

	// Configuración del pool de conexiones
	sqlDB.SetMaxIdleConns(10)           // Máximo de conexiones inactivas
	sqlDB.SetMaxOpenConns(100)          // Máximo de conexiones abiertas
	sqlDB.SetConnMaxLifetime(time.Hour) // Tiempo máximo de vida de una conexión

	fmt.Println("Conexión exitosa a la base de datos")
	Database = db
}
