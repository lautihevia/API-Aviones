package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// MaxMin toma una serie de números como entrada y muestra el número mayor y el número menor.
func MaxMin(c *gin.Context) {
	numbersStr := c.Query("numbers")
	if numbersStr == "" {
		c.String(http.StatusBadRequest, "No numbers provided")
		return
	}

	numberStrings := strings.Split(numbersStr, ",")
	numbers := []int{}
	for _, numStr := range numberStrings {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid number: %s", numStr)
			return
		}
		numbers = append(numbers, num)
	}

	if len(numbers) == 0 {
		c.String(http.StatusBadRequest, "No valid numbers provided")
		return
	}

	max, min := numbers[0], numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"max": max,
		"min": min,
	})
}
