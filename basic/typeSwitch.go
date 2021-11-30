package main

import (
	"fmt"
)

func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v, "string")
	case int:
		fmt.Println(v, "int")
	default:
		fmt.Println(v, "i don't know ")

	}
}
func main() {
	anything("aaa")
	anything([3]string{"a", "b", "c"})

	var x interface{} = [3]string{"a", "b", "c"}
	i, isInt := x.(int)
	fmt.Println(i, isInt)
	//interfaceと数値の足し算はできないよ
	// fmt.Println(x + 3)

	// panic: interface conversion: interface {} is int, not float64
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("none")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is string")
	} else {
		fmt.Println("i dont know")
	}

	switch v := x.(type) {
	case int:
		fmt.Println(v, "int")

	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println("i don't know")
	}
}
