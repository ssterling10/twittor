package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors" //permisos para que el api se pueda acceder de cualquier lugar
)

/* Manjeadores: seteo mi puerto, el handler y pongo a escuchar el servidor */
func Manejadores() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
