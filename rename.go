package main

import (
	"fmt"
	"os"
)

func main() {

	src := "names.txt"
	dst := "nomes.txt"

	if os.Rename(src, dst) != nil {
		fmt.Println("Problem in the rename")
	}

	fmt.Println("File renamed")
}