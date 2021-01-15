package bd

import (
	"context"
	"time"

	"github.com/JoaoPaulo87/microblogginjoao/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un mail de parámetro y chequea si ya
está en la BD*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email} //Le pasamos un json
	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
