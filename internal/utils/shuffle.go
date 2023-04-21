package utils

import (
	"math/rand"
	"time"
)

func Shuffle(strings []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := len(strings) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		strings[i], strings[j] = strings[j], strings[i]
	}
	return strings
}
