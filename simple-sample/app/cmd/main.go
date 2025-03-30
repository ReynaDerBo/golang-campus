package main

import (
	"fmt"
	"log"
	"net/http"

	"app/internals/database"
	"app/internals/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/preguntas", handlers.GetPreguntas(db)).Methods("GET")
	r.HandleFunc("/preguntas", handlers.CreatePregunta(db)).Methods("POST")

	fmt.Println("Servidor corriendo en el puerto 8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}
