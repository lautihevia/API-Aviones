package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// IsPalindrome determina si una palabra dada es un palíndromo.
func IsPalindrome(c *gin.Context) {
	word := c.Param("word")
	normalizedWord := strings.ToLower(word)
	reversedWord := reverseString(normalizedWord)

	isPalindrome := normalizedWord == reversedWord

	c.JSON(http.StatusOK, gin.H{
		"word":         word,
		"isPalindrome": isPalindrome,
	})
}

// reverseString invierte una cadena de texto.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
