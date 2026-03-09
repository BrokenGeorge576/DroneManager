#!/bin/bash

SERVER="localhost:50051"

echo "=========================================="
echo "INICIANDO PRUEBAS E2E DE LA FLOTA (gRPC)"

echo -e "\nConsultando drones en la base de datos:"
grpcurl -plaintext $SERVER drone.DroneService/GetDrones

echo -e "\nCreando Dron Alfa (Batería al 100%)..."
grpcurl -plaintext -d '{"model": "DJI Mavic 3", "serial": "ALFA-001", "charge": 100}' $SERVER drone.DroneService/CreateDrone

echo -e "\nCreando Dron Beta (Batería al 15%)..."
grpcurl -plaintext -d '{"model": "DJI Mini Pro", "serial": "BETA-002", "charge": 15}' $SERVER drone.DroneService/CreateDrone

echo -e "\nVerificando que la flota se actualizó:"
grpcurl -plaintext $SERVER drone.DroneService/GetDrones

echo -e "\nPruebas E2E finalizadas con éxito."
echo "=========================================="
