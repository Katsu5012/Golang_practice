package main

import "fmt"

func double(i int) {
	i = i * 2
}

func doubleV2(i *int) {

	*i = *i * 2
}

func doubleV3(s []int) {
	for i, v := range s {
		s[i] = v * 2

	}
}
func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)
	double(n)
	fmt.Println(n)
	//参照渡し
	var p *int = &n
	double(*p)
	fmt.Println(p)
	fmt.Println(*p)

	*p = 300
	fmt.Println(p)
	doubleV2(p)
	fmt.Println(n)

	var sl []int = []int{1, 2, 3, 4, 5}
	doubleV3(sl)
	fmt.Println(sl)
}
