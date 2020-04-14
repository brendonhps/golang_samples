package main

import "fmt"

// Channel receive only


func f(c <- chan string ) {
	fmt.Println(<-c)
}

func main() {
	d := make(chan string, 1)

	d <- "testando"
	f(d)
}

// Channel send only

//func f2(c chan <- string) {
//	c <- "string in the function"
//}
//
//func main() {
//	d := make(chan string, 1)
//
//	f2(d)
//	fmt.Println(<-d)
//}