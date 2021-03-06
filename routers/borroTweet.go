package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JoaoPaulo87/microblog-server/bd"
)

/*BorroTweet - Borra un tweet*/
func BorroTweet(w http.ResponseWriter, r *http.Request) {
	// Aca obtengo el id del tweet que tengo como parametro en el request del postman
	tweetID := r.URL.Query().Get("tweetID")

	if len(tweetID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	// IDUsuario es una variable global que esta en procesoToken.go para asi no tener que definir siempre lo mismo
	respuesta, borrado := bd.BorrarTweet(tweetID, IDUsuario)

	if borrado == false {
		http.Error(w, "error al borrar el Tweet", http.StatusBadRequest)
		return
	}

	/* En la sig línea le decimos al navegador que si va respuesta será del tipo json
	   aunque no le mandamos respuesta, es una buena práctica, por si más adelante queremos enviarle una */
	w.Header().Set("Content-Type", "application/json")
	// y le damos un status created:
	w.WriteHeader(http.StatusCreated)

	// Aca le devolvemos una respuesta al navegador a su peticion de borrar tweet
	json.NewEncoder(w).Encode(respuesta)
}
