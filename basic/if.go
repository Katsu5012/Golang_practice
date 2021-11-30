package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("i don't know")
	}
	str := "hoge"
	if num, err := strconv.Atoi(str); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}

}
