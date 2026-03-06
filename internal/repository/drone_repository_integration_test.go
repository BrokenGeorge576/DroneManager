//go:build integration

package repository

import (
	"testing"

	"DroneManager/internal/database"
	"DroneManager/internal/models"
)

func TestIntegration_CreateDrone_RealDB(t *testing.T) {
	database.ConnectDB()
	repo := NewMongoDroneRepo()

	testDrone := models.Drone{
		Model:  "Prueba Integracion DB",
		Serial: "INT-001",
		Charge: 50,
	}

	insertedID, err := repo.CreateDrone(testDrone)

	if err != nil {
		t.Fatalf("Falló la conexión o inserción en la BD real: %v", err)
	}

	if insertedID == nil {
		t.Errorf("Se esperaba un ID de Mongo, pero se recibió nil")
	}

	t.Logf("¡Éxito! Dron insertado en BD real con ID: %v", insertedID)
}
