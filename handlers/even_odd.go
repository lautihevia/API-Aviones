package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EvenOdd imprime los primeros N números pares e impares.
func EvenOdd(c *gin.Context) {
	nStr := c.Param("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		c.String(http.StatusBadRequest, "Invalid number")
		return
	}

	evenNumbers := []int{}
	oddNumbers := []int{}

	for i := 0; len(evenNumbers) < n || len(oddNumbers) < n; i++ {
		if i%2 == 0 {
			if len(evenNumbers) < n {
				evenNumbers = append(evenNumbers, i)
			}
		} else {
			if len(oddNumbers) < n {
				oddNumbers = append(oddNumbers, i)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"even_numbers": evenNumbers,
		"odd_numbers":  oddNumbers,
	})
}
