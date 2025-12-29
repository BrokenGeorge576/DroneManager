// Este codigo crea la conexion con Docker, mantiene una variable global
// para no tener que reconectar cada vez que haya un query y con la funcion GetCollection
// dirige donde se guarda la info
package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client es la variable global que guarda la conexión activa
var Client *mongo.Client

// ConnectDB inicia la conexión con Docker
func ConnectDB() {
	// Definimos un tiempo límite de 10 segundos para intentar conectarse
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// La URL de conexión (localhost:27017 es tu contenedor Docker)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error creando cliente Mongo:", err)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("No se pudo conectar a MongoDB:", err)
	}

	fmt.Println("Conectado exitosamente a MongoDB (Docker)")
}

// Ayudante para obtener rápido una colección (tabla)
func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("drone_fleet_db").Collection(collectionName)
}
