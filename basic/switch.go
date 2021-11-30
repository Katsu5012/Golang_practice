package main

import (
	"fmt"
)

func main() {
	// n := 1
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("i don't know ")
	// }

	// switch m := 2; m {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("i don't know ")
	// }

	p := 6
	switch {
	case p > 0 && p < 4:
		fmt.Println("0<p<4")
	case p >= 4 || p <= 0:
		fmt.Println("other")
	}
}
