package main

import "fmt"

func main() {
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 299}
	fmt.Println(sl2)

	sl3 := []string{"a", "b", "c"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	// sl4[7] = 10
	// fmt.Println(sl4)
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", sl5)

	//append

	slAppend := []int{1, 2, 3, 4}
	fmt.Println(slAppend)
	slAppend = append(slAppend, 5)
	fmt.Println(slAppend)

	slAppend = append(slAppend, 6, 7)
	fmt.Println(slAppend)

	//make
	slmake := make([]int, 5)
	fmt.Println(slmake)
	//length
	fmt.Println(len(slmake))
	//cap
	fmt.Println(cap(slAppend))
}
