package main

import (
	"fmt"
	"time"
)

// Example 1
//func main() {
//	c := make(chan int, 10)
//
//	for i:= 0; i < 10; i++ {
//		c <- i
//	}
//	close(c)
//
//	for element := range c {
//		fmt.Println(element)
//	}
//}

func f( c chan int) {
	for {
		c <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

func main() {
	channel := make(chan int)

	go f(channel)
	for c := range channel {
		fmt.Println(c)
	}
}
