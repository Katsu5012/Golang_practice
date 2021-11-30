package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	m := map[int]User{
		1: {name: "john", age: 1},
		2: {name: "tarou", age: 2},
	}
	fmt.Println(m)
	m2 := map[User]string{
		{name: "john", age: 1}:  "tokyo",
		{name: "tarou", age: 2}: "osaka",
	}
	fmt.Println(m2)
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	m3 := make(map[int]User)
	m3[-1] = User{name: "john", age: 1}
	fmt.Println(m3)
}
