// Lee el JSON que manda el usuario revisa los datos y llama al repo
// para guardar las cosas y decide que codigo http devolver
package handlers

import (
	"DroneManager/internal/models"
	"DroneManager/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Crear un drone nuevo
func CreateDrone(c *gin.Context) {
	var drone models.Drone

	if err := c.ShouldBindJSON(&drone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := repository.CreateDrone(drone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error guardando en BD"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Drone registrado exitosamente",
		"id":      result,
	})
}

func GetDrones(c *gin.Context) {
	drones, err := repository.GetAllDrones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error leyendo datos"})
		return
	}

	c.JSON(http.StatusOK, drones)
}

func UpdateDrone(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Status string `json:"status"`
		Charge int    `json:"charge"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	droneData := models.Drone{
		Status: input.Status,
		Charge: input.Charge,
	}

	updatedCount, err := repository.UpdateDrone(id, droneData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error actualizando drone"})
		return
	}

	if updatedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No se encontró el drone"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Drone actualizado correctamente"})
}
