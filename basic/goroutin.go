package main

import (
	"fmt"
	"time"
)

func sub() {
	for {

		fmt.Println("sub loop")
		time.Sleep(1000 * time.Millisecond)
	}
}
func main() {
	go sub()
	for {
		fmt.Println("main loop")
		time.Sleep(500 * time.Millisecond)
	}

}
