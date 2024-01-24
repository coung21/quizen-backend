package common

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[RandomInt(0, len(letter)-1)]
	}
	return string(b)
}

func RandomInt(min, max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + rand.Intn(max-min+1)
}
