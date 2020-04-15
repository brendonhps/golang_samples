package main

import (
	"encoding/json"
	"fmt"
)

type DayPrice struct {
	Gold float64
	Silver float64
	Platinum float64
}
func main() {
	data := `{"gold": 999,"silver": 666,"platinum": 333}`

	obj := new(DayPrice)

	err := json.Unmarshal([]byte(data), &obj)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Gold     : $ %f\n", obj.Gold);
	fmt.Printf("Silver   : $ %f\n", obj.Silver);
	fmt.Printf("Platinum : $ %f\n", obj.Platinum);
}

