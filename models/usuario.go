package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Usuario es el modelo de usuario de la base de mongoDB */
type Usuario struct {
	/* omitempty : si el campo está vacío lo omite*/
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre, omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos, omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento, omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password, omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar, omitempty"`
	Banner          string             `bson:"banner" json:"banner, omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia, omitempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb, omitempty"`
}
