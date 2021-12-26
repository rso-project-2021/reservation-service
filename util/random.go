package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	var sb strings.Builder
	length := len(letters)

	for i := 0; i < n; i++ {
		c := letters[rand.Intn(length)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomTimestamp() time.Time {
	return time.Now()
}
