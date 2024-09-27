package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Calculate realiza la operación aritmética básica entre dos números.
func Calculate(c *gin.Context) {
	num1Str := c.Query("num1")
	num2Str := c.Query("num2")
	operator := c.Query("operator")

	num1, err1 := strconv.ParseFloat(num1Str, 64)
	num2, err2 := strconv.ParseFloat(num2Str, 64)

	if err1 != nil || err2 != nil {
		c.String(http.StatusBadRequest, "Invalid numbers")
		return
	}

	var result float64
	var err error

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			err = fmt.Errorf("division by zero")
		} else {
			result = num1 / num2
		}
	default:
		err = fmt.Errorf("invalid operator")
	}

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
