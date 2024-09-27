package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateID genera un ID aleatorio.
func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Intn(1000000))
}
