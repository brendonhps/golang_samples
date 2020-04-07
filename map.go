package main

import "fmt"

func main() {

	// First way
	firstHash := map[string]int{
		"Coritiba":      1,
		"Atletico":      2,
		"Parana":        3,
		"Santos":        4,
		"Corintithians": 5,
	}

	// Second way
	secondHash := make(map[string]int)

	secondHash["Tempo"] = 12
	secondHash["Dia"] = 17
	secondHash["Mes"] = 12

	fmt.Println(firstHash)
	fmt.Println(secondHash)

}
