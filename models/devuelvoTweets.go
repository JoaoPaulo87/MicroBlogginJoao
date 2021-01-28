package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Vamos a ir a la carpeta models a crear una estructura para ver como voy a devolver los tweets,
como voy a devolver esa información al http para que sea procesada por el frontend*/

/*DevuelvoTweets es la estructura con la que devolveremos los Tweets*/
type DevuelvoTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userid" json:"userId,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
