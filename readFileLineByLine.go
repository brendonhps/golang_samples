package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readLines(path string) ([]string, error) {

	file, err := os.Open(path)

	check(err)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, err
}

func main() {

	lines, err := readLines("names.txt")

	check(err)

	for i, line := range lines {
		fmt.Println(i, line)
	}
}
