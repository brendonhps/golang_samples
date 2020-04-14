package main

import (
	"fmt"
	"time"
)

func f1(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		c <- "1"
	}
}


func f2(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		c <- "2"
	}
}

func main() {

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go f1(c1)
	go f2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

	}
	fmt.Println("finish GoRoutines")

}