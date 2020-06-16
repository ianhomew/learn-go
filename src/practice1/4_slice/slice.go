package main

import "fmt"

func main() {

	// append 函數 可以動態追加
	var intSlice = []int{1, 2, 3, 4, 5}

	var intSlice2 []int
	for i := 0; i < 5; i++ {
		intSlice2 = append(intSlice, i)
	}
	fmt.Println(intSlice2) // 123454

	for i := 0; i < 5; i++ {
		// 用原本的切片接收 才能追加
		intSlice = append(intSlice, i)
	}
	fmt.Println(intSlice) // 1234501234

	var intSlice3 []int
	intSlice3 = append(intSlice, intSlice...)
	fmt.Println(intSlice3)

}
