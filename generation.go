package main

import (
	"crypto/rand"
	"math/big"
)

func generateRandomString(length int) string {
	big_length := big.NewInt(int64(length))
	result := ""

	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big_length)
		result += string(StandardCharset[index.Int64()])
	}

	return result
}
