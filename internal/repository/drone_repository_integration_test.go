//go:build integration

package repository

import (
	"context"
	"testing"

	"DroneManager/internal/database"
	"DroneManager/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestIntegration_FullCRUD_RealDB(t *testing.T) {
	database.ConnectDB()
	repo := NewMongoDroneRepo()

	testDrone := models.Drone{
		Model:  "Prueba Integracion CRUD",
		Serial: "CRUD-999",
		Charge: 100,
	}

	insertedID, err := repo.CreateDrone(testDrone)
	if err != nil {
		t.Fatalf("Falló CreateDrone: %v", err)
	}
	oid, ok := insertedID.(primitive.ObjectID)
	if !ok {
		t.Fatalf("El ID devuelto no es válido")
	}
	idString := oid.Hex()
	t.Logf("CREATE: Dron insertado con ID: %s", idString)

	updateData := models.Drone{
		Status: "En Mantenimiento",
		Charge: 80,
	}

	modifiedCount, err := repo.UpdateDrone(idString, updateData)
	if err != nil || modifiedCount == 0 {
		t.Fatalf("Falló UpdateDrone. Modificados: %d, Error: %v", modifiedCount, err)
	}
	t.Log("UPDATE: Dron actualizado correctamente.")

	drones, err := repo.GetAllDrones()
	if err != nil || len(drones) == 0 {
		t.Fatalf("Falló GetAllDrones o vino vacío: %v", err)
	}

	encontrado := false
	for _, d := range drones {
		if d.ID == oid {
			encontrado = true
			if d.Status != "En Mantenimiento" || d.Charge != 80 {
				t.Errorf("Los datos no se actualizaron bien en la BD. Status: %s", d.Status)
			}
			break
		}
	}

	if !encontrado {
		t.Errorf("El dron insertado no apareció en GetAllDrones")
	} else {
		t.Log("GET: Dron encontrado con los datos actualizados.")
	}

	t.Cleanup(func() {
		collection := database.GetCollection("drones")
		_, err := collection.DeleteOne(context.Background(), bson.M{"_id": oid})

		if err != nil {
			t.Fatalf("Advertencia: No se pudo limpiar el dron de prueba: %v", err)
		} else {
			t.Log("CLEANUP: Dron de prueba eliminado. La base de datos quedó vacía e impecable.")
		}
	})
}
