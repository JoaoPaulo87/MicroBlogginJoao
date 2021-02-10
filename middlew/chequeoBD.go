package middlew

import (
	"net/http"

	"github.com/JoaoPaulo87/microblog-server/bd"
)

/*ChequeoBD es el middleware que me permite conocer el estado de la BD */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		// En esta linea es como decir "si hay conección con la BD, continuá"
		next.ServeHTTP(w, r)
	}
}
