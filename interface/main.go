package main

import "fmt"

type Stringify interface {
	ToString() string
}
type Person struct {
	name string
	age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v,age=%v", p.name, p.age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v,Model=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringify{
		&Person{name: "eji", age: 10},
		&Car{Number: "anbababa", Model: "aaaa"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}

}
