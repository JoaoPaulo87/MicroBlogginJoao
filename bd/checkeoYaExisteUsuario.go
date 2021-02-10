package bd

import (
	"context"
	"time"

	"github.com/JoaoPaulo87/microblog-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un mail de parámetro y chequea si ya
está en la BD*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	//ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*El contexto es como un tiempo de espera o un plazo que indica cuándo una operación debe dejar de funcionar y reanudarse.
	Ayuda a prevenir la degradación del rendimiento en los sistemas de producción cuando determinadas operaciones se ejecutan con lentitud.
	En este código, está pasando context.TODO() para indicar que no está seguro de qué contexto usar en este momento, pero que planea añadir uno en el futuro.*/
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")

	// M es una función que formatea o mapea a bson lo que recibe como json
	condicion := bson.M{"email": email} //Le pasamos un json

	// en la variable resultado voy a modelar un usuario
	var resultado models.Usuario

	//FindOne me devuelve un sólo registro que cumple con la condición
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	//Convertimos el ID en hexagesimal a string con la función Hex
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
