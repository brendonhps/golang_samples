package main 

import (
	"fmt"
	"os"
)

func fileExist(name string) (bool, error) {

	_, err := os.Stat(name)
	
	if err != nil {
		return false,err
	} 
	
	return true, err
}

func main() {

	if exist, _ := fileExist("gostudya.go"); exist == true {
		fmt.Println("The file exist")
	} else {
		fmt.Println("The file not exist")
	}	
}