package main

import "fmt"

func Sum(s ...int) int {
	n := 0

	for _, v := range s {
		n += v
	}
	return n

}
func main() {
	fmt.Println(Sum(1, 2, 3, 45))
	fmt.Println(Sum(10, 20, 50))
}
