package main

import "fmt"

// function returning other function
func makeSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Nested function example

	teste := func() string {
		return "teste printed"
	}

	fmt.Println("Printe o teste\n", teste())

	// Clojure example
	// A closure is a type of function, that uses variables defined outside of the function itself.

	x := 0

	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	// Generator

	sequenceGenerator := makeSequence()
	fmt.Println(sequenceGenerator())
	fmt.Println(sequenceGenerator())
	fmt.Println(sequenceGenerator())
	fmt.Println(sequenceGenerator())
}
