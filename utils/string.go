package utils

import (
	"math/rand"
	"strconv"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 ")

// GenerateString generates a random string with len n
func GenerateString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//Atoi converts string to int
func Atoi(data string) int {
	val, err := strconv.Atoi(data)
	if err != nil {
		return 20
	}
	return val
}
