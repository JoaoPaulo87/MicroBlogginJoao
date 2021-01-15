package bd

import (
	"context"
	"time"

	"github.com/JoaoPaulo87/microblogginjoao/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es la parada final con la BD para insertar datos del usuario.*/

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	/* context creo que es para establecer coneccion con la BD de mongo. Si pasan 15 seg. va
	a hacer un cancel*/
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("usuarios")
	u.Password, _ = EncriptarPassword(u.Password)
	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
