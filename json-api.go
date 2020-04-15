package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	TradeID int
	Price string
	Size string
	Bid string
	Ask string
	Volume string
	Time string
}

func get_content() {

	url := "https://api.gdax.com/products/BTC-EUR/ticker"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}
	data := new(Response)

	json.Unmarshal(body, &data)

	// print values of the object
	fmt.Printf("Price: $ %s\n", data.Price)
	fmt.Printf("Price: $ %s\n", data.Bid)
	fmt.Printf("Price: $ %s\n", data.Ask)

	os.Exit(0)
}

func main() {
	get_content()
}