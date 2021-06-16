package utils

import "math/rand"

// GenerateNumbers generates a random int32 list with amount len
func GenerateNumbers(amount int) []int32 {
	b := make([]int32, amount)
	for i := range b {
		b[i] = int32(rand.Intn(2147483647))
	}
	return b
}

// Average calculates the average for a given int32 array
func Average(numbers []int32) int32 {
	var total int32 = 0
	for _, number := range numbers {
		total += number
	}
	return total / int32(len(numbers))
}
