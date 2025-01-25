package service

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomID() int {
	return rand.Intn(200)
}
func GenerateRandomName(length int) string {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var sb strings.Builder

	for i := 0; i < length; i++ {
		sb.WriteRune(letters[randGen.Intn(len(letters))]) // استفاده از randGen به جای rand
	}

	return sb.String()
}
