package bd

import (
	"github.com/JoaoPaulo87/microblog-server/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de login a la BD*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}
	// Ahora comparo la password que en la BD está encriptada
	// creo una variable que sea un slice de bytes
	// el slice de bytes es porque esta encriptada con letras y nros
	passwordBytes := []byte(password)
	// creo otra variable con la password que tengo en la BD para el usuario
	passwordBD := []byte(usu.Password)
	// Ahora llamo a una función del package bcrypt que compara las password
	// No podemos hacer if passwordBytes == passwordBD porque estan encriptadas
	// por eso utilizamos este metodo
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
