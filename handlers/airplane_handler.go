package handlers

import (
	"ejemplo1/dto"
	"ejemplo1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAirplane maneja la creación de un nuevo avión.
func CreateAirplane(c *gin.Context) {
	//aca intentamos hacer un bind del request que llega al endpoint con el dto que creamos,
	//y si hay un error, devolvemos un error 400
	var req dto.CreateAirplaneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//llamamos al servicio que creamos, y si hay un error, devolvemos un error 500
	airplane, err := services.CreateAirplane(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, airplane)
}

// GetAirplanes maneja la obtención de aviones, con filtrado opcional por aerolínea.
func GetAirplanes(c *gin.Context) {
	airline := c.Query("airline")
	airplanes, err := services.GetAirplanes(airline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, airplanes)
}

// GetAirplaneByID maneja la obtención de un avión por ID.
func GetAirplaneByID(c *gin.Context) {
	id := c.Param("id")
	airplane, err := services.GetAirplaneByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, airplane)
}

// UpdateAirplane maneja la actualización de un avión por ID.
func UpdateAirplane(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateAirplaneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	airplane, err := services.UpdateAirplane(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, airplane)
}

// DeleteAirplane maneja la eliminación de un avión por ID.
func DeleteAirplane(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteAirplane(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
