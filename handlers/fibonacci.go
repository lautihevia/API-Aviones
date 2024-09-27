package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Fibonacci genera los primeros N términos de la secuencia de Fibonacci.
func Fibonacci(c *gin.Context) {
	nStr := c.Param("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		c.String(http.StatusBadRequest, "Invalid number")
		return
	}

	fibSequence := make([]int, n)
	if n > 0 {
		fibSequence[0] = 0
	}
	if n > 1 {
		fibSequence[1] = 1
	}
	for i := 2; i < n; i++ {
		fibSequence[i] = fibSequence[i-1] + fibSequence[i-2]
	}

	c.JSON(http.StatusOK, gin.H{
		"fibonacci": fibSequence,
	})
}
