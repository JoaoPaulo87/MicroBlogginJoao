package bd

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func BorrarTweet(tweetID string, idUsuario string) (*mongo.DeleteResult, bool) {

	var borrado bool
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblogging")
	col := db.Collection("tweet")

	// convierto el string ID que viene como parámetro en hexadecimal a un ObjID
	objID, errObj := primitive.ObjectIDFromHex(tweetID)

	if errObj != nil {
		log.Fatal(errObj)
	}

	condicion := bson.M{
		"_id":    objID,
		"userid": idUsuario,
	}

	resultadoBorrado, err := col.DeleteOne(ctx, condicion)

	if err != nil {
		log.Fatal(err)
		borrado = false
	}
	//fmt.Printf("Deleted %v documents in the trainers collection\n", resultadoBorrado.DeletedCount)

	borrado = true
	return resultadoBorrado, borrado
}
