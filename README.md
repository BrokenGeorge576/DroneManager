# 🚁 DroneManager API

API RESTful desarrollada en **Go** para la gestión de flotas de drones. Implementa una **Clean Architecture** y utiliza **MongoDB** como base de datos NoSQL.

## 🛠 Tech Stack
* **Lenguaje:** Go (Golang)
* **Framework Web:** Gin Gonic
* **Base de Datos:** MongoDB (Driver oficial)
* **Infraestructura:** Docker (para la instancia de Mongo)

## 🚀 Funcionalidades (CRUD)
* **POST** `/api/v1/drones` - Registrar un nuevo drone.
* **GET** `/api/v1/drones` - Listar flota (soporta filtros, ej: `?status=Disponible`).
* **PUT** `/api/v1/drones/:id` - Actualizar estado y batería.
* **DELETE** `/api/v1/drones/:id` - Eliminar registros.

## 📦 Instalación y Uso

1. Clonar el repositorio:
   ```bash
   git clone [https://github.com/TU_USUARIO/drone-manager-go.git](https://github.com/TU_USUARIO/drone-manager-go.git)
