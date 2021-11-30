package main

import (
	"fmt"
	"math"
)

func main() {
	//円周率
	fmt.Println(math.Pi)
	//２の平方根
	fmt.Println(math.Sqrt2)
	//float32で表現可能な最大値
	fmt.Println(math.MaxFloat32)
	//float32で表現可能な最小値
	fmt.Println(math.SmallestNonzeroFloat32)
	//int64
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)

	//絶対値
	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))
	//累乗
	fmt.Println(math.Pow(0, 0))
	fmt.Println(math.Pow(2, 3))
	//平方根,立方根
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Cbrt(8))

	//最大値、最小値
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))

	//小数点以下を切り捨てる
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))
	//引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))
	//引数の数値より大きい最小の整数を返す
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))
}
