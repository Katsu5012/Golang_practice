package main

import (
	"fmt"
	"time"
)

func receive(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(i, name)
	}
	fmt.Println(name + "end")
}
func main() {
	ch1 := make(chan int, 2)

	// close(ch1)

	// ch1 <- 1
	// i, ok := <-ch1
	//okはchannelのオープン状態がcloseかつ中身が空の時false
	// fmt.Println(i, ok)
	// fmt.Println(<-ch1)

	go receive("1.goroutin", ch1)
	go receive("2.goroutin", ch1)
	go receive("3.goroutin", ch1)

	i := 0
	for i < 10 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(300 * time.Second)
}
