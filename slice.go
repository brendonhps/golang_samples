package main

import (
	"fmt"
	"math"
)

func divideSlice(str string) ([]string, []string) {
	var first, second []string
	x := len(str) * 1.0
	for i, v := range str {
		if float64(i) < math.Ceil(float64(x/2.0)) {
			first = append(first, string(v))
		} else {
			second = append(second, string(v))
		}
	}
	return first, second
}

func main() {
	texto := "Hello World"
	substr1, substr2 := divideSlice(texto)
	fmt.Printf("%s , %s\n", substr1, substr2)
}
