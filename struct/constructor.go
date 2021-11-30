package main

import "fmt"

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) *User {
	return &User{name: name, age: age}
}
func main() {
	user1 := NewUser("user1", 10)
	fmt.Println(user1)

	fmt.Println((*user1))
}
