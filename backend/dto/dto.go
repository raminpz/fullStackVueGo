package dto

type EjemploDto struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

type CategoriaDto struct {
	Nombre string `json:"nombre" binding:"required"`
}

type RecetaDto struct {
	Nombre      string `json:"nombre" binding:"required"`
	Tiempo      string `json:"tiempo" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	CategoriaId uint   `json:"categoria_id"`
}

// response
type RecetaResponse struct {
	Id          uint   `json:"id"`
	Nombre      string `json:"nombre" binding:"required"`
	Slug        string `json:"slug"`
	CategoriaId uint   `json:"categoria_id"`
	Categoria   string `json:"categoria"`
	Tiempo      string `json:"tiempo"`
	Foto        string `json:"foto"`
	Descripcion string `json:"descripcion"`
	Fecha       string `json:"fecha"`
}

type RecetasResponses []RecetaResponse

type ContactanosDto struct {
	Nombre   string `json:"nombre" binding:"required"`
	Correo   string `json:"correo" binding:"required,email"` // Ojo no pongamos espacio después del required si no, fallará
	Telefono string `json:"telefono" binding:"required"`
	Mensaje  string `json:"mensaje" binding:"required"`
}

type UsuarioDto struct {
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

type LoginDto struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}
