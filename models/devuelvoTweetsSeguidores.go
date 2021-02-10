/*Vamos a realizar ahora un endpoint para poder listar todos los tweets de los usuarios que sigo,
 ordenados por fecha (del último hacia atrás).

Vamos a tener que trabajar con las tablas unidas, es decir los usuarios, los tweets y las relaciones,
 teniendo en cuenta todo. Para eso utilizaremos un framework nuevo que se llama Agregate que se utiliza con Mongo DB.

Vamos a tener que crear un modelo para contener lo que devuelve esta función
que lee los tweets: struct DevuelvoTweetsSeguidores.
*/
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoTweetsSeguidores es la estructura con la que devolveremos los tweets*/
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID            string             `bson:"usuarioid" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"userRelationId"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
