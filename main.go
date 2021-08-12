package main

import (
	"log"

	"github.com/ssterling10/twittor/bd"
	"github.com/ssterling10/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return

	handlers.Manejadores()

}
