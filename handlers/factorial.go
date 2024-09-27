package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Factorial calcula el factorial de un número dado.
func Factorial(c *gin.Context) {
	numberStr := c.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil || number < 0 {
		c.String(http.StatusBadRequest, "Invalid number")
		return
	}

	result := factorial(number)

	c.JSON(http.StatusOK, gin.H{
		"number":    number,
		"factorial": result,
	})
}

// factorial calcula el factorial de un número de forma recursiva.
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
