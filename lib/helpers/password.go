package helpers

import (
	"math/rand"
)

const (
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes    = "0123456789"
	specialBytes   = "!@#$%^&*()-_=+[]{}|;:,.<>?/~`"
	allBytes       = letterBytes + numberBytes + specialBytes
	passwordLength = 12 // Change this to your desired password length
)

func GetRandomPassword() string {
	b := make([]byte, passwordLength)
	for i := range b {
		b[i] = allBytes[rand.Intn(len(allBytes))]
	}
	return string(b)
}
