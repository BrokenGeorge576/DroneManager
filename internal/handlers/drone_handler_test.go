package handlers

import (
	"context"
	"errors"
	"testing"

	"DroneManager/internal/mocks"
	"DroneManager/pb"

	"go.uber.org/mock/gomock"
)

func TestCreateDrone_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockDroneRepository(ctrl)
	mockRepo.EXPECT().
		CreateDrone(gomock.Any()).
		Return("mock-id-123", nil)

	server := &DroneGRPCServer{
		Repo: mockRepo,
	}

	req := &pb.CreateDroneRequest{
		Model:  "DJI Mavic Pro",
		Serial: "XYZ-999",
		Charge: 100,
	}

	res, err := server.CreateDrone(context.Background(), req)

	if err != nil {
		t.Fatalf("Se esperaba éxito, pero dio error: %v", err)
	}

	if res.GetId() != "mock-id-123" {
		t.Errorf("Se esperaba el ID 'mock-id-123', pero llegó: %s", res.GetId())
	}

	expectedMessage := "Drone registrado exitosamente"
	if res.GetMessage() != expectedMessage {
		t.Errorf("Se esperaba '%s', pero llegó: '%s'", expectedMessage, res.GetMessage())
	}
}

func TestCreateDrone_Failure_DBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockDroneRepository(ctrl)
	dbError := errors.New("timeout conectando a MongoDB")

	mockRepo.EXPECT().
		CreateDrone(gomock.Any()).
		Return(nil, dbError)

	server := &DroneGRPCServer{
		Repo: mockRepo,
	}

	req := &pb.CreateDroneRequest{
		Model:  "DJI Mavic Pro",
		Serial: "XYZ-999",
		Charge: 100,
	}

	res, err := server.CreateDrone(context.Background(), req)

	if err == nil {
		t.Fatalf("Se esperaba que fallara por error de BD, pero la operación fue exitosa")
	}

	if res != nil {
		t.Errorf("Se esperaba un response nulo, pero llegó: %v", res)
	}
}
