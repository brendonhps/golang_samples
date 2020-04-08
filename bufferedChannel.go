package main

import "fmt"

func main() {

	var c chan int = make(chan int, 5)

	for i := 100; i >= 96; i-- {
		c <- i
	}

	j := 0
	for j < 5 {
		fmt.Println(<-c)
		j++
	}

}
