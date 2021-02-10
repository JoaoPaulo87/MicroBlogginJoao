package main

import (
	"log"

	"github.com/JoaoPaulo87/microblog-server/bd"
	"github.com/JoaoPaulo87/microblog-server/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
