package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutina que me permite encriptar la password
recibida*/
func EncriptarPassword(pass string) (string, error) {
	/*El algoritmo de encriptacion hará (2 elevado al costo) pasados
	por el texto*/
	costo := 8
	//me retorna la password encriptada y un err si dio error
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
