package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Countdown imprime una cuenta regresiva desde 10 hasta 1.
func Countdown(c *gin.Context) {
	countdown := []int{}
	for i := 10; i >= 1; i-- {
		countdown = append(countdown, i)
	}

	c.JSON(http.StatusOK, gin.H{
		"countdown": countdown,
	})
}
