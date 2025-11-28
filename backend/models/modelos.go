package models

import (
	"backend/database"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Categoria struct {
	ID        uint           `json:"id"`
	Nombre    string         `gorm:"type:varchar(100);not null" json:"nombre"`
	Slug      string         `gorm:"type:varchar(100);not null" json:"slug"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Categorias []Categoria

type Receta struct {
	ID          uint           `json:"id"`
	CategoriaID uint           `json:"categoria_id"`
	UsuarioID   uint           `json:"usuario_id"`
	Usuario     *Usuario       `gorm:"foreignKey:UsuarioID;references:ID" json:"usuario"`
	Categoria   *Categoria     `gorm:"foreignKey:CategoriaID;references:ID" json:"categoria"`
	Nombre      string         `gorm:"type:varchar(100);not null" json:"nombre"`
	Slug        string         `gorm:"type:varchar(100);not null" json:"slug"`
	Tiempo      string         `gorm:"type:varchar(100);not null" json:"tiempo"`
	Foto        string         `gorm:"type:varchar(100);not null" json:"foto"`
	Descripcion string         `json:"descripcion"`
	Fecha       time.Time      `json:"fecha"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Recetas []Receta

type Contacto struct {
	Id       uint      `json:"id"`
	Nombre   string    `gorm:"type:varchar(100)" json:"nombre"`
	Correo   string    `gorm:"type:varchar(100)" json:"correo"`
	Telefono string    `gorm:"type:varchar(100)" json:"telefono"`
	Mensaje  string    `json:"mensaje"`
	Fecha    time.Time `json:"fecha"`
}

type Contactos []Contacto

type Estado struct {
	ID     uint   `json:"id"`
	Nombre string `gorm:"type:varchar(100);not null" json:"nombre"`
}
type Estados []Estado

type Usuario struct {
	ID       uint      `json:"id"`
	EstadoID uint      `json:"estado_id"`
	Estado   *Estado   `gorm:"foreignKey:EstadoID;references:ID" json:"estado"`
	Nombre   string    `gorm:"type:varchar(100);not null" json:"nombre"`
	Correo   string    `gorm:"type:varchar(100);not null" json:"correo"`
	Password string    `gorm:"type:varchar(100);not null" json:"password"`
	Token    string    `gorm:"type:varchar(100);not null" json:"token"`
	Fecha    time.Time `json:"fecha"`
}
type Usuarios []Usuario

func Migraciones() {
	err := database.Database.AutoMigrate(&Categoria{}, &Receta{}, &Contacto{}, &Estado{}, &Usuario{})
	if err != nil {
		panic("Error en migración de Categoria, Receta, Contacto, Estado, Usuario: " + err.Error())
	}
	fmt.Println("Migración de Categoria, Receta, Contacto, Estado, Usuario, ejecutada correctamente")
}
