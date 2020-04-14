package main

import (
	"fmt"
	"time"
)

func f(c chan string) {
	for {
		time.Sleep(2 * time.Second)
		c <- "1 seconds passed"
	}
}

func main() {
	c1:= make(chan string, 1)
	go f(c1)

	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout!")

		}
	}
}