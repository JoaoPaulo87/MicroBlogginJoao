package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JoaoPaulo87/microblogginjoao/middlew"
	"github.com/JoaoPaulo87/microblogginjoao/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores: seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()

	// Por cada EndPoint vamos a tener un renglon que permita manejar la función correspondiente
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
