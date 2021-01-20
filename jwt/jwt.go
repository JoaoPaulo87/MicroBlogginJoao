package jwt

import (
	"time"

	"github.com/JoaoPaulo87/microblogginjoao/models"
	jwt "github.com/dgrijalva/jwt-go" //Creo un alias para manejarlo más fácil
)

func GeneroJWT(usu models.Usuario) (string, error) {
	// esta es la clave secreta de la pag. web jwt.io, es algo unico asi cada jwt es diferente
	miClave := []byte("SkillFactoryGo_Avalith")

	// Armamos el jwt con los datos del usuario que viene como parametro, usu.
	payload := jwt.MapClaims{
		"email":            usu.Email,
		"nombre":           usu.Nombre,
		"apellidos":        usu.Apellidos,
		"fecha_nacimiento": usu.FechaNacimiento,
		"biografia":        usu.Biografia,
		"ubicacion":        usu.Ubicacion,
		"sitioWeb":         usu.SitioWeb,
		"_id":              usu.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	// Creamos el token con el algoritmos HS256 (u otro que querramos) y con los datos (payload)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// Aca le agregamos la firma o clave secreta al jwt.
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
