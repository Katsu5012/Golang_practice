package main

import "fmt"

func main() {
	sl := []string{"a", "b", "c"}

	fmt.Println(sl)

	for _, v := range sl {
		fmt.Println(v)
	}
}
