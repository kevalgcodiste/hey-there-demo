package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumber(min, max int) int {
	if max <= min {
		return min
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func main() {
	fmt.Println(RandomNumber(1, 100))
}
