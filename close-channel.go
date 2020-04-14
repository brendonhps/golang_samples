package main

import "fmt"

func main(){
	c := make(chan string, 4)
	c <- "teste1"
	c <- "teste2"
	c <- "teste3"
	fmt.Println(<-c)
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
