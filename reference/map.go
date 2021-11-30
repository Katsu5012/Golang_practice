package main

import "fmt"

func main() {
	var m = map[string]int{"a": 1, "b": 2}
	fmt.Println(m)

	m2 := make(map[int]string)
	m2[1] = "japan"
	m2[3] = "USA"
	fmt.Println(m2)

	s, err := m2[4]
	if !err {
		fmt.Println(err)
	}
	fmt.Println(s)
}
