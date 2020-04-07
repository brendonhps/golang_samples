package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := ioutil.ReadFile("names.txt")

	check(err)

	fmt.Println(file)

	fmt.Println(string(file))
}