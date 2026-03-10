package handlers

import (
	"context" // es el motor que ejecuta las pruebas.
	"errors"  // es el motor que ejecuta las pruebas.
	"testing" // es el motor que ejecuta las pruebas.

	"DroneManager/internal/mocks" //Importamos la carpeta donde guardaste tu clon falso (el código generado por MockGen).
	"DroneManager/pb"

	"go.uber.org/mock/gomock" // La librería de Uber que controla a nuestro clon falso.
)

func TestCreateDrone_Success(t *testing.T) { // Funcion de testing que recibe una solicitud de creacion de drone y verifica que la respuesta sea correcta.
	ctrl := gomock.NewController(t)                // Crea un nuevo controlador de Mocks.
	defer ctrl.Finish()                            // Cierra el controlador al finalizar la prueba y avisa si responde como deberia.
	mockRepo := mocks.NewMockDroneRepository(ctrl) // En lugar de conectarnos a Mongo, instanciamos un repositorio falso (mockRepo).
	mockRepo.EXPECT().
		CreateDrone(gomock.Any()).
		Return("mock-id-123", nil)

	server := &DroneGRPCServer{ // Se crea un servidor con el repositorio falso.
		Repo: mockRepo,
	}

	req := &pb.CreateDroneRequest{ // Inyeccion de dependencias
		Model:  "DJI Mavic Pro",
		Serial: "XYZ-999",
		Charge: 100,
	}

	res, err := server.CreateDrone(context.Background(), req) // Se llama al metodo CreateDrone del servidor con la solicitud de prueba.

	if err != nil { // Si hubo algún error, abortamos la prueba (t.Fatalf).
		t.Fatalf("Se esperaba éxito, pero dio error: %v", err)
	}

	if res.GetId() != "mock-id-123" { //Verificamos que el código empaquetó correctamente en el response el ID que el clon le inventó.
		t.Errorf("Se esperaba el ID 'mock-id-123', pero llegó: %s", res.GetId())
	}

	expectedMessage := "Drone registrado exitosamente" //Nos aseguramos de que el mensaje de éxito que manda tu servidor gRPC no haya sido alterado por accidente.
	if res.GetMessage() != expectedMessage {
		t.Errorf("Se esperaba '%s', pero llegó: '%s'", expectedMessage, res.GetMessage())
	}
}

func TestCreateDrone_Failure_DBError(t *testing.T) { // Prueba de fallo por DB
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockDroneRepository(ctrl)
	dbError := errors.New("timeout conectando a MongoDB")

	mockRepo.EXPECT().
		CreateDrone(gomock.Any()).
		Return(nil, dbError) // Regresara error en DB

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
