package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(time.Second * 2 )

	go func() {

		<-t1.C
		fmt.Println("Timer expired")
	} ()

	fmt.Scanln()
	stop := t1.Stop()
	if stop {
		fmt.Print("Timer stopped")
	}


}