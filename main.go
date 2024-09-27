package main

import (
	"ejemplo1/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Rutas para la entidad Avión
	r.POST("/airplanes", handlers.CreateAirplane)
	r.GET("/airplanes", handlers.GetAirplanes)
	r.GET("/airplanes/:id", handlers.GetAirplaneByID)
	r.PUT("/airplanes/:id", handlers.UpdateAirplane)
	r.DELETE("/airplanes/:id", handlers.DeleteAirplane)

	//Ejercicios cortos del Slide

	// r.GET("/consecutive_numbers/:lastNumber", func(c *gin.Context) {
	// 	lastNumberStr := c.Param("lastNumber")
	// 	numbers, err := handlers.GenerateConsecutiveNumbers(lastNumberStr)
	// 	if err != nil {
	// 		c.String(400, "Invalid number")
	// 		return
	// 	}
	// 	c.JSON(200, numbers)
	// })

	// r.GET("/calculate", handlers.Calculate)

	// r.GET("/multiplication_table/:number", handlers.MultiplicationTable)

	// r.GET("/convert_temperature/:celsius", handlers.ConvertTemperature)

	// r.GET("/is_prime/:number", handlers.IsPrime)

	// r.GET("/fibonacci/:n", handlers.Fibonacci)

	// r.GET("/countdown", handlers.Countdown)

	// r.GET("/palindrome/:word", handlers.IsPalindrome)

	// r.GET("/even_odd/:n", handlers.EvenOdd)

	// r.GET("/factorial/:number", handlers.Factorial)

	// r.GET("/max_min", handlers.MaxMin)

	// r.POST("/vowel_counter", handlers.VowelCounter)

	// r.GET("/reverse_string", handlers.ReverseString)

	r.Run() // Por defecto, corre en localhost:8080
}
