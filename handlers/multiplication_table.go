package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MultiplicationTable genera la tabla de multiplicar para un número dado.
func MultiplicationTable(c *gin.Context) {
	numberStr := c.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid number")
		return
	}

	table := make([]string, 10)
	for i := 1; i <= 10; i++ {
		table[i-1] = strconv.Itoa(number) + " x " + strconv.Itoa(i) + " = " + strconv.Itoa(number*i)
	}

	c.JSON(http.StatusOK, gin.H{
		"table": table,
	})
}
