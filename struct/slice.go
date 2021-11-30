package main

import "fmt"

type User struct {
	name string
	age  int
}

type Users []*User

func main() {
	user1 := User{name: "John", age: 10}
	user2 := User{name: "tanaka", age: 20}
	user3 := User{name: "tarou", age: 30}
	users := Users{}
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3)
	for _, u := range users {
		fmt.Println(*u)
	}
	users2 := make([]*User, 0)
	users2 = append(users2, &user1)
	users2 = append(users2, &user2)
	for _, u := range users2 {
		fmt.Println(u)
	}
}
