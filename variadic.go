package main

import "fmt"

func sum(numbers ...int) int {
	ret := 0
	for _, num := range numbers {
		ret += num
	}
	return ret
}

func nameStudents(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func main() {
	nameStudents("maria", "Carlos", "Jose", "Diego", "Joana", "Bianca", "Brendon")
	fmt.Println(sum(1, 2, 5, 7, 2, 5, 2, 9, 28723, -1722, -20000, 845))
}
