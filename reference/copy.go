package main

import "fmt"

func main() {
	sl := []int{100, 200, 300, 400, 500}
	sl2 := make([]int, 5, 10)
	n := copy(sl2, sl)
	sl2[0] = 1000
	// nはコピーに成功した数
	fmt.Println(n, sl, sl2)
}
