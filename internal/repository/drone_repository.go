// El unico que tiene permiso de entrar a MondoDB
// Aqui se realiza el CRUD tecnico, recibe structs que convierte a BSON y viceversa

package repository

import (
	"DroneManager/internal/database"
	"DroneManager/internal/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDrone(drone models.Drone) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	drone.ID = primitive.NewObjectID()
	drone.CreatedAt = time.Now()
	drone.UpdatedAt = time.Now()

	if drone.Status == "" {
		drone.Status = "Disponible"
	}

	collection := database.GetCollection("drones")

	result, err := collection.InsertOne(ctx, drone)
	return result.InsertedID, err
}

func GetAllDrones() ([]models.Drone, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := database.GetCollection("drones")

	var drones []models.Drone

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var drone models.Drone
		if err := cursor.Decode(&drone); err != nil {
			return nil, err
		}
		drones = append(drones, drone)
	}

	return drones, nil
}

func UpdateDrone(idString string, updateData models.Drone) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := database.GetCollection("drones")

	objectId, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return 0, err // Si el ID tiene formato inválido, regresamos error
	}

	// "$set" comando de Mongo para decir "Solo toca estos campos, no borres lo demás"
	update := bson.M{
		"$set": bson.M{
			"status":     updateData.Status,
			"charge":     updateData.Charge,
			"updated_at": time.Now(), // actualizar la fecha de modificación
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objectId}, update)

	return result.ModifiedCount, err
}
