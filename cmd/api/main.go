// Enciende todo y asigna puestos
// Funcionamiento general:
// 1.- Internet golpea el puerto 8080.
// 2.- Main ve la ruta y desvía el tráfico al Handler.
// 3.- Handler revisa que el JSON esté bien y se lo pasa al Repository.
// 4.- Repository convierte los datos y los empuja por la tubería de Database.
// 5.- MongoDB guarda el dato.

package main

import (
	"DroneManager/internal/database" // Importamos la conexión
	"DroneManager/internal/handlers" // Importamos los controladores

	"github.com/gin-gonic/gin"
)

func main() {
	// nos aseguramos que Mongo esté listo.
	database.ConnectDB()

	// Creamos una instancia de Gin (el framework web) con configuración por defecto
	r := gin.Default()

	// Agrupamos las rutas bajo "/api/v1"
	v1 := r.Group("/api/v1")
	{
		// Si mandan POST a /api/v1/drones -> Ejecuta CreateDrone
		v1.POST("/drones", handlers.CreateDrone)

		// Si mandan GET a /api/v1/drones -> Ejecuta GetDrones
		v1.GET("/drones", handlers.GetDrones)

		v1.PUT("/drones/:id", handlers.UpdateDrone)
	}

	// Escuchamos en el puerto 8080.
	// El programa se queda "bloqueado" aquí esperando peticiones.
	r.Run(":8080")
}
