package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReverseString invierte una cadena dada.
func ReverseString(c *gin.Context) {
	text := c.Query("text")
	if text == "" {
		c.String(http.StatusBadRequest, "Text query parameter is required")
		return
	}

	reversedText := reverse(text)

	c.JSON(http.StatusOK, gin.H{
		"reversed_text": reversedText,
	})
}

// reverse invierte una cadena.
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
