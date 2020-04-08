package main

import "fmt"

func f(c chan string) {
	msg := "Inseri essa mensagem em c na função f()"
	fmt.Println(msg)
	c <- msg
}

func f2(c chan string) {
	msg := <-c
	fmt.Println("Aqui em f2 recebi tal mensagem: ", msg)
}

func main() {
	var c chan string = make(chan string)
	go f(c)
	go f2(c)

	fmt.Scanln()
}
