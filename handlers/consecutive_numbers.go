package handlers

import (
	"strconv"
)

func GenerateConsecutiveNumbers(lastNumberStr string) ([]int, error) {
	lastNumber, err := strconv.Atoi(lastNumberStr)
	if err != nil {
		return nil, err
	}

	numbers := make([]int, lastNumber)
	for i := 0; i < lastNumber; i++ {
		numbers[i] = i + 1
	}

	return numbers, nil
}
