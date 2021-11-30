package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u User) SayName() {
	fmt.Println(u.name)
}

func (u *User) updateName(name string) {
	u.name = name
}
func (u User) updateName2(name string) {
	u.name = name
}
func main() {
	user1 := User{name: "user1", age: 20}
	user1.SayName()
	user2 := &User{name: "user2", age: 100}
	user2.SayName()
	user1.updateName2("update")
	user1.SayName()
	user2.updateName("update")
	user2.SayName()
}
