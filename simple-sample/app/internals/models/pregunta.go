package models

type Pregunta struct {
	ID        int    `json:"id"`
	Pregunta  string `json:"pregunta"`
	Respuesta string `json:"respuesta"`
}
