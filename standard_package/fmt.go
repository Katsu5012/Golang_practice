package main

import "fmt"

func main() {
	fmt.Println("表示")
	fmt.Print("Hello\n")
	fmt.Println("Welcome")

	s := fmt.Sprint("hello")
	s1 := fmt.Sprintf("%v\n", "hello2")
	s2 := fmt.Sprintln("Hello3")

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
}
