package dto

type EjemploDto struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

type CategoriaDto struct {
	Nombre string `json:"nombre" binding:"required"`
}
