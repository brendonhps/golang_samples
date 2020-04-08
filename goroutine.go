package main

import "fmt"

func f(i int) {
	fmt.Println(i)
}

func main() {

	for j := 0; j < 100; j++ {
		go f(j)
	}

	fmt.Scanln()
}
