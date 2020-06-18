package main

import "fmt"

func main() {

	// declare
	var arr [5][5]int = [5][5]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println(arr)

	// arr2 其實是指向一個一維陣列 裡面存放真實陣列的位址
	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(arr2)
	fmt.Printf("arr[0]位址 = %p\n", &arr2[0])
	fmt.Printf("arr[1]位址 = %p\n", &arr2[1])

	// for
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Printf("arr2[%d][%d] = %d\n", i, j, arr2[i][j])
		}

	}

	// for range
	for _, value1 := range arr2 {
		for _, value2 := range value1 {
			fmt.Println(value2)
		}
	}
}
