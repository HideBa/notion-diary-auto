package rand

import (
	"math/rand"
	"time"
)

func GenerateRandomStr(letters []rune, n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
