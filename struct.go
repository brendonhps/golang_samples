package main

import "fmt"

type Motor struct {
	potencia      float64
	fabricante    string
	anoFabricacao int
}

func (m Motor) anoStart() int {
	return m.anoFabricacao
}

type Carro struct {
	cor       string
	motor     Motor
	anoL      int
	velMaxima float64
}

func (c Carro) anoStart() int {
	return c.anoL
}

type devAno interface {
	anoStart() int
}

func main() {

	newMotor := Motor{200.0, "Fiat", 1994}
	newCar := Carro{"azul", newMotor, 1998, 198.9}

	fmt.Println("O ano de lancamento do motor é:", newMotor.anoStart())
	fmt.Println("O ano de lancamento do carro é:", newCar.anoStart())

}
