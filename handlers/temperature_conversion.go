package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ConvertTemperature convierte una temperatura de Celsius a Fahrenheit.
func ConvertTemperature(c *gin.Context) {
	celsiusStr := c.Param("celsius")
	celsius, err := strconv.ParseFloat(celsiusStr, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid temperature")
		return
	}

	fahrenheit := celsius*9/5 + 32

	c.JSON(http.StatusOK, gin.H{
		"celsius":    celsius,
		"fahrenheit": fahrenheit,
	})
}
