package handlers

import (
	"context"
	"fmt"

	"DroneManager/internal/models"
	"DroneManager/internal/repository"
	"DroneManager/pb" // Importamos los archivos generados por protoc
)

// DroneGRPCServer implementa la interfaz generada por gRPC para DroneService
type DroneGRPCServer struct {
	pb.UnimplementedDroneServiceServer
}

// ------------------------------------------------------------------------
// 1. Crear Drone (Equivalente al POST /api/v1/drones)
// ------------------------------------------------------------------------
func (s *DroneGRPCServer) CreateDrone(ctx context.Context, req *pb.CreateDroneRequest) (*pb.DroneResponse, error) {
	// Mapeamos los datos de gRPC a nuestro modelo interno de Go
	drone := models.Drone{
		Model:  req.GetModel(),
		Serial: req.GetSerial(),
		Charge: int(req.GetCharge()),
	}

	// Usamos tu repositorio existente de MongoDB sin cambiarle nada
	result, err := repository.CreateDrone(drone)
	if err != nil {
		// En gRPC devolvemos el error directo, no un c.JSON(http...)
		return nil, fmt.Errorf("error guardando en BD: %v", err)
	}

	// Convertimos el ObjectID de Mongo (interface{}) a string
	idStr := fmt.Sprintf("%v", result)

	// Retornamos el mensaje de respuesta de gRPC
	return &pb.DroneResponse{
		Id:      idStr,
		Message: "Drone registrado exitosamente",
	}, nil
}

// ------------------------------------------------------------------------
// 2. Obtener todos los Drones (Equivalente al GET /api/v1/drones)
// ------------------------------------------------------------------------
func (s *DroneGRPCServer) GetDrones(ctx context.Context, req *pb.Empty) (*pb.DroneListResponse, error) {
	// Llamamos a tu repositorio actual
	drones, err := repository.GetAllDrones()
	if err != nil {
		return nil, fmt.Errorf("error leyendo datos de la BD: %v", err)
	}

	// gRPC espera un slice de *pb.DroneItem, así que convertimos los resultados
	var pbDrones []*pb.DroneItem
	for _, d := range drones {
		pbDrones = append(pbDrones, &pb.DroneItem{
			Id:     d.ID.Hex(), // Usamos .Hex() para sacar el string limpio de Mongo
			Model:  d.Model,
			Status: d.Status,
		})
	}

	// Retornamos la lista empaquetada en el formato gRPC
	return &pb.DroneListResponse{
		Drones: pbDrones,
	}, nil
}

// ------------------------------------------------------------------------
// 3. Actualizar Drone (Equivalente al PUT /api/v1/drones/:id)
// ------------------------------------------------------------------------
func (s *DroneGRPCServer) UpdateDrone(ctx context.Context, req *pb.UpdateDroneRequest) (*pb.DroneResponse, error) {
	// Mapeamos los datos que nos envía el cliente gRPC
	droneData := models.Drone{
		Status: req.GetStatus(),
		Charge: int(req.GetCharge()),
	}

	// Llamamos al repositorio usando el ID que viene en el request
	updatedCount, err := repository.UpdateDrone(req.GetId(), droneData)
	if err != nil {
		return nil, fmt.Errorf("error actualizando drone en BD: %v", err)
	}

	// Si Mongo no modificó nada, significa que no lo encontró
	if updatedCount == 0 {
		return nil, fmt.Errorf("no se encontró el drone con ID: %s", req.GetId())
	}

	// Respuesta exitosa
	return &pb.DroneResponse{
		Id:      req.GetId(),
		Message: "Drone actualizado correctamente",
	}, nil
}
