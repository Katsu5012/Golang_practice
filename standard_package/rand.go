package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(256)

	fmt.Println(rand.Float64())

	//現在の時刻をシードに使った擬似乱数の作成
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	//0から99の擬似乱数
	fmt.Println(rand.Intn(100))
	//int型の擬似乱数
	fmt.Println(rand.Int())
	//int32型の擬似乱数
	fmt.Println(rand.Int31())
	//int64型の擬似乱数
	fmt.Println(rand.Int63())
}
