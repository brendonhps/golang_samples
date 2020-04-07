package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomNum := random(1, 6)
		fmt.Printf("Play %d - Random number : %d\n", i, randomNum)
	}
}


