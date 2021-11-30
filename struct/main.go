package main

import "fmt"

type User struct {
	Name string
	age  int
	// x, y int
}

func updateUser(user User) User {
	user.Name = "a"
	user.age = 1111
	return user
}

func updateUser2(user *User) *User {
	user.Name = "b"
	user.age = 222222
	return user
}
func main() {
	var user1 User
	user1.Name = "user1"
	user1.age = 10
	fmt.Println(user1)
	user2 := User{}
	user2.Name = "user2"
	user2.age = 10
	fmt.Println(user2)
	user3 := User{Name: "user3", age: 10}
	fmt.Println(user3)
	user4 := new(User)
	fmt.Println(user4)
	user5 := &User{}
	fmt.Println(user5)

	fmt.Println(updateUser(user1))
	fmt.Println(updateUser2(user5))
	fmt.Println(user1)
	fmt.Println(user5)
}
