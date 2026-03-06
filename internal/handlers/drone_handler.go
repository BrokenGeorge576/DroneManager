package handlers

import (
	"context"
	"fmt"

	"DroneManager/internal/models"
	"DroneManager/pb"
)

type DroneRepository interface {
	CreateDrone(drone models.Drone) (interface{}, error)
	GetAllDrones() ([]models.Drone, error)
	UpdateDrone(idString string, updateData models.Drone) (int64, error)
}

type DroneGRPCServer struct {
	pb.UnimplementedDroneServiceServer
	Repo DroneRepository
}

// 1. Crear Drone
func (s *DroneGRPCServer) CreateDrone(ctx context.Context, req *pb.CreateDroneRequest) (*pb.DroneResponse, error) {
	drone := models.Drone{
		Model:  req.GetModel(),
		Serial: req.GetSerial(),
		Charge: int(req.GetCharge()),
	}

	result, err := s.Repo.CreateDrone(drone)
	if err != nil {
		return nil, fmt.Errorf("error guardando en BD: %v", err)
	}

	idStr := fmt.Sprintf("%v", result)

	return &pb.DroneResponse{
		Id:      idStr,
		Message: "Drone registrado exitosamente",
	}, nil
}

// 2. Obtener todos los Drones
func (s *DroneGRPCServer) GetDrones(ctx context.Context, req *pb.Empty) (*pb.DroneListResponse, error) {
	drones, err := s.Repo.GetAllDrones()
	if err != nil {
		return nil, fmt.Errorf("error leyendo datos de la BD: %v", err)
	}

	var pbDrones []*pb.DroneItem
	for _, d := range drones {
		pbDrones = append(pbDrones, &pb.DroneItem{
			Id:     d.ID.Hex(),
			Model:  d.Model,
			Status: d.Status,
		})
	}

	return &pb.DroneListResponse{
		Drones: pbDrones,
	}, nil
}

// 3. Actualizar Drone
func (s *DroneGRPCServer) UpdateDrone(ctx context.Context, req *pb.UpdateDroneRequest) (*pb.DroneResponse, error) {
	droneData := models.Drone{
		Status: req.GetStatus(),
		Charge: int(req.GetCharge()),
	}

	updatedCount, err := s.Repo.UpdateDrone(req.GetId(), droneData)
	if err != nil {
		return nil, fmt.Errorf("error actualizando drone en BD: %v", err)
	}

	if updatedCount == 0 {
		return nil, fmt.Errorf("no se encontró el drone con ID: %s", req.GetId())
	}

	return &pb.DroneResponse{
		Id:      req.GetId(),
		Message: "Drone actualizado correctamente",
	}, nil
}
