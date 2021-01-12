package main

import (
	"log"

	"github.com/JoaoPaulo87/microblogginjoao/bd"
	"github.com/JoaoPaulo87/microblogginjoao/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
