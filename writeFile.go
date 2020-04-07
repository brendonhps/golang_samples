package main

import "os"

func main() {

	file, err := os.Create("cities.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	cities := [5]string{"Curitiba", "Sp", "Rj", "Porto Alegre", "Pinhais"}

	for _, x := range cities {
		file.WriteString(string(x))
		file.WriteString("\n")
	}

}
