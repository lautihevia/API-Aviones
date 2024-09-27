package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// VowelCounterRequest representa la estructura del cuerpo de la solicitud.
type VowelCounterRequest struct {
	Text string `json:"text"`
}

// VowelCounterResponse representa la estructura de la respuesta.
type VowelCounterResponse struct {
	VowelCount int `json:"vowel_count"`
}

// VowelCounter cuenta el número de vocales en una cadena dada.
func VowelCounter(c *gin.Context) {
	var req VowelCounterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}

	vowelCount := countVowels(req.Text)

	c.JSON(http.StatusOK, VowelCounterResponse{
		VowelCount: vowelCount,
	})
}

// countVowels cuenta el número de vocales en una cadena.
func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}
