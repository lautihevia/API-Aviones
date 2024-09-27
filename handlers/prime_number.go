package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IsPrime determina si un número dado es primo.
func IsPrime(c *gin.Context) {
	numberStr := c.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid number")
		return
	}

	if number <= 1 {
		c.JSON(http.StatusOK, gin.H{
			"number":  number,
			"isPrime": false,
		})
		return
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			c.JSON(http.StatusOK, gin.H{
				"number":  number,
				"isPrime": false,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"number":  number,
		"isPrime": true,
	})
}
