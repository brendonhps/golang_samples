package main

import (
	"fmt"
	"time"
)

func showMessage(c chan string ) {

	for range time.Tick(time.Second * 3) {
		for el := range c {
			fmt.Println(el)
		}
		fmt.Println(<-c)
	}
}

func main() {
	c := make(chan string,30)

	c <- "Hello"
	go showMessage(c)

	c <- "World"
	go showMessage(c)
	time.Sleep(time.Second * 20)
}
