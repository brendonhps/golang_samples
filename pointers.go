package main

import "fmt"

func main() {
	var pointerI *int
	var aux [20]int
	for i := 0; i < 20; i++ {
		aux[i] = i
		pointerI = &aux[i]
		fmt.Printf("The address of aux is: %x\n", &aux[i])
		fmt.Printf("The address aponted by pointerIis : %x\n", pointerI)
		fmt.Printf("pointerI content is: %d\n", *pointerI)
	}
}
