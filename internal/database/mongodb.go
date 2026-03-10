// Este codigo crea la conexion con Docker o local
package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client // Cliente para hacer consultas

func ConnectDB() { // Funcion para conectar a la base de datos se llama desde main.go
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Timeout de 10 segundos
	defer cancel()                                                           // Cancela el contexto al finalizar la funcion
	mongoURI := os.Getenv("MONGO_URI")                                       // Obtiene la URI de la base de datos desde las variables de entorno
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" // URI por defecto si no se especifica
	}

	clientOptions := options.Client().ApplyURI(mongoURI) // Paquete de configuracion de la base de datos

	var err error
	Client, err = mongo.Connect(ctx, clientOptions) // Conecta al cliente con la configuracion especificada. Se intenta abrir la conexion
	if err != nil {
		log.Fatal("Error creando cliente Mongo:", err)
	}

	err = Client.Ping(ctx, nil) // Verifica que la conexion este activa
	if err != nil {
		log.Fatal("No se pudo conectar a MongoDB:", err)
	}

	fmt.Printf("Conectado exitosamente a MongoDB en: %s\n", mongoURI)
}

func GetCollection(collectionName string) *mongo.Collection { // Obtiene una coleccion de la base de datos
	return Client.Database("drone_fleet_db").Collection(collectionName)
}
