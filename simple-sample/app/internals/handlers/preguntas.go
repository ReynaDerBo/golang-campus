package handlers

import (
	"app/internals/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetPreguntas(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, pregunta, respuesta FROM preguntas")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var preguntas []models.Pregunta
		for rows.Next() {
			var p models.Pregunta
			if err := rows.Scan(&p.ID, &p.Pregunta, &p.Respuesta); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			preguntas = append(preguntas, p)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(preguntas)
	}
}

func CreatePregunta(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.Pregunta
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "Solicitud inválida", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("INSERT INTO preguntas (pregunta, respuesta) VALUES ($1, $2)", p.Pregunta, p.Respuesta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
