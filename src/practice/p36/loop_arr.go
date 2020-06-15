package main

import (
	"fmt"
)

// go 的陣列是 傳值 call by value 也就是丟進去函式內 會進行值拷貝

func testArr(num [5]int) {
	//func testArr (num [4]int){ 會錯誤

	// [5]int 是一種資料類型    [5]int 與 [4]int 是不同的資料型別！！！！！！！
	num[0] = 999
	fmt.Println("num[0] =", num[0]) // 999
}

func testArrReference(num *[5]int) {
	// *[5]int 是一種資料類型
	num[0] = 888
	fmt.Println("num[0] =", num[0]) // 888
}

func main() {

	num := [...]int{1, 2, 3, 4, 5}

	// for 的 index, value 只有for 內有作用
	for index, value := range num {
		fmt.Println("index = ", index, "value = ", value)
	}

	fmt.Println("=====================")

	// call by value
	testArr(num)
	fmt.Println("num[0] =", num[0]) // 1

	fmt.Println("=====================")

	// call by reference
	testArrReference(&num)
	fmt.Println("num[0] =", num[0]) // 888

}
