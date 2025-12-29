// Definiendo el contrato de datos
// Este codigo define lo que es un drone, no tiene ninguna logica, solo definiciones, es el MODELO
package models // Este archivo pertenece a la carpeta models

import (
	"time" //Libreria de Go para fechas y horas

	"go.mongodb.org/mongo-driver/bson/primitive" // Reconocer codigos de los objetos en Mongo
)

type Drone struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`                 // ID unico para el dron
	Model     string             `bson:"model" json:"model" binding:"required"`   // Modelo
	Serial    string             `bson:"serial" json:"serial" binding:"required"` // Numero de serie
	Charge    int                `bson:"charge" json:"charge"`                    // Estado de la bateria 0-100
	Status    string             `bson:"status" json:"status"`                    // Estatus del dron
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`            // Fecha de creacion
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`            // Fecha de actualizacion
}
