package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomTip returns a small random tip string.
func RandomTip() string {
	tips := []string{
		"Keep functions small and focused.",
		"Prefer errors over panics for expected failures.",
		"Use gofmt and keep imports tidy.",
		"Write table-driven tests.",
	}
	rand.Seed(time.Now().UnixNano())
	return tips[rand.Intn(len(tips))]
}

func main() {
	fmt.Println(RandomTip())
}
