package main

import (
	"DroneManager/internal/database"
	"DroneManager/internal/handlers"
	"DroneManager/internal/repository"
	"DroneManager/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	database.ConnectDB()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Fallo al escuchar en puerto: %v", err)
	}

	grpcServer := grpc.NewServer()
	repoReal := repository.NewMongoDroneRepo()
	droneServer := &handlers.DroneGRPCServer{
		Repo: repoReal,
	}

	pb.RegisterDroneServiceServer(grpcServer, droneServer)
	reflection.Register(grpcServer)

	fmt.Println("Servidor gRPC corriendo en puerto 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo al levantar gRPC: %v", err)
	}
}
