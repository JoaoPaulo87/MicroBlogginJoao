package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JoaoPaulo87/microblog-server/bd"
	"github.com/JoaoPaulo87/microblog-server/jwt"
	"github.com/JoaoPaulo87/microblog-server/models"
)

/*Login realiza el login de usuario
Recibe como parámetro lo mismo de todos los endpoints y no devuelve nada,
como los otros endPoints, son prácticamente métodos*/
func Login(w http.ResponseWriter, r *http.Request) {
	// Vamos a setear en el header que el contenido que devolveremos (w)
	// será de tipo Json
	w.Header().Add("Content-Type", "application/json")

	var usu models.Usuario

	// Aca usamos Decode para insertar en el struct Usuario al usuario con
	// sus datos en formato json. Ej. Nombre: Juan, Apellido: Perez, etc.
	err := json.NewDecoder(r.Body).Decode(&usu)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(usu.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}
	// documento es un usuario, es el primer paramtro que me devuelve ItentoLogin.
	documento, existe := bd.IntentoLogin(usu.Email, usu.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña inválidos ", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente"+err.Error(), 400)
		return
	}
	// Si llegamos hasta acá es porque viene todo bien.

	// Los tokens son permisos (a groso modo) que tienen un tiempo de expiración.
	// Si ese tiempo se acaba, hay que generar otro.
	// Si el token está generado:
	resp := models.RespuestaLogin{
		// jwtKey es el token de respuesta generado, es la parte codificada que vemos en
		// la pagina web jtw.io como esa secuencia de letras y nros.
		Token: jwtKey,
	}
	/* La respuesta generada la tenemos que poner en el header. Una vez que desde el frontend
	el usuario puso en un formulario su usuario y pass nos devuelve algo, y en lo que devuelve,
	lo devuelve con el token generado. Nos lo devuelve creado entonces ponemos un StatusCreated
	y codificamos la respuesta. */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// Vamos también a grabar una cookie
	// generamos un campo fecha para ver la expiración de esa cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
