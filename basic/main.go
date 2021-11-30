package main

import (
	"fmt"
	"strconv"
)

func CalFunc(f func()) {
	f()
}

const pi = 3.14

func ReturnFunc() func() {
	return func() {
		fmt.Println("im function")
	}
}

var i5 int = 500

func add(a, b int) int {
	return a + b
}

func Div(x, y int) (int, int) {
	return x / y, x % y
}
func double(x int) (result int) {
	result = x * 2
	return
}
func main() {
	aa := ReturnFunc()
	aa()
	var i int = 100
	fmt.Println(i)
	var s string = "Hello, world"
	fmt.Println(s)
	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 150
		s2 string = "Hello, again"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "hello"
	fmt.Println(i3, s3)

	i = 400
	fmt.Println(i)

	//暗黙的な定義
	// i4:=450

	fmt.Println(i5)
	fmt.Println(add(i2, i3))
	//型を調べる
	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))

	// float型
	var fl64 float64 = 2.4
	fmt.Println(fl64)
	var fl32 float32 = 2.1
	fmt.Printf("%T\n", float64(fl32))
	fmt.Println(fl64 + float64(fl32))

	//文字列
	fmt.Println("\"")
	fmt.Println(string(s[0]))

	//byte型
	byteA := [3]byte{72, 73}
	fmt.Println(byteA)

	// 配列
	var arr1 [3]int
	fmt.Println(arr1)
	var arr2 [3]string = [3]string{"a", "b"}
	fmt.Println(arr2)

	arr3 := []int{1, 2, 3, 4}
	fmt.Println(arr3)

	arr4 := [...]string{"a", "b", "c"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr4[0])
	arr2[2] = "c"
	fmt.Println(len(arr2))

	var x interface{}
	fmt.Println(x)
	x = 1
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	var int int = 100
	fl2_64 := float64(int)
	fmt.Printf("fl2_64=%T\n", fl2_64)
	// var string string = "100"
	int2, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", int2)

	int44 := "100"
	s55, _ := strconv.Atoi(int44)
	fmt.Printf("%T\n", s55)

	fmt.Println(pi)

	fmt.Println((1 == 1))

	fmt.Println(Div(10, 3))
	fmt.Println(double(10))

	CalFunc(func() {
		fmt.Println("hello")
	})
}
