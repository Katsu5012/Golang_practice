package main

import (
	"fmt"
)

func main() {
	//無限ループ
	// for {
	// 	fmt.Println("loop")
	// }

	// i := 0
	// for {
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}
	//３を飛ばすよ

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	//上と同じ結果が返ってくる
	for _, v := range arr {
		fmt.Println(v)
	}

	sl := []string{"a", "b", "c", "d", "e", "f"}

	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
