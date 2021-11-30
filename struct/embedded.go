package main

import "fmt"

type T struct {
	User User
}

type User struct {
	name string
	age  int
}

func main() {
	t := T{User: User{name: "John", age: 10}}
	fmt.Println(t.User)

}
